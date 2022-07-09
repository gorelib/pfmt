// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalUint16s(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint16 slice": pfmt.Uint16s([]uint16{42, 77})},
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uint16 slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uint16": pfmt.Uint16s(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without uint16":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uint16 = 42, 77
				return map[string]json.Marshaler{"slice of any uint16": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uint16":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uint16 = 42, 77
				return map[string]json.Marshaler{"slice of uint16 reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of uint16 reflects":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
