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

package driverhub

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const contentType = "application/json; charset=utf-8"

func success(w http.ResponseWriter, value interface{}) {
	body, err := json.Marshal(map[string]interface{}{
		"status": 0,
		"value":  value,
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func sessionNotCreated(w http.ResponseWriter, err error) {
	body, err := json.Marshal(map[string]interface{}{
		"status":  33,
		"error":   "session not created",
		"message": err.Error(),
		"value": map[string]string{
			"message": err.Error(),
		},
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(body)
}

func unknownError(w http.ResponseWriter, err error) {
	body, err := json.Marshal(map[string]interface{}{
		"status":  13,
		"error":   "unknown error",
		"message": err.Error(),
		"value": map[string]string{
			"message": err.Error(),
		},
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(body)
}

func unknownMethod(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s is not a supported method for %s", r.Method, r.URL.Path)
	body, err := json.Marshal(map[string]interface{}{
		"status":  9,
		"error":   "unknown method",
		"message": message,
		"value": map[string]string{
			"message": message,
		},
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(body)
}

func unknownCommand(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s %s is an unsupported webdriver command", r.Method, r.URL.Path)
	body, err := json.Marshal(map[string]interface{}{
		"status":  9,
		"error":   "unknown command",
		"message": message,
		"value": map[string]string{
			"message": message,
		},
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusNotFound)
	w.Write(body)
}

func invalidSessionID(w http.ResponseWriter, id string) {
	message := fmt.Sprintf("session %s does not exist", id)
	body, err := json.Marshal(map[string]interface{}{
		"status":  6,
		"error":   "invalid session id",
		"message": message,
		"value": map[string]string{
			"message": message,
		},
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusNotFound)
	w.Write(body)
}

func timeout(w http.ResponseWriter, endpoint string) {
	message := fmt.Sprintf("request to %q timed out", endpoint)
	body, err := json.Marshal(map[string]interface{}{
		"status":  21,
		"error":   "timeout",
		"message": message,
		"value": map[string]string{
			"message": message,
		},
	})
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusRequestTimeout)
	w.Write(body)
}
