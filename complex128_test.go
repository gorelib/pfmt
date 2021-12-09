// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalComplex128Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"complex128": pfmt.Complex128(complex(1, 23))},
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any complex128": pfmt.Any(complex(1, 23))},
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"reflect complex128": pfmt.Reflect(complex(1, 23))},
		want:      "(1+23i)",
		wantText:  "(1+23i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128(t *testing.T) {
	testMarshal(t, MarshalComplex128Tests)
}
