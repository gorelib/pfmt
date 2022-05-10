// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalUint32(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint32": pfmt.Uint32(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint32":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint32": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint32":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint32": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint32":42
		}`,
		},
	}

	testMarshal(t, tests)
}
