// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalIntp(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int = 42
				return map[string]json.Marshaler{"int pointer": pfmt.Intp(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int = 42
				return map[string]json.Marshaler{"any int pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int = 42
				return map[string]json.Marshaler{"reflect int pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int pointer":42
		}`,
		},
	}

	testMarshal(t, tests)
}
