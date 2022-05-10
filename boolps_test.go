// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalBoolps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				b, b2 := true, false
				return map[string]json.Marshaler{"bool pointers to true and false": pfmt.Boolps([]*bool{&b, &b2})}
			}(),
			want:     "true false",
			wantText: "true false",
			wantJSON: `{
			"bool pointers to true and false":[true,false]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"bool pointers to nil": pfmt.Boolps([]*bool{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"bool pointers to nil":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without bool pointers": pfmt.Boolps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"without bool pointers":null
		}`,
		},
	}

	testMarshal(t, tests)
}
