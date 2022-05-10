// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalInt8s(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int8 slice": pfmt.Int8s([]int8{42, 77})},
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"int8 slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without int8": pfmt.Int8s(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without int8":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 int8 = 42, 77
				return map[string]json.Marshaler{"slice of any int8": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any int8":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 int8 = 42, 77
				return map[string]json.Marshaler{"slice of int8 reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of int8 reflects":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
