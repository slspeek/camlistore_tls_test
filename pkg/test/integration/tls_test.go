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

package integration

import (
	"camlistore.org/pkg/test"
	"testing"
)

func TestTls(t *testing.T) {
	w, err := test.WorldFromConfig("server-config-tls.json", true)
	if err != nil {
		t.Fatal(err)
	}
	SimplePut(t, w)
}

func TestPlaintext(t *testing.T) {
	w, err := test.WorldFromConfig("server-config.json", false)
	if err != nil {
		t.Fatal(err)
	}
	SimplePut(t, w)
}

func SimplePut(t *testing.T, w *test.World) {
	err := w.Start()
	if err != nil {
		t.Fatalf("could not start server : %v", err)
	}
	test.MustRunCmd(t, w.Cmd("camput", "file", "tls_test.go"))
	w.Stop()
}

func TestBuffer(t *testing.T) {
}
