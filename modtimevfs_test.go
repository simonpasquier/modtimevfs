// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package modtimevfs

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestModTime(t *testing.T) {
	tim := time.Unix(1, 0)
	fs := New(http.Dir("./testdata"), tim)

	f, err := fs.Open("foo")
	if err != nil {
		t.Fatalf("unexpected error calling Open(): %v", err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		t.Fatalf("unexpected error calling Stat(): %v", err)
	}
	if fi.ModTime() != tim {
		t.Fatalf("expecting: %v, got: %v", tim, fi.ModTime())
	}

	exp := []byte(`foo`)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("unexpected error calling ReadAll(): %v", err)
	}
	if bytes.Compare(b, exp) != 0 {
		t.Fatalf("expecting: %v, got: %v", exp, b)
	}
}
