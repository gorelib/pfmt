// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pfmt/pfmt"
)

func TestMarshalJSONMarshalers(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"json slice": pfmt.JSONMarshalers([]json.Marshaler{time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)})},
			want:     `["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]`,
			wantText: `["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]`,
			wantJSON: `{
			"json slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without jsons": pfmt.JSONMarshalers(nil)},
			want:     `null`,
			wantText: `null`,
			wantJSON: `{
			"without jsons":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of empty jsons": pfmt.JSONMarshalers([]json.Marshaler{pfmt.String(""), pfmt.String("")})},
			want:     `["",""]`,
			wantText: `["",""]`,
			wantJSON: `{
			"slice of empty jsons":["",""]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of json nils": pfmt.JSONMarshalers([]json.Marshaler{nil, nil})},
			want:     `[null,null]`,
			wantText: `[null,null]`,
			wantJSON: `{
			"slice of json nils":[null,null]
		}`,
		},
	}

	testMarshal(t, tests)
}
