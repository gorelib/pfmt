// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalRunes(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"runes": pfmt.Runes([]rune("Hello, Wörld!"))},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"runes":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty runes": pfmt.Runes([]rune{})},
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty runes":""
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var p []rune
				return map[string]json.Marshaler{"nil runes": pfmt.Runes(p)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil runes":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"rune slice with zero rune": pfmt.Runes([]rune{rune(0)})},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"rune slice with zero rune":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any runes": pfmt.Any([]rune("Hello, Wörld!"))},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any runes":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any empty runes": pfmt.Any([]rune{})},
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty runes":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any rune slice with zero rune": pfmt.Any([]rune{rune(0)})},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"any rune slice with zero rune":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect runes": pfmt.Reflect([]rune("Hello, Wörld!"))},
			want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
			wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
			wantJSON: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect empty runes": pfmt.Reflect([]rune{})},
			want:     "[]",
			wantText: "[]",
			wantJSON: `{
			"reflect empty runes":[]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect rune slice with zero rune": pfmt.Reflect([]rune{rune(0)})},
			want:     "[0]",
			wantText: "[0]",
			wantJSON: `{
			"reflect rune slice with zero rune":[0]
		}`,
		},
	}

	testMarshal(t, tests)
}
