// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalInt64(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int64": pfmt.Int64(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int64": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int64": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int64":42
		}`,
		},
	}

	testMarshal(t, tests)
}
