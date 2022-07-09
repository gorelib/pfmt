// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalComplex64p(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex64 = complex(1, 23)
				return map[string]json.Marshaler{"complex64 pointer": pfmt.Complex64p(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"complex64 pointer":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil complex64 pointer": pfmt.Complex64p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil complex64 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex64 = complex(1, 23)
				return map[string]json.Marshaler{"any complex64 pointer": pfmt.Any(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"any complex64 pointer":"1+23i"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex64 = complex(1, 23)
				return map[string]json.Marshaler{"reflect complex64 pointer": pfmt.Reflect(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"reflect complex64 pointer":"1+23i"
		}`,
		},
	}

	testMarshal(t, tests)
}
