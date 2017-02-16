// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package webdriver provides a simple and incomplete WebDriver client for use by web test launcher.
package webdriver

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/launcher/errors"
	"github.com/bazelbuild/rules_webtesting/go/launcher/healthreporter"
)

const (
	compName           = "Go WebDriver Client"
	seleniumElementKey = "ELEMENT"
	w3cElementKey      = "element-6066-11e4-a52e-4f735466cecf"
)

// WebDriver provides access to a running WebDriver session
type WebDriver interface {
	healthreporter.HealthReporter
	// ExecuteScript executes script inside the browser's current execution context.
	ExecuteScript(ctx context.Context, script string, args []interface{}, value interface{}) error
	// ExecuteScriptAsync executes script asynchronously inside the browser's current execution context.
	ExecuteScriptAsync(ctx context.Context, script string, args []interface{}, value interface{}) error
	// Quit closes the WebDriver session.
	Quit(ctx context.Context) error
	// CommandURL builds a fully resolved URL for the specified end-point.
	CommandURL(endpoint ...string) (*url.URL, error)
	// SetScriptTimeout sets the timeout for the callback of an ExecuteScriptAsync call to be called.
	SetScriptTimeout(ctx context.Context, timeout time.Duration) error
	// Logs gets logs of the specified type from the remote end.
	Logs(ctx context.Context, logType string) ([]LogEntry, error)
	// SessionID returns the id for this session.
	SessionID() string
	// Address returns the base address for this sessions (ending with session/<SessionID>)
	Address() *url.URL
	// Capabilities returns the capabilities returned from the remote end when session was created.
	Capabilities() map[string]interface{}
	// Screenshot takes a screenshot of the current browser window.
	Screenshot(ctx context.Context) (image.Image, error)
	// WindowHandles returns a slice of the current window handles.
	WindowHandles(ctx context.Context) ([]string, error)
	// ElementFromID returns a new WebElement object for the given id.
	ElementFromID(string) WebElement
	// ElementFromMap returns a new WebElement from a map representing a JSON object.
	ElementFromMap(map[string]interface{}) (WebElement, error)
}

// WebElement provides access to a specific DOM element in a WebDriver session.
type WebElement interface {
	// ID returns the WebDriver element id.
	ID() string
	// ToMap returns a Map representation of a WebElement suitable for use in other WebDriver commands.
	ToMap() map[string]string
	// ScrollIntoView scrolls a WebElement to the top of the browsers viewport.
	ScrollIntoView(ctx context.Context) error
	// Bounds returns the bounds of the WebElement within the viewport.
	// This will not scroll the element into the viewport first.
	// Will return an error if the element is not in the viewport.
	Bounds(ctx context.Context) (image.Rectangle, error)
}

// LogEntry is an entry parsed from the logs retrieved from the remote WebDriver.
type LogEntry struct {
	Timestamp float64 `json:"timestamp"`
	Level     string  `json:"level"`
	Message   string  `json:"message"`
}

type webDriver struct {
	address      *url.URL
	sessionID    string
	capabilities map[string]interface{}
	client       *http.Client
}

type webElement struct {
	driver *webDriver
	id     string
}

