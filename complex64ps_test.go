// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalComplex64ps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c, c2 complex64 = complex(1, 23), complex(3, 21)
				return map[string]json.Marshaler{"slice of complex64 pointers": pfmt.Complex64ps([]*complex64{&c, &c2})}
			}(),
			want:     "1+23i 3+21i",
			wantText: "1+23i 3+21i",
			wantJSON: `{
			"slice of complex64 pointers":["1+23i","3+21i"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil complex64 pointers": pfmt.Complex64ps([]*complex64{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil complex64 pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without complex64 pointers": pfmt.Complex64ps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"slice without complex64 pointers":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c, c2 complex64 = complex(1, 23), complex(3, 21)
				return map[string]json.Marshaler{"slice of any complex64 pointers": pfmt.Anys([]interface{}{&c, &c2})}
			}(),
			want:     "1+23i 3+21i",
			wantText: "1+23i 3+21i",
			wantJSON: `{
			"slice of any complex64 pointers":["1+23i","3+21i"]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c, c2 complex64 = complex(1, 23), complex(3, 21)
				return map[string]json.Marshaler{"slice of reflects of complex64 pointers": pfmt.Reflects([]interface{}{&c, &c2})}
			}(),
			want:     "1+23i 3+21i",
			wantText: "1+23i 3+21i",
			wantJSON: `{
			"slice of reflects of complex64 pointers": ["1+23i", "3+21i"]
		}`,
		},
	}

	testMarshal(t, tests)
}
