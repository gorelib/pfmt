// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalStringpTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"string pointer": pfmt.Stringp(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"empty string pointer": pfmt.Stringp(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty string pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil string pointer": pfmt.Stringp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"any string pointer": pfmt.Any(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"any empty string pointer": pfmt.Any(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"reflect string pointer": pfmt.Reflect(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"reflect empty string pointer": pfmt.Reflect(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"reflect empty string pointer":""
		}`,
	},
}

func TestMarshalStingp(t *testing.T) {
	testMarshal(t, MarshalStringpTests)
}