type jsonResp struct {
	Status     int         `json:"status"`
	SessionID  string      `json:"sessionId"`
	Value      interface{} `json:"value"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	StackTrace interface{} `json:"stacktrace"`
}

// CreateSession creates a new WebDriver session with desired capabilities from server at addr
// and ensures that the browser connection is working. It retries up to attempts - 1 times.
func CreateSession(ctx context.Context, addr string, attempts int, desired map[string]interface{}) (WebDriver, error) {
	reqBody := map[string]interface{}{"desiredCapabilities": desired}

	urlPrefix, err := url.Parse(addr)
	if err != nil {
		return nil, errors.New(compName, err)
	}

	urlSuffix, err := url.Parse("session/")
	if err != nil {
		return nil, errors.New(compName, err)
	}

	fullURL := urlPrefix.ResolveReference(urlSuffix)
	c, err := command(fullURL, "")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	for ; attempts > 0; attempts-- {
		d, err := func() (*webDriver, error) {
			respBody, err := post(ctx, client, c, reqBody, nil)
			if err != nil {
				return nil, err
			}

			caps, ok := respBody.Value.(map[string]interface{})
			if !ok {
				caps = make(map[string]interface{})
			}

			session := respBody.SessionID
			if session == "" {
				// if we cannot cast to string, then empty string is fine.
				session, _ = caps["webdriver.remote.sessionid"].(string)
			}
			if session == "" {
				return nil, errors.New(compName, fmt.Errorf("no session id specified in %v", respBody))
			}

			sessionURL, err := url.Parse(session + "/")
			if err != nil {
				return nil, errors.New(compName, err)
			}

			d := &webDriver{
				address:      fullURL.ResolveReference(sessionURL),
				sessionID:    session,
				capabilities: caps,
				client:       client,
			}

			if err := d.Healthy(ctx); err != nil {
				if err := d.Quit(ctx); err != nil {
					log.Printf("error quitting WebDriver session: %v", err)
				}
				return nil, err
			}
			return d, nil
		}()

		if err == nil {
			return d, nil
		}
		if errors.IsPermanent(err) || attempts <= 1 {
			return nil, err
		}
	}

	// This should only occur if called with attempts <= 0
	return nil, errors.New(compName, fmt.Errorf("attempts %d <= 0", attempts))
}

func (d *webDriver) Address() *url.URL {
	return d.address
}

func (d *webDriver) Capabilities() map[string]interface{} {
	return d.capabilities
}

func (d *webDriver) SessionID() string {
	return d.sessionID
}

func (*webDriver) Name() string {
	return compName
}

func (d *webDriver) Healthy(ctx context.Context) error {
	return d.ExecuteScript(ctx, "return navigator.userAgent", nil, nil)
}

func (d *webDriver) ExecuteScript(ctx context.Context, script string, args []interface{}, value interface{}) error {
	if args == nil {
		args = []interface{}{}
	}
	body := map[string]interface{}{
		"script": script,
		"args":   args,
	}
	return d.post(ctx, "execute", body, value)
}

func (d *webDriver) ExecuteScriptAsync(ctx context.Context, script string, args []interface{}, value interface{}) error {
	if args == nil {
		args = []interface{}{}
	}
	body := map[string]interface{}{
		"script": script,
		"args":   args,
	}
	err := d.post(ctx, "execute_async", body, value)
	return err
}

// Screenshot takes a screenshot of the current browser window.
func (d *webDriver) Screenshot(ctx context.Context) (image.Image, error) {
	var value string
	if err := d.get(ctx, "screenshot", &value); err != nil {
		return nil, err
	}
	return png.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(value)))
}

// WindowHandles returns a slice of the current window handles.
func (d *webDriver) WindowHandles(ctx context.Context) ([]string, error) {
	var value []string
	if err := d.get(ctx, "window_handles", &value); err != nil {
		return nil, err
	}
	return value, nil
}

func (d *webDriver) post(ctx context.Context, suffix string, body interface{}, value interface{}) error {
	c, err := d.CommandURL(suffix)
	if err != nil {
		return err
	}

	_, err = post(ctx, d.client, c, body, value)
	return err
}

func (d *webDriver) get(ctx context.Context, suffix string, value interface{}) error {
	c, err := d.CommandURL(suffix)
	if err != nil {
		return err
	}
	_, err = getReq(ctx, d.client, c, value)
	return err
}

func (d *webDriver) delete(ctx context.Context, suffix string, value interface{}) error {
	c, err := d.CommandURL(suffix)
	if err != nil {
		return err
	}
	_, err = deleteReq(ctx, d.client, c, value)
	return err
}

func (d *webDriver) Quit(ctx context.Context) error {
	return d.delete(ctx, "", nil)
}

func (d *webDriver) CommandURL(endpoint ...string) (*url.URL, error) {
	return command(d.Address(), endpoint...)
}

func (d *webDriver) SetScriptTimeout(ctx context.Context, timeout time.Duration) error {
	body := map[string]interface{}{
		"type": "script",
		"ms":   int(timeout / time.Millisecond),
	}
	return d.post(ctx, "timeouts", body, nil)
}

func (d *webDriver) Logs(ctx context.Context, logType string) ([]LogEntry, error) {
	body := map[string]interface{}{"type": logType}
	var entries []LogEntry
	err := d.post(ctx, "log", body, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// ElementFromID returns a new WebElement object for the given id.
func (d *webDriver) ElementFromID(id string) WebElement {
	return &webElement{driver: d, id: id}
}

// ElementFromMap returns a new WebElement from a map representing a JSON object.
func (d *webDriver) ElementFromMap(m map[string]interface{}) (WebElement, error) {
	i, ok := m[w3cElementKey]
	if !ok {
		i, ok = m[seleniumElementKey]
		if !ok {
			return nil, errors.New(d.Name(), fmt.Errorf("map %v does not appear to represent a WebElement", m))
		}
	}

	id, ok := i.(string)
	if !ok {
		return nil, errors.New(d.Name(), fmt.Errorf("map %v does not appear to represent a WebElement", m))
	}
	return d.ElementFromID(id), nil
}

func command(addr *url.URL, endpoint ...string) (*url.URL, error) {
	u, err := addr.Parse(path.Join(endpoint...))
	if err != nil {
		return nil, err
	}
	return &url.URL{
		Scheme: u.Scheme,
		Opaque: u.Opaque,
		User:   u.User,
		Host:   u.Host,
		// Some remote ends (notably chromedriver) do not like a trailing slash
		Path:     strings.TrimRight(u.Path, "/"),
		RawPath:  strings.TrimRight(u.RawPath, "/"),
		RawQuery: u.RawQuery,
		Fragment: u.Fragment,
	}, nil
}

func processResponse(body io.Reader, value interface{}) (*jsonResp, error) {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New(compName, err)
	}

	respBody := &jsonResp{Value: value}
	var jsonErr error
	if err := json.Unmarshal(bytes, respBody); err == nil && respBody.Status == 0 && respBody.Error == "" {
		// WebDriver returned success, we are done.
		return respBody, nil
	} else if err != nil {
		jsonErr = err
		return nil, errors.New(compName, fmt.Errorf("%v unmarshalling %q", err, string(bytes)))
	}
	// if no value object was passed in then we can use the parsed Value
	if value == nil {
		return respBody, newWebDriverError(respBody)
	}

	// otherwise we can't trust the parsed Value has what we want, so need to re-parse.
	errBody := &jsonResp{}
	if err := json.Unmarshal(bytes, errBody); err != nil {
		return nil, errors.New(compName, fmt.Errorf("%v unmarshalling %q", err, respBody))
	}

	if errBody.Status == 0 && errBody.Error == "" {
		return nil, errors.New(compName, fmt.Errorf("%v unmarshalling %q", jsonErr, string(bytes)))
	}

	return errBody, newWebDriverError(errBody)
}

func post(ctx context.Context, client *http.Client, command *url.URL, body interface{}, value interface{}) (*jsonResp, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.NewPermanent(compName, err)
	}

	request, err := http.NewRequest("POST", command.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, errors.NewPermanent(compName, err)
	}

	request.TransferEncoding = []string{"identity"}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.ContentLength = int64(len(reqBody))

	return doRequest(ctx, client, request, value)
}

func deleteReq(ctx context.Context, client *http.Client, command *url.URL, value interface{}) (*jsonResp, error) {
	request, err := http.NewRequest("DELETE", command.String(), nil)
	if err != nil {
		return nil, errors.NewPermanent(compName, err)
	}

	return doRequest(ctx, client, request, value)
}

func getReq(ctx context.Context, client *http.Client, command *url.URL, value interface{}) (*jsonResp, error) {
	request, err := http.NewRequest("GET", command.String(), nil)
	if err != nil {
		return nil, errors.NewPermanent(compName, err)
	}

	return doRequest(ctx, client, request, value)
}

func doRequest(ctx context.Context, client *http.Client, request *http.Request, value interface{}) (*jsonResp, error) {
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Accept-Encoding", "identity")
	request = request.WithContext(ctx)
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New(compName, err)
	}

	defer resp.Body.Close()
	r, err := processResponse(resp.Body, value)
	return r, err
}

// ID returns the WebDriver element id.
func (e *webElement) ID() string {
	return e.id
}

// ToMap returns a Map representation of a WebElement suitable for use in other WebDriver commands.
func (e *webElement) ToMap() map[string]string {
	return map[string]string{
		seleniumElementKey: e.ID(),
		w3cElementKey:      e.ID(),
	}
}

// ScrollIntoView scrolls a WebElement to the top of the browsers viewport.
func (e *webElement) ScrollIntoView(ctx context.Context) error {
	const script = `return arguments[0].scrollIntoView(true);`
	args := []interface{}{e.ToMap()}
	return e.driver.ExecuteScript(ctx, script, args, nil)
}

// Bounds returns the bounds of the WebElement within the viewport.
// This will not scroll the element into the viewport first.
// Will return an error if the element is not in the viewport.
func (e *webElement) Bounds(ctx context.Context) (image.Rectangle, error) {
	const script = `
var element = arguments[0];
var rect = element.getBoundingClientRect();
var top = rect.top; var left = rect.left;
element = window.frameElement;
var currentWindow = window.parent;
while (element != null) {
  var currentRect = element.getBoundingClientRect();
  top += currentRect.top;
  left += currentRect.left;
  element = currentWindow.frameElement;
  currentWindow = currentWindow.parent;
}
return {"X0": Math.round(left), "Y0": Math.round(top), "X1": Math.round(left + rect.width), "Y1": Math.round(top + rect.height)};
`
	bounds := struct {
		X0 int
		Y0 int
		X1 int
		Y1 int
	}{}
	args := []interface{}{e.ToMap()}
	err := e.driver.ExecuteScript(ctx, script, args, &bounds)
	log.Printf("Err: %v", err)
	return image.Rect(bounds.X0, bounds.Y0, bounds.X1, bounds.Y1), err
}
