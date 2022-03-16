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
	"math"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/errors"
	"github.com/bazelbuild/rules_webtesting/go/healthreporter"
	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
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
	// ExecuteScriptAsyncWithTimeout executes the script asynchronously, but sets the script timeout to timeout before,
	// and attempts to restore it to its previous value after.
	ExecuteScriptAsyncWithTimeout(ctx context.Context, timeout time.Duration, script string, args []interface{}, value interface{}) error
	// Quit closes the WebDriver session.
	Quit(context.Context) error
	// CommandURL builds a fully resolved URL for the specified end-point.
	CommandURL(endpoint ...string) (*url.URL, error)
	// SetScriptTimeout sets the timeout for the callback of an ExecuteScriptAsync call to be called.
	SetScriptTimeout(context.Context, time.Duration) error
	// Logs gets logs of the specified type from the remote end.
	Logs(ctx context.Context, logType string) ([]LogEntry, error)
	// SessionID returns the id for this session.
	SessionID() string
	// Address returns the base address for this sessions (ending with session/<SessionID>)
	Address() *url.URL
	// Capabilities returns the capabilities returned from the remote end when session was created.
	Capabilities() map[string]interface{}
	// Screenshot takes a screenshot of the current browser window.
	Screenshot(context.Context) (image.Image, error)
	// ActiveElement returns the active element of the current browsing context's document element
	ActiveElement(ctx context.Context) (WebElement, error)
	// ElementScreenshot takes a screenshot of the visible region encompassed by the bounding rectangle of element.
	ElementScreenshot(ctx context.Context, el WebElement) (image.Image, error)
	// ElementGetAttribute gets the attribute of an element.
	ElementGetAttribute(ctx context.Context, el WebElement, attribute string) (string, error)
	// ElementGetText gets the text of an element.
	ElementGetText(ctx context.Context, el WebElement) (string, error)
	// ElementSendKeys sends keys to the element.
	ElementSendKeys(ctx context.Context, el WebElement, keys string) error
	// WindowHandles returns a slice of the current window handles.
	WindowHandles(context.Context) ([]string, error)
	// CurrentWindowHandle returns the handle of the active window.
	CurrentWindowHandle(context.Context) (string, error)
	// SwitchToFrame switches the current browsing context to:
	// nil: the top window (first frame)
	// int: the frame with the given index
	// WebElement: not yet supported
	SwitchToFrame(ctx context.Context, frame interface{}) error
	// SwitchToParentFrame switches the current browsing context to the parent of
	// the current browsing context.
	SwitchToParentFrame(ctx context.Context) error
	// SwitchToWindow switches the current browsing context to the window with
	// the given handle.
	SwitchToWindow(ctx context.Context, handle string) error
	// ElementFromID returns a new WebElement object for the given id.
	ElementFromID(string) WebElement
	// ElementFromMap returns a new WebElement from a map representing a JSON object.
	ElementFromMap(map[string]interface{}) (WebElement, error)
	// GetWindowRect returns the current windows size and location.
	GetWindowRect(context.Context) (Rectangle, error)
	// SetWindowRect sets the current window size and location.
	SetWindowRect(context.Context, Rectangle) error
	// SetWindowSize sets the current window size.
	SetWindowSize(ctx context.Context, width, height float64) error
	// SetWindowPosition sest the current window position.
	SetWindowPosition(ctx context.Context, x, y float64) error
	// W3C return true iff connected to a W3C compliant remote end.
	W3C() bool
	// CurrentURL returns the URL that the current browser window is looking at.
	CurrentURL(context.Context) (*url.URL, error)
	// PageSource returns the source of the current browsing context active document.
	PageSource(context.Context) (string, error)
	// NavigateTo navigates the controlled browser to the specified URL.
	NavigateTo(context.Context, *url.URL) error
	// ExecuteCDPCommand sends a CDP command to ChromeDriver.
	ExecuteCDPCommand(ctx context.Context, cmd string, params map[string]interface{}, value interface{}) error
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
	Bounds(ctx context.Context) (Rectangle, error)
}

