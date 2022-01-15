// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalUintps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint = 42, 77
				return map[string]json.Marshaler{"uint pointer slice": pfmt.Uintps([]*uint{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uint pointer slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil uint pointers": pfmt.Uintps([]*uint{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil uint pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uint pointers": pfmt.Uintps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without uint pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint = 42, 77
				return map[string]json.Marshaler{"slice of any uint pointers": pfmt.Anys([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uint pointers":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint = 42, 77
				return map[string]json.Marshaler{"slice of reflects of uint pointers": pfmt.Reflects([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of reflects of uint pointers":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
