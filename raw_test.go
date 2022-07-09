// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pfmt/pfmt"
)

func TestMarshalRaw(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of raw jsons": pfmt.Raw([]byte(`{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`))},
			want:     `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`,
			wantText: `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`,
			wantJSON: `{
			"slice of raw jsons":{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}
		}`,
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw with quote": pfmt.Raw([]byte(`Hello, "Wörld"!`))},
			want:      `Hello, "Wörld"!`,
			wantText:  `Hello, "Wörld"!`,
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'H' looking for beginning of value"),
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"quoted raw": pfmt.Raw([]byte(`"Hello, Wörld!"`))},
			want:     `"Hello, Wörld!"`,
			wantText: `"Hello, Wörld!"`,
			wantJSON: `{
			"quoted raw":"Hello, Wörld!"
		}`,
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw with nested quote": pfmt.Raw([]byte(`"Hello, "Wörld"!"`))},
			want:      `"Hello, "Wörld"!"`,
			wantText:  `"Hello, "Wörld"!"`,
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'W' after top-level value"),
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw quoted json": pfmt.Raw([]byte(`"{"foo":"bar"}"`))},
			want:      `"{"foo":"bar"}"`,
			wantText:  `"{"foo":"bar"}"`,
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'f' after top-level value"),
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw malformed json object": pfmt.Raw([]byte(`xyz{"foo":"bar"}`))},
			want:      `xyz{"foo":"bar"}`,
			wantText:  `xyz{"foo":"bar"}`,
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw malformed json key/value": pfmt.Raw([]byte(`{"foo":"bar""}`))},
			want:      `{"foo":"bar""}`,
			wantText:  `{"foo":"bar""}`,
			wantError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw json with unescaped null byte": pfmt.Raw(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...))},
			want:      "{\"foo\":\"\u0000xyz\"}",
			wantText:  "{\"foo\":\"\u0000xyz\"}",
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"raw nil": pfmt.Raw(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"raw nil":null
		}`,
		},
	}

	testMarshal(t, tests)
}
