// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalBools(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"bools true false": pfmt.Bools([]bool{true, false})},
			want:     "true false",
			wantText: "true false",
			wantJSON: `{
			"bools true false":[true,false]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without bools": pfmt.Bools(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"without bools":[]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any bools": pfmt.Anys([]interface{}{true, false})},
			want:     "true false",
			wantText: "true false",
			wantJSON: `{
			"any bools":[true, false]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflects bools": pfmt.Reflects([]interface{}{true, false})},
			want:     "true false",
			wantText: "true false",
			wantJSON: `{
			"reflects bools":[true, false]
		}`,
		},
	}

	testMarshal(t, tests)
}
