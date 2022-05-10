// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalInt16(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int16": pfmt.Int16(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int16":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int16 = 42
				return map[string]json.Marshaler{"any int16": pfmt.Any(i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int16":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int16 = 42
				return map[string]json.Marshaler{"reflect int16": pfmt.Reflect(i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int16":42
		}`,
		},
	}

	testMarshal(t, tests)
}
