// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalJSON(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without jsons": pfmt.JSON(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"without jsons":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of empty jsons": pfmt.JSON([]pfmt.KV{pfmt.String(""), pfmt.String("")})},
			want:     ` `,
			wantText: ` `,
			wantJSON: `{
			"slice of empty jsons":{"":""}
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of json nils": pfmt.JSON([]pfmt.KV{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of json nils":{"":""}
		}`,
		},
	}

	testMarshal(t, tests)
}
