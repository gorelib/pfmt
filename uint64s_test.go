// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalUint64s(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint64 slice": pfmt.Uint64s([]uint64{42, 77})},
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uint64 slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uint64": pfmt.Uint64s(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without uint64":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uint64 = 42, 77
				return map[string]json.Marshaler{"slice of any uint64": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uint64":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uint64 = 42, 77
				return map[string]json.Marshaler{"slice of uint64 reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of uint64 reflects":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
