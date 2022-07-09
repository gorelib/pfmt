// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalComplex64(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"complex64": pfmt.Complex64(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"complex64":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"complex64": pfmt.Complex64(complex(3, 21))},
			want:     "3+21i",
			wantText: "3+21i",
			wantJSON: `{
			"complex64":"3+21i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any complex64": pfmt.Any(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"any complex64":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any complex64": pfmt.Any(complex(3, 21))},
			want:     "3+21i",
			wantText: "3+21i",
			wantJSON: `{
			"any complex64":"3+21i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect complex64": pfmt.Reflect(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"reflect complex64":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect complex64": pfmt.Reflect(complex(3, 21))},
			want:     "3+21i",
			wantText: "3+21i",
			wantJSON: `{
			"reflect complex64":"3+21i"
		}`,
		},
	}

	testMarshal(t, tests)
}
