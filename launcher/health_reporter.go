// Package healthreporter provides polling wait functions.
package healthreporter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/web_test_launcher/launcher/errors"
)

const (
	pollMin     = 50 * time.Millisecond
	pollMax     = time.Second
	pollDefault = pollMin
	pollCount   = 20
)

// A HealthReporter knows its name and can check if it is healthy.
type HealthReporter interface {
	// Component is the name used in errors.
	Name() string
	// Healthy returns nil if this is healthy, and an Error describing the problem if not healthy.
	// Healthy should respect Context's Done.
	Healthy(context.Context) error
}

// ForHealthy waits until ctx is Done or h is Healthy, returning an error
// if h does not become healthy.
func WaitForHealthy(ctx context.Context, h HealthReporter) error {
	poll := pollDefault
	if deadline, ok := ctx.Deadline(); ok {
		poll = deadline.Sub(time.Now()) / pollCount
		if poll < pollMin {
			poll = pollMin
		}
		if poll > pollMax {
			poll = pollMax
		}
	} else {
		log.Printf("%s WaitForHealthy being called without deadline; will potentially wait forever.", h.Name())
	}

	ticker := time.NewTicker(poll)
	defer ticker.Stop()

	for {
		err := h.Healthy(ctx)
		if err == nil {
			return nil
		}
		if errors.IsPermanent(err) {
			return err
		}
		select {
		case <-ctx.Done():
			return errors.New(h.Name(), fmt.Errorf("%v waiting for healthy: %v", ctx.Err(), err))
		case <-ticker.C:
		}
	}
}
