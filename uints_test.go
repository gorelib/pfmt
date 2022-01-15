// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalUints(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint slice": pfmt.Uints([]uint{42, 77})},
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uint slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uint": pfmt.Uints(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without uint":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uint = 42, 77
				return map[string]json.Marshaler{"slice of any uint": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uint":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uint = 42, 77
				return map[string]json.Marshaler{"slice of uint reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of uint reflects":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
