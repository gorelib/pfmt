// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalInt8ps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 int8 = 42, 77
				return map[string]json.Marshaler{"int8 pointer slice": pfmt.Int8ps([]*int8{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"int8 pointer slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil int8 pointers": pfmt.Int8ps([]*int8{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil int8 pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without int8 pointers": pfmt.Int8ps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without int8 pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 int8 = 42, 77
				return map[string]json.Marshaler{"slice of any int8 pointers": pfmt.Anys([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any int8 pointers":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 int8 = 42, 77
				return map[string]json.Marshaler{"slice of reflects of int8 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of reflects of int8 pointers":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
