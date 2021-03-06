/*
Copyright 2014 The Camlistore Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package netutil

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

// AwaitReachable tries to make a TCP connection to addr regularly.
// It returns an error if it's unable to make a connection before maxWait.
func AwaitReachable(addr string, maxWait time.Duration) error {
	done := time.Now().Add(maxWait)
	for time.Now().Before(done) {
		c, err := net.Dial("tcp", addr)
		if err == nil {
			c.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("%v unreachable for %v", addr, maxWait)
}

// HostPort takes a urlStr string URL, and returns a host:port string suitable
// to passing to net.Dial, with the port set as the scheme's default port if
// absent.
func HostPort(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("could not parse %q as a url: %v", urlStr, err)
	}
	if u.Scheme == "" {
		return "", fmt.Errorf("url %q has no scheme", urlStr)
	}
	hostPort := u.Host
	if hostPort == "" || strings.HasPrefix(hostPort, ":") {
		return "", fmt.Errorf("url %q has no host", urlStr)
	}
	idx := strings.Index(hostPort, "]")
	if idx == -1 {
		idx = 0
	}
	if !strings.Contains(hostPort[idx:], ":") {
		if u.Scheme == "https" {
			hostPort += ":443"
		} else {
			hostPort += ":80"
		}
	}
	return hostPort, nil
}
