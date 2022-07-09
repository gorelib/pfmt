// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalInt32s(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int32 slice": pfmt.Int32s([]int32{123, 321})},
			want:     "123 321",
			wantText: "123 321",
			wantJSON: `{
			"int32 slice":[123,321]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without int32": pfmt.Int32s(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without int32":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 int32 = 123, 321
				return map[string]json.Marshaler{"slice of any int32": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "123 321",
			wantText: "123 321",
			wantJSON: `{
			"slice of any int32":[123,321]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 int32 = 123, 321
				return map[string]json.Marshaler{"slice of int32 reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "123 321",
			wantText: "123 321",
			wantJSON: `{
			"slice of int32 reflects":[123,321]
		}`,
		},
	}

	testMarshal(t, tests)
}
