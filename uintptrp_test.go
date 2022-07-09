// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalUintptrp(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uintptr = 42
				return map[string]json.Marshaler{"uintptr pointer": pfmt.Uintptrp(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uintptr pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uintptr pointer": pfmt.Uintptrp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uintptr pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uintptr = 42
				return map[string]json.Marshaler{"any uintptr pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uintptr pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uintptr = 42
				return map[string]json.Marshaler{"reflect uintptr pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uintptr pointer":42
		}`,
		},
	}

	testMarshal(t, tests)
}
