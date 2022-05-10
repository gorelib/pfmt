// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalUint64ps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint64 = 42, 77
				return map[string]json.Marshaler{"uint64 pointer slice": pfmt.Uint64ps([]*uint64{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uint64 pointer slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil uint64 pointers": pfmt.Uint64ps([]*uint64{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil uint64 pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uint64 pointers": pfmt.Uint64ps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without uint64 pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint64 = 42, 77
				return map[string]json.Marshaler{"slice of any uint64 pointers": pfmt.Anys([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uint64 pointers":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f, f2 uint64 = 42, 77
				return map[string]json.Marshaler{"slice of reflects of uint64 pointers": pfmt.Reflects([]interface{}{&f, &f2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of reflects of uint64 pointers":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
