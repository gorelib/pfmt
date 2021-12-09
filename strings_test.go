// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalStringsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"strings": pfmt.Strings("Hello, Wörld!", "Hello, World!")},
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty strings": pfmt.Strings("", "")},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"empty strings":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"strings with zero byte": pfmt.Strings(string(byte(0)), string(byte(0)))},
		want:     "\\u0000 \\u0000",
		wantText: "\\u0000 \\u0000",
		wantJSON: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without strings": pfmt.Strings()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"without strings":null
		}`,
	},
}

func TestMarshalStrings(t *testing.T) {
	testMarshal(t, MarshalStringsTests)
}
