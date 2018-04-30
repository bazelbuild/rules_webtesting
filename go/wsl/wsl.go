// Copyright 2018 Google Inc.
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

// WSL (Webdriver Server Light) is a lightweight replacement for Selenium Server.
package wsl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/wsl/hub"
	"github.com/bazelbuild/rules_webtesting/go/wsl/upload"
)

func Run(localHost string, port int, downloadRoot, uploadRoot string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := createHandler(hub.New(localHost, &upload.Uploader{Root: uploadRoot}), downloadRoot, cancel)

	if err := startServer(ctx, port, handler); err != nil {
		log.Print(err)
	}
}

func startServer(ctx context.Context, port int, handler http.Handler) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	errChan := make(chan error)

	go func() {
		log.Printf("Listening on %s", server.Addr)
		errChan <- server.ListenAndServe()
		close(errChan)
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	case err := <-errChan:
		return err
	}
}

func createHandler(hub http.Handler, downloadRoot string, shutdown func()) http.Handler {
	handler := http.NewServeMux()

	shutdownFunc := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("shutting down"))
		shutdown()
	}

	handler.HandleFunc("/quitquitquit", shutdownFunc)

	handler.HandleFunc("/shutdown", shutdownFunc)

	handler.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	handler.Handle("/session", hub)
	handler.Handle("/session/", hub)

	handler.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)

		respJSON := map[string]interface{}{
			"status": 0,
			"value": map[string]interface{}{
				"build": map[string]interface{}{
					"version":  "unknown",
					"revision": "unknown",
					"time":     "unknown",
				},
				"os": map[string]interface{}{
					"arch":    runtime.GOARCH,
					"name":    runtime.GOOS,
					"version": "unknown",
				},
				"ready":   true,
				"message": "ready to create new sessions",
			},
		}

		json.NewEncoder(w).Encode(respJSON)
	})

	handler.Handle("/google/staticfile/", http.StripPrefix("/google/staticfile/", http.FileServer(http.Dir(downloadRoot))))

	return handler
}
