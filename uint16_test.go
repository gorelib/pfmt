// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalUint16(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint16": pfmt.Uint16(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint16":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint16": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint16":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint16": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint16":42
		}`,
		},
	}

	testMarshal(t, tests)
}
