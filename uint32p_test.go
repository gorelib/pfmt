// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalUint32p(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint32 = 42
				return map[string]json.Marshaler{"uint32 pointer": pfmt.Uint32p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint32 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint32 pointer": pfmt.Uint32p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint32 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint32 = 42
				return map[string]json.Marshaler{"any uint32 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint32 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint32 = 42
				return map[string]json.Marshaler{"reflect uint32 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint32 pointer":42
		}`,
		},
	}

	testMarshal(t, tests)
}
