// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalFloat64(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"float64": pfmt.Float64(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"float64":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"high precision float64": pfmt.Float64(0.123456789)},
			want:     "0.123456789",
			wantText: "0.123456789",
			wantJSON: `{
			"high precision float64":0.123456789
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"zero float64": pfmt.Float64(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"zero float64":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any float64": pfmt.Any(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"any float64":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any zero float64": pfmt.Any(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"any zero float64":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect float64": pfmt.Reflect(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"reflect float64":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect zero float64": pfmt.Reflect(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"reflect zero float64":0
		}`,
		},
	}

	testMarshal(t, tests)
}
