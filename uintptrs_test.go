// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalUintptrs(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uintptr slice": pfmt.Uintptrs([]uintptr{42, 77})},
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"uintptr slice":[42,77]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without uintptr": pfmt.Uintptrs(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice without uintptr":[]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uintptr = 42, 77
				return map[string]json.Marshaler{"slice of any uintptr": pfmt.Anys([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of any uintptr":[42,77]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i, i2 uintptr = 42, 77
				return map[string]json.Marshaler{"slice of uintptr reflects": pfmt.Reflects([]interface{}{i, i2})}
			}(),
			want:     "42 77",
			wantText: "42 77",
			wantJSON: `{
			"slice of uintptr reflects":[42,77]
		}`,
		},
	}

	testMarshal(t, tests)
}
