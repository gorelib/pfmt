// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalFloat32ps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 float32 = 0.123456789, 0.987654321
				return map[string]json.Marshaler{"float32 pointer slice": pfmt.Float32ps([]*float32{&f, &f2})}
			}(),
			want:     "0.12345679 0.9876543",
			wantText: "0.12345679 0.9876543",
			wantJSON: `{
			"float32 pointer slice":[0.123456789,0.987654321]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil float32 pointers": pfmt.Float32ps([]*float32{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil float32 pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without float32 pointers": pfmt.Float32ps(nil)},
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
				return map[string]json.Marshaler{"slice of any float32 pointers": pfmt.Anys([]interface{}{&f, &f2})}
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
				return map[string]json.Marshaler{"slice of reflects of float32 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
			}(),
			want:     "0.12345679 0.9876543",
			wantText: "0.12345679 0.9876543",
			wantJSON: `{
			"slice of reflects of float32 pointers":[0.123456789,0.987654321]
		}`,
		},
	}

	testMarshal(t, tests)
}
