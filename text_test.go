// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalText(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"text": pfmt.Text(pfmt.String("Hello, Wörld!"))},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"text":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"text json": pfmt.Text(pfmt.String(`{"foo":"bar"}`))},
			want:     `{\"foo\":\"bar\"}`,
			wantText: `{\"foo\":\"bar\"}`,
			wantJSON: `{
			"text json":"{\"foo\":\"bar\"}"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"text with zero byte": pfmt.Text(pfmt.String("Hello, Wörld!\x00"))},
			want:     "Hello, Wörld!\\u0000",
			wantText: "Hello, Wörld!\\u0000",
			wantJSON: `{
			"text with zero byte":"Hello, Wörld!\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty text": pfmt.Text(pfmt.String(""))},
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty text":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"text nil": pfmt.Text(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"text nil":null
		}`,
		},
	}

	testMarshal(t, tests)
}
