// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalStringpsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 string = "Hello, Wörld!", "Hello, World!"
			return map[string]json.Marshaler{"string pointer slice": pfmt.Stringps([]*string{&f, &f2})}
		}(),
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"string pointer slice":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil string pointers": pfmt.Stringps([]*string{nil, nil})},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil string pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without string pointers": pfmt.Stringps(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without string pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 string = "Hello, Wörld!", "Hello, World!"
			return map[string]json.Marshaler{"slice of any string pointers": pfmt.Anys([]interface{}{&f, &f2})}
		}(),
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"slice of any string pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 string = "Hello, Wörld!", "Hello, World!"
			return map[string]json.Marshaler{"slice of reflects of string pointers": pfmt.Reflects([]interface{}{&f, &f2})}
		}(),
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"slice of reflects of string pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
}

func TestMarshalStringps(t *testing.T) {
	testMarshal(t, MarshalStringpsTests)
}
