// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalString(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"string": pfmt.String("Hello, Wörld!")},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"string":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty string": pfmt.String("")},
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty string":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"string with zero byte": pfmt.String(string(byte(0)))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"string with zero byte":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any string": pfmt.Any("Hello, Wörld!")},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any string":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any empty string": pfmt.Any("")},
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty string":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any string with zero byte": pfmt.Any(string(byte(0)))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"any string with zero byte":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect string": pfmt.Reflect("Hello, Wörld!")},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"reflect string":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect empty string": pfmt.Reflect("")},
			want:     "",
			wantText: "",
			wantJSON: `{
			"reflect empty string":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect string with zero byte": pfmt.Reflect(string(byte(0)))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"reflect string with zero byte":"\u0000"
		}`,
		},
	}

	testMarshal(t, tests)
}
