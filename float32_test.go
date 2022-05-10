// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalFloat32(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"high precision float32": pfmt.Float32(0.123456789)},
			want:     "0.12345679",
			wantText: "0.12345679",
			wantJSON: `{
			"high precision float32":0.123456789
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"zero float32": pfmt.Float32(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"zero float32":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any float32": pfmt.Any(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"any float32":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any zero float32": pfmt.Any(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"any zero float32":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect float32": pfmt.Reflect(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"reflect float32":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect zero float32": pfmt.Reflect(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"reflect zero float32":0
		}`,
		},
	}

	testMarshal(t, tests)
}
