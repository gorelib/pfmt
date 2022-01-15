// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalInt64psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int64 = 123, 321
			return map[string]json.Marshaler{"int64 pointer slice": pfmt.Int64ps([]*int64{&f, &f2})}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"int64 pointer slice":[123,321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil int64 pointers": pfmt.Int64ps([]*int64{nil, nil})},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil int64 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int64 pointers": pfmt.Int64ps(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without int64 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int64 = 123, 321
			return map[string]json.Marshaler{"slice of any int64 pointers": pfmt.Anys([]interface{}{&f, &f2})}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of any int64 pointers":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int64 = 123, 321
			return map[string]json.Marshaler{"slice of reflects of int64 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of reflects of int64 pointers":[123,321]
		}`,
	},
}

func TestMarshalInt64ps(t *testing.T) {
	testMarshal(t, MarshalInt64psTests)
}
