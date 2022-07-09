// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalRunesps(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p, p2 := []rune("Hello, Wörld!"), []rune("Hello, World!")
				return map[string]json.Marshaler{"slice of rune slice pointers": pfmt.Runesps([]*[]rune{&p, &p2})}
			}(),
			want:     "Hello, Wörld! Hello, World!",
			wantText: "Hello, Wörld! Hello, World!",
			wantJSON: `{
			"slice of rune slice pointers":["Hello, Wörld!","Hello, World!"]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p, p2 := []rune{}, []rune{}
				return map[string]json.Marshaler{"slice of empty rune slice pointers": pfmt.Runesps([]*[]rune{&p, &p2})}
			}(),
			want:     " ",
			wantText: " ",
			wantJSON: `{
			"slice of empty rune slice pointers":["",""]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of nil rune slice pointers": pfmt.Runesps([]*[]rune{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"slice of nil rune slice pointers":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty slice of rune slice pointers": pfmt.Runesps(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"empty slice of rune slice pointers":null
		}`,
		},
	}

	testMarshal(t, tests)
}
