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

// Package httphelper provides simple wrappers for working with HTTP.
package httphelper

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

// Forward forwards r to host and writes the response from host to w.
func Forward(ctx context.Context, host, trimPrefix string, w http.ResponseWriter, r *http.Request) error {
	url, err := constructURL(host, r.URL.Path, trimPrefix)
	if err != nil {
		return err
	}

	// Construct request based on Method, URL Path, and Body from r
	request, err := http.NewRequest(r.Method, url.String(), r.Body)
	if err != nil {
		return err
	}
	request = request.WithContext(ctx)

	// Copy headers from r to request
	request.Header = r.Header

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Copy response headers from resp to w
	for k, vs := range resp.Header {
		w.Header().Del(k)
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	// Copy status code from resp to w
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
	return nil
}

// Get returns the contents located at url.
func Get(ctx context.Context, url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	return client.Do(request)
}

func constructURL(base, path, prefix string) (*url.URL, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}

	if !strings.HasPrefix(path, prefix) {
		return nil, fmt.Errorf("%q does not have expected prefix %q", path, prefix)
	}

	ref, err := url.Parse(strings.TrimPrefix(path, prefix))
	if err != nil {
		return nil, err
	}

	return u.ResolveReference(ref), err
}

type longestToShortest []string

func (s longestToShortest) Len() int {
	return len(s)
}

func (s longestToShortest) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s longestToShortest) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

// FQDN returns the fully-qualified domain name (or localhost if lookup 
// according to the hostname fails).
func FQDN() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		// Fail if the kernel fails to report a hostname.
		return "", err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return "localhost", nil
	}

	for _, addr := range addrs {
		if names, err := net.LookupAddr(addr); err == nil && len(names) > 0 {
			sort.Sort(longestToShortest(names))
			for _, name := range names {
				name = strings.TrimRight(name, ".")
				if strings.HasPrefix(name, hostname) {
					return name, nil
				}
			}
			return names[0], nil
		}
	}

	return "localhost", nil
}
