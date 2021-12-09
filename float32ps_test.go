// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalFloat32psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"float32 pointer slice": pfmt.Float32ps(&f, &f2)}
		}(),
		want:     "0.12345679 0.9876543",
		wantText: "0.12345679 0.9876543",
		wantJSON: `{
			"float32 pointer slice":[0.123456789,0.987654321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil float32 pointers": pfmt.Float32ps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil float32 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without float32 pointers": pfmt.Float32ps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without float32 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of any float32 pointers": pfmt.Anys(&f, &f2)}
		}(),
		want:     "0.12345679 0.9876543",
		wantText: "0.12345679 0.9876543",
		wantJSON: `{
			"slice of any float32 pointers":[0.123456789,0.987654321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of reflects of float32 pointers": pfmt.Reflects(&f, &f2)}
		}(),
		want:     "0.12345679 0.9876543",
		wantText: "0.12345679 0.9876543",
		wantJSON: `{
			"slice of reflects of float32 pointers":[0.123456789,0.987654321]
		}`,
	},
}

func TestMarshalFloat32ps(t *testing.T) {
	testMarshal(t, MarshalFloat32psTests)
}
