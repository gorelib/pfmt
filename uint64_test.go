// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalUint64(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint64": pfmt.Uint64(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint64": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint64": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint64":42
		}`,
		},
	}

	testMarshal(t, tests)
}
