// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalBytespsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte("Hello, Wörld!"), []byte("Hello, World!")
			return map[string]json.Marshaler{"slice of byte slice pointers": pfmt.Bytesps([]*[]byte{&p, &p2})}
		}(),
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"slice of byte slice pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte{}, []byte{}
			return map[string]json.Marshaler{"slice of empty byte slice pointers": pfmt.Bytesps([]*[]byte{&p, &p2})}
		}(),
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of empty byte slice pointers":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil byte slice pointers": pfmt.Bytesps([]*[]byte{nil, nil})},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil byte slice pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty slice of byte slice pointers": pfmt.Bytesps(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"empty slice of byte slice pointers":null
		}`,
	},
}

func TestMarshalBytesps(t *testing.T) {
	testMarshal(t, MarshalBytespsTests)
}
