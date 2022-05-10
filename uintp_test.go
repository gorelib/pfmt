// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalUintp(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint = 42
				return map[string]json.Marshaler{"uint pointer": pfmt.Uintp(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint pointer": pfmt.Uintp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint = 42
				return map[string]json.Marshaler{"any uint pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint = 42
				return map[string]json.Marshaler{"reflect uint pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint pointer":42
		}`,
		},
	}

	testMarshal(t, tests)
}
