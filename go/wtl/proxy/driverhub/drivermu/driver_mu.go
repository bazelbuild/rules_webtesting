// Package drivermu implements a handler that temporally serializes all commands to a session.
package drivermu

import (
	"context"
	"sync"

	"github.com/bazelbuild/rules_webtesting/go/metadata/capabilities"
	"github.com/bazelbuild/rules_webtesting/go/wtl/proxy/driverhub"
)

// ProviderFunc provides a handler that temporally serializes all commands to a session.
func ProviderFunc(_ *driverhub.WebDriverSession, _ *capabilities.Capabilities, base driverhub.HandlerFunc) (driverhub.HandlerFunc, bool) {
	var mu sync.Mutex

	return func(ctx context.Context, rq driverhub.Request) (driverhub.Response, error) {
		mu.Lock()
		defer mu.Unlock()
		return base(ctx, rq)
	}, true
}