// Rectangle represents a rectangle with floating point precision.
type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// ToImageRectangle converts webdriver.Rectangle to an image.Rectangle.
func (r *Rectangle) ToImageRectangle() image.Rectangle {
	return image.Rect(int(math.Trunc(r.X)), int(math.Trunc(r.Y)), int(math.Ceil(r.X+r.Width)), int(math.Ceil(r.Y+r.Height)))
}

// LogEntry is an entry parsed from the logs retrieved from the remote WebDriver.
type LogEntry struct {
	Timestamp float64 `json:"timestamp"`
	Level     string  `json:"level"`
	Message   string  `json:"message"`
}

type webDriver struct {
	address       *url.URL
	sessionID     string
	capabilities  map[string]interface{}
	client        *http.Client
	scriptTimeout time.Duration
	w3c           bool
}

type webElement struct {
	driver *webDriver
	id     string
}

type jsonResp struct {
	Status     *int        `json:"status"`
	SessionID  string      `json:"sessionId"`
	Value      interface{} `json:"value"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	StackTrace interface{} `json:"stacktrace"`
}

func (j *jsonResp) isError() bool {
	if j.Status != nil && *j.Status != 0 {
		return true
	}

	if j.Error != "" {
		return true
	}

	value, ok := j.Value.(map[string]interface{})
	if !ok {
		return false
	}

	e, ok := value["error"].(string)
	return ok && e != ""
}

// CreateSession creates a new WebDriver session with desired capabilities from server at addr
// and ensures that the browser connection is working. It retries up to attempts - 1 times.
func CreateSession(ctx context.Context, addr string, attempts int, requestedCaps *capabilities.Capabilities) (WebDriver, error) {
	reqBody := requestedCaps.ToMixedMode()

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

			val, ok := respBody.Value.(map[string]interface{})
			if !ok {
				return nil, errors.New(compName, fmt.Errorf("value field must be an object in %+v", respBody))
			}

			var caps map[string]interface{}

			session := respBody.SessionID
			if session != "" {
				// OSS protocol puts Session ID at the top level:
				// {
				//   "value": { capabilities object },
				//   "sessionId": "id",
				//   "status": 0
				// }
				caps = val
			} else {
				// W3C protocol wraps everything in a "value" key:
				// {
				//   "value": {
				//     "capabilities": { capabilities object },
				//     "sessionId": "id"
				//   }
				// }
				session, _ = val["sessionId"].(string)
				if session == "" {
					return nil, errors.New(compName, fmt.Errorf("no session id specified in %+v", respBody))
				}
				caps, ok = val["capabilities"].(map[string]interface{})
				if !ok {
					return nil, errors.New(compName, fmt.Errorf("no capabilities in value of %+v", respBody))
				}
			}

			sessionURL, err := url.Parse(session + "/")
			if err != nil {
				return nil, errors.New(compName, err)
			}

			d := &webDriver{
				address:       fullURL.ResolveReference(sessionURL),
				sessionID:     session,
				capabilities:  caps,
				client:        client,
				scriptTimeout: scriptTimeout(requestedCaps),
				w3c:           respBody.Status == nil,
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

func (d *webDriver) W3C() bool {
	return d.w3c
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
	command := "execute"
	if d.W3C() {
		command = "execute/sync"
	}
	return d.post(ctx, command, body, value)
}

func (d *webDriver) ExecuteScriptAsync(ctx context.Context, script string, args []interface{}, value interface{}) error {
	if args == nil {
		args = []interface{}{}
	}
	body := map[string]interface{}{
		"script": script,
		"args":   args,
	}
	command := "execute_async"
	if d.W3C() {
		command = "execute/async"
	}
	err := d.post(ctx, command, body, value)
	return err
}

func (d *webDriver) ExecuteScriptAsyncWithTimeout(ctx context.Context, timeout time.Duration, script string, args []interface{}, value interface{}) error {
	if err := d.setScriptTimeout(ctx, timeout); err != nil {
		log.Printf("error setting script timeout to %v", timeout)
	}
	if d.scriptTimeout != 0 {
		defer func() {
			if err := d.setScriptTimeout(ctx, d.scriptTimeout); err != nil {
				log.Printf("error restoring script timeout to %v", d.scriptTimeout)
			}
		}()
	}
	return d.ExecuteScriptAsync(ctx, script, args, value)
}

// CurrentURL returns the URL that the current browser window is looking at.
func (d *webDriver) CurrentURL(ctx context.Context) (*url.URL, error) {
	var result string

	if err := d.get(ctx, "url", &result); err != nil {
		return nil, err
	}

	current, err := url.Parse(result)
	if err != nil {
		return current, errors.New(d.Name(), err)
	}
	return current, nil
}

// PageSource returns the source of the current browsing context active document.
func (d *webDriver) PageSource(ctx context.Context) (string, error) {
	var result string

	if err := d.get(ctx, "source", &result); err != nil {
		return "", err
	}
	return result, nil
}

// NavigateTo navigates the controlled browser to the specified URL.
func (d *webDriver) NavigateTo(ctx context.Context, u *url.URL) error {
	return d.post(ctx, "url", map[string]interface{}{
		"url": u.String(),
	}, nil)
}

// Screenshot takes a screenshot of the current browser window.
func (d *webDriver) Screenshot(ctx context.Context) (image.Image, error) {
	var value string
	if err := d.get(ctx, "screenshot", &value); err != nil {
		return nil, err
	}
	return png.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(value)))
}

// ActiveElement returns the active element of the current browsing context's document element
func (d *webDriver) ActiveElement(ctx context.Context) (WebElement, error) {
	var value map[string]interface{}
	if err := d.get(ctx, "element/active", &value); err != nil {
		return nil, err
	}

	return d.ElementFromMap(value)
}

// ElementScreenshot takes a screenshot of the visible region encompassed by the bounding rectangle of element.
func (d *webDriver) ElementScreenshot(ctx context.Context, el WebElement) (image.Image, error) {
	var value string
	if err := d.get(ctx, fmt.Sprintf("element/%s/screenshot", el.ID()), &value); err != nil {
		return nil, err
	}
	return png.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(value)))
}

// ElementGetAttribute gets the attribute of an element.
func (d *webDriver) ElementGetAttribute(ctx context.Context, el WebElement, attribute string) (string, error) {
	var value string
	if err := d.get(ctx, fmt.Sprintf("element/%s/attribute/%s", el.ID(), attribute), &value); err != nil {
		return "", err
	}
	return value, nil
}

// ElementGetText gets the text of an element.
func (d *webDriver) ElementGetText(ctx context.Context, el WebElement) (string, error) {
	var value string
	if err := d.get(ctx, fmt.Sprintf("element/%s/text", el.ID()), &value); err != nil {
		return "", err
	}
	return value, nil
}

// ElementSendKeys sends keys to an element.
func (d *webDriver) ElementSendKeys(ctx context.Context, el WebElement, keys string) error {
	if err := d.post(ctx, fmt.Sprintf("element/%s/value", el.ID()), map[string]string{"text": keys}, nil); err != nil {
		return err
	}
	return nil
}

// WindowHandles returns a slice of the current window handles.
func (d *webDriver) WindowHandles(ctx context.Context) ([]string, error) {
	var value []string
	command := "window_handles"
	if d.W3C() {
		command = "window/handles"
	}
	if err := d.get(ctx, command, &value); err != nil {
		return nil, err
	}
	return value, nil
}

// CurrentWindowHandle returns the handle of the currently active window.
func (d *webDriver) CurrentWindowHandle(ctx context.Context) (string, error) {
	var value string
	command := "window_handle"
	if d.W3C() {
		command = "window"
	}
	if err := d.get(ctx, command, &value); err != nil {
		return "", err
	}
	return value, nil
}

func (d *webDriver) SwitchToFrame(ctx context.Context, frame interface{}) error {
	body := map[string]interface{}{}
	switch f := frame.(type) {
	case int, nil:
		body["id"] = f
	default:
		return fmt.Errorf("invalid type %T", frame)
	}
	return d.post(ctx, "frame", body, nil)
}

func (d *webDriver) SwitchToParentFrame(ctx context.Context) error {
	return d.post(ctx, "frame/parent", map[string]interface{}{}, nil)
}

func (d *webDriver) SwitchToWindow(ctx context.Context, handle string) error {
	body := make(map[string]string)
	if d.W3C() {
		body["handle"] = handle
	} else {
		body["name"] = handle
	}
	return d.post(ctx, "window", body, nil)
}

func (d *webDriver) GetWindowRect(ctx context.Context) (result Rectangle, err error) {
	if d.W3C() {
		err = d.get(ctx, "window/rect", &result)
		return
	}

	err = d.get(ctx, "window/current/size", &result)
	if err != nil {
		return
	}
	err = d.get(ctx, "window/current/position", &result)
	return
}

func (d *webDriver) SetWindowRect(ctx context.Context, rect Rectangle) error {
	if d.W3C() {
		return d.post(ctx, "window/rect", rect, nil)
	}

	if err := d.SetWindowSize(ctx, rect.Width, rect.Height); err != nil {
		return err
	}

	return d.SetWindowPosition(ctx, rect.X, rect.Y)
}

func (d *webDriver) SetWindowSize(ctx context.Context, width, height float64) error {
	args := map[string]float64{
		"width":  width,
		"height": height,
	}
	command := "window/current/size"
	if d.W3C() {
		command = "window/rect"
	}
	return d.post(ctx, command, args, nil)
}

func (d *webDriver) SetWindowPosition(ctx context.Context, x, y float64) error {
	args := map[string]float64{
		"x": x,
		"y": y,
	}
	command := "window/current/position"
	if d.W3C() {
		command = "window/rect"
	}
	return d.post(ctx, command, args, nil)
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
	d.scriptTimeout = timeout
	return d.setScriptTimeout(ctx, timeout)
}

func (d *webDriver) setScriptTimeout(ctx context.Context, timeout time.Duration) error {
	if d.W3C() {
		return d.post(ctx, "timeouts", map[string]interface{}{
			"script": int(timeout / time.Millisecond),
		}, nil)
	}
	return d.post(ctx, "timeouts", map[string]interface{}{
		"type": "script",
		"ms":   int(timeout / time.Millisecond),
	}, nil)
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

	// Parse once for errors.
	respBody := &jsonResp{}
	if err := json.Unmarshal(bytes, respBody); err != nil {
		return nil, errors.New(compName, fmt.Errorf("%v unmarshaling %s", err, string(bytes)))
	}
	if respBody.isError() {
		return respBody, newWebDriverError(respBody)
	}
	if value == nil {
		return respBody, nil
	}

	// Parse again to get value.
	respBodyValue := &jsonResp{Value: value}
	if err := json.Unmarshal(bytes, respBodyValue); err != nil {
		return respBody, errors.New(compName, fmt.Errorf("%v unmarshaling %v into %#v", err, string(bytes), value))
	}

	return respBody, nil
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
func (e *webElement) Bounds(ctx context.Context) (Rectangle, error) {
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
return {"X": left, "Y": top, "Width": rect.width, "Height": rect.height};
`
	bounds := Rectangle{}
	args := []interface{}{e.ToMap()}
	err := e.driver.ExecuteScript(ctx, script, args, &bounds)
	if err != nil {
		log.Printf("Err: %v", err)
	}
	return bounds, err
}

func scriptTimeout(caps *capabilities.Capabilities) time.Duration {
	if caps == nil {
		return 0
	}
	timeouts, ok := caps.AlwaysMatch["timeouts"].(map[string]interface{})
	if !ok {
		return 0
	}

	if script, ok := timeouts["script"].(int); ok {
		return time.Duration(script) * time.Millisecond
	}

	if script, ok := timeouts["script"].(float64); ok {
		return time.Duration(script) * time.Millisecond
	}

	return 0
}

func (d *webDriver) ExecuteCDPCommand(ctx context.Context, cmd string, params map[string]interface{}, value interface{}) error {
	return d.post(ctx, "goog/cdp/execute", map[string]interface{}{
		"cmd":    cmd,
		"params": params,
	}, value)
}
