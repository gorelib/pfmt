// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalInt16psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int16 = 123, 321
			return map[string]json.Marshaler{"int16 pointer slice": pfmt.Int16ps([]*int16{&f, &f2})}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"int16 pointer slice":[123,321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil int16 pointers": pfmt.Int16ps([]*int16{nil, nil})},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil int16 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int16 pointers": pfmt.Int16ps(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without int16 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int16 = 123, 321
			return map[string]json.Marshaler{"slice of any int16 pointers": pfmt.Anys([]interface{}{&f, &f2})}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of any int16 pointers":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int16 = 123, 321
			return map[string]json.Marshaler{"slice of reflects of int16 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of reflects of int16 pointers":[123,321]
		}`,
	},
}

func TestMarshalInt16ps(t *testing.T) {
	testMarshal(t, MarshalInt16psTests)
}
