// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalInt16s(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int16 slice": pfmt.Int16s([]int16{123, 321})},
			want:     "123 321",
			wantText: "123 321",
			wantJSON: `{
			"int16 slice":[123,321]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without int16": pfmt.Int16s(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without int16":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 int16 = 123, 321
				return map[string]json.Marshaler{"slice of any int16": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "123 321",
			wantText: "123 321",
			wantJSON: `{
			"slice of any int16":[123,321]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 int16 = 123, 321
				return map[string]json.Marshaler{"slice of int16 reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "123 321",
			wantText: "123 321",
			wantJSON: `{
			"slice of int16 reflects":[123,321]
		}`,
		},
	}

	testMarshal(t, tests)
}
