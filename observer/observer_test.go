// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package observer implements an event emitter and listener with builtin file watcher.
package observer

import (
	"testing"
)

func TestOpen(t *testing.T) {
	var err error
	var o Observer

	err = o.Open()
	if err != nil {
		t.Error("error Open a new Observer.")
	}

	err = o.Open()
	if err == nil {
		t.Error("no error trying to reopen a running Observer.")
	}
}

func TestClose(t *testing.T) {
	var err error
	var o Observer

	err = o.Close()
	if err != nil {
		t.Error("error Closing a not running Observer.")
	}

	o.Open()
	err = o.Close()
	if err != nil {
		t.Error("error Closing a running Observer.")
	}
}

func TestAddListener(t *testing.T) {
	var err error
	var o Observer

	o.Open()
	defer o.Close()

	err = o.AddListener(func(e interface{}) {
		// Do nothing
	})
	if err != nil {
		t.Error("error Adding a litner.")
	}
}

func TestEmit(t *testing.T) {
	var output string
	var o Observer

	o.Open()
	defer o.Close()

	done := make(chan bool)
	defer close(done)

	o.AddListener(func(e interface{}) {
		output = e.(string)
		done <- true
	})

	o.Emit("done")

	<-done // blocks until listener is triggered

	if output != "done" {
		t.Error("error Emitting strings.")
	}
}
