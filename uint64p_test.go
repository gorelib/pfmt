// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalUint64p(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint64 = 42
				return map[string]json.Marshaler{"uint64 pointer": pfmt.Uint64p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint64 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint64 pointer": pfmt.Uint64p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint64 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint64 = 42
				return map[string]json.Marshaler{"any uint64 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint64 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint64 = 42
				return map[string]json.Marshaler{"reflect uint64 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint64 pointer":42
		}`,
		},
	}

	testMarshal(t, tests)
}
