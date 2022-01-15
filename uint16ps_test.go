// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalUint16psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint16 = 42, 77
			return map[string]json.Marshaler{"uint16 pointer slice": pfmt.Uint16ps([]*uint16{&f, &f2})}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"uint16 pointer slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil uint16 pointers": pfmt.Uint16ps([]*uint16{nil, nil})},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil uint16 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without uint16 pointers": pfmt.Uint16ps(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without uint16 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint16 = 42, 77
			return map[string]json.Marshaler{"slice of any uint16 pointers": pfmt.Anys([]interface{}{&f, &f2})}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any uint16 pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint16 = 42, 77
			return map[string]json.Marshaler{"slice of reflects of uint16 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of reflects of uint16 pointers":[42,77]
		}`,
	},
}

func TestMarshalUint16ps(t *testing.T) {
	testMarshal(t, MarshalUint16psTests)
}
