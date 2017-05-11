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

package webdriver

import (
	"encoding/json"
	"fmt"
)

type errorDatum struct {
	Status     int
	Error      string
	HTTPStatus int
	W3C        bool
}

var errorData = []errorDatum{
	{
		0, "Success", 200, false,
	},
	{
		6, "invalid session id", 404, true,
	},
	{
		7, "no such element", 404, true,
	},
	{
		8, "no such frame", 404, true,
	},
	{
		9, "unknown command", 404, true,
	},
	{
		10, "stale element reference", 400, true,
	},
	{
		11, "ElementNotVisible", 400, false,
	},
	{
		12, "invalid element state", 400, true,
	},
	{
		13, "unknown error", 500, true,
	},
	{
		15, "element not selectable", 400, true,
	},
	{
		17, "javascript error", 500, true,
	},
	{
		19, "XPathLookupError", 400, false,
	},
	{
		21, "timeout", 408, true,
	},
	{
		23, "no such window", 400, true,
	},
	{
		24, "invalid cookie domain", 400, true,
	},
	{
		25, "unable to set cookie", 500, true,
	},
	{
		26, "unexpected alert open", 500, true,
	},
	{
		27, "no such alert", 400, true,
	},
	{
		28, "script timeout", 408, true,
	},
	{
		29, "invalid element coordinates", 400, true,
	},
	{
		30, "IMENotAvailable", 500, false,
	},
	{
		31, "IMEEngineActivationFailed", 500, false,
	},
	{
		32, "invalid selector", 400, true,
	},
	{
		33, "session not created", 500, true,
	},
	{
		34, "move target out of bounds", 400, true,
	},
	{
		-1, "element not interactable", 400, true,
	},
	{
		-1, "invalid argument", 400, true,
	},
	{
		-1, "no such cookie", 404, true,
	},
	{
		-1, "unable to capture screen", 500, true,
	},
	{
		-1, "unknown method", 405, true,
	},
	{
		-1, "unsupported operation", 500, true,
	},
}

type webDriverError struct {
	errDatum   errorDatum
	value      interface{}
	message    string
	stackTrace interface{}
}

func newWebDriverError(resp *jsonResp) error {
	return &webDriverError{
		errDatum:   errDatum(resp),
		value:      errValue(resp),
		message:    errMessage(resp),
		stackTrace: errStackTrace(resp),
	}
}

func errDatum(resp *jsonResp) errorDatum {
	if resp.Error != "" {
		for _, cand := range errorData {
			if cand.Error == resp.Error {
				return cand
			}
		}
	}
	if resp.Status != nil && *resp.Status != 0 {
		for _, cand := range errorData {
			if cand.Status == *resp.Status {
				return cand
			}
		}
	}
	return errorDatum{*resp.Status, resp.Error, 500, false}
}

func errMessage(resp *jsonResp) string {
	if resp.Message != "" {
		return resp.Message
	}
	value, _ := resp.Value.(map[string]interface{})
	message, _ := value["message"].(string)
	return message
}

func errStackTrace(resp *jsonResp) interface{} {
	if resp.StackTrace != nil {
		return resp.StackTrace
	}
	value, _ := resp.Value.(map[string]interface{})
	return value["stacktrace"]
}

func errValue(resp *jsonResp) interface{} {
	if resp.Value != nil {
		return resp.Value
	}
	val := map[string]interface{}{}
	if resp.Message != "" {
		val["message"] = resp.Message
	}
	if resp.StackTrace != nil {
		val["stacktrace"] = resp.StackTrace
	}
	return val
}

func (e *webDriverError) Component() string {
	return compName
}

func (e *webDriverError) Error() string {
	message := e.value
	if mapValue, ok := message.(map[string]interface{}); ok {
		if m, ok := mapValue["message"]; ok {
			message = m
		}
	}

	if e.errDatum.W3C {
		return fmt.Sprintf("[%s] (%s) %v", e.Component(), e.errDatum.Error, message)
	}

	return fmt.Sprintf("[%s] (%d) %v", e.Component(), e.errDatum.Status, message)
}

// IsWebDriverError returns true if err is a WebDriver Error.
func IsWebDriverError(err error) bool {
	_, ok := err.(*webDriverError)
	return ok
}

// ErrorStatus returns the WebDriver status for err.
func ErrorStatus(err error) int {
	we, ok := err.(*webDriverError)
	if !ok {
		return 13
	}
	if we.errDatum.Status <= 0 {
		return 13
	}
	return we.errDatum.Status
}

// ErrorValue returns the WebDriver value for err.
func ErrorValue(err error) interface{} {
	we, ok := err.(*webDriverError)
	if !ok {
		return map[string]interface{}{"message": err.Error()}
	}
	return we.value
}

// ErrorStackTrace returns the WebDriver value for err.
func ErrorStackTrace(err error) interface{} {
	we, ok := err.(*webDriverError)
	if !ok {
		return nil
	}
	return we.stackTrace
}

// ErrorMessage returns the WebDriver value for err.
func ErrorMessage(err error) string {
	we, ok := err.(*webDriverError)
	if !ok {
		return err.Error()
	}
	return we.message
}

// ErrorError returns the WebDriver error for err.
func ErrorError(err error) string {
	we, ok := err.(*webDriverError)
	if !ok {
		return "unknown error"
	}
	if !we.errDatum.W3C || we.errDatum.Error == "" {
		return "unknown error"
	}
	return we.errDatum.Error
}

// ErrorHTTPStatus returns the HTTP status code that is associated with err.
func ErrorHTTPStatus(err error) int {
	we, ok := err.(*webDriverError)
	if !ok {
		return 500
	}
	return we.errDatum.HTTPStatus

}

// MarshalError generates the WebDriver JSON wire protocol HTTP response body for err.
func MarshalError(err error) ([]byte, error) {
	body := map[string]interface{}{
		"status":  ErrorStatus(err),
		"value":   ErrorValue(err),
		"error":   ErrorError(err),
		"message": ErrorMessage(err),
	}

	st := ErrorStackTrace(err)
	if st != nil {
		body["stacktrace"] = st
	}

	return json.Marshal(body)
}
