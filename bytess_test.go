// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

// TODO: Add any/relfect test cases.
var MarshalBytessTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices": pfmt.Bytess([][]byte{[]byte("Hello, Wörld!"), []byte("Hello, World!")})},
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"slice of byte slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with quote": pfmt.Bytess([][]byte{[]byte(`Hello, "Wörld"!`), []byte(`Hello, "World"!`)})},
		want:     `Hello, \"Wörld\"! Hello, \"World\"!`,
		wantText: `Hello, \"Wörld\"! Hello, \"World\"!`,
		wantJSON: `{
			"slice of byte slices with quote":["Hello, \"Wörld\"!","Hello, \"World\"!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"quoted slice of byte slices": pfmt.Bytess([][]byte{[]byte(`"Hello, Wörld!"`), []byte(`"Hello, World!"`)})},
		want:     `\"Hello, Wörld!\" \"Hello, World!\"`,
		wantText: `\"Hello, Wörld!\" \"Hello, World!\"`,
		wantJSON: `{
			"quoted slice of byte slices":["\"Hello, Wörld!\"","\"Hello, World!\""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with nested quote": pfmt.Bytess([][]byte{[]byte(`"Hello, "Wörld"!"`), []byte(`"Hello, "World"!"`)})},
		want:     `\"Hello, \"Wörld\"!\" \"Hello, \"World\"!\"`,
		wantText: `\"Hello, \"Wörld\"!\" \"Hello, \"World\"!\"`,
		wantJSON: `{
			"slice of byte slices with nested quote":["\"Hello, \"Wörld\"!\"","\"Hello, \"World\"!\""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with json": pfmt.Bytess([][]byte{[]byte(`{"foo":"bar"}`), []byte(`{"baz":"xyz"}`)})},
		want:     `{\"foo\":\"bar\"} {\"baz\":\"xyz\"}`,
		wantText: `{\"foo\":\"bar\"} {\"baz\":\"xyz\"}`,
		wantJSON: `{
			"slice of byte slices with json":["{\"foo\":\"bar\"}","{\"baz\":\"xyz\"}"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with quoted json": pfmt.Bytess([][]byte{[]byte(`"{"foo":"bar"}"`), []byte(`"{"baz":"xyz"}"`)})},
		want:     `\"{\"foo\":\"bar\"}\" \"{\"baz\":\"xyz\"}\"`,
		wantText: `\"{\"foo\":\"bar\"}\" \"{\"baz\":\"xyz\"}\"`,
		wantJSON: `{
			"slice of byte slices with quoted json":["\"{\"foo\":\"bar\"}\"","\"{\"baz\":\"xyz\"}\""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of empty byte slice": pfmt.Bytess([][]byte{[]byte{}, []byte{}})},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of empty byte slice":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil byte slice": pfmt.Bytess([][]byte{nil, nil})},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil byte slice":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty slice of byte slices": pfmt.Bytess(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"empty slice of byte slices":null
		}`,
	},
}

func TestMarshalBytess(t *testing.T) {
	testMarshal(t, MarshalBytessTests)
}
