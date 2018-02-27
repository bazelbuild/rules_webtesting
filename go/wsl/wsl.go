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
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bazelbuild/rules_webtesting/go/wsl/hub"
)

func Run(port int, downloadRoot string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := createHandler(hub.New(), downloadRoot, cancel)

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

	handler.HandleFunc("/quitquitquit", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("shutting down"))
		shutdown()
	})

	handler.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	handler.Handle("/session", hub)
	handler.Handle("/session/", hub)

	handler.Handle("/google/staticfile/", http.StripPrefix("/google/staticfile/", http.FileServer(http.Dir(downloadRoot))))

	return handler
}
