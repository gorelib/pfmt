// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalBytesTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bytes": pfmt.Bytes([]byte("Hello, Wörld!")...)},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bytes with quote": pfmt.Bytes([]byte(`Hello, "World"!`)...)},
		want:     `Hello, \"World\"!`,
		wantText: `Hello, \"World\"!`,
		wantJSON: `{
			"bytes with quote":"Hello, \"World\"!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bytes quote": pfmt.Bytes([]byte(`"Hello, World!"`)...)},
		want:     `\"Hello, World!\"`,
		wantText: `\"Hello, World!\"`,
		wantJSON: `{
			"bytes quote":"\"Hello, World!\""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bytes nested quote": pfmt.Bytes([]byte(`"Hello, "World"!"`)...)},
		want:     `\"Hello, \"World\"!\"`,
		wantText: `\"Hello, \"World\"!\"`,
		wantJSON: `{
			"bytes nested quote":"\"Hello, \"World\"!\""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bytes json": pfmt.Bytes([]byte(`{"foo":"bar"}`)...)},
		want:     `{\"foo\":\"bar\"}`,
		wantText: `{\"foo\":\"bar\"}`,
		wantJSON: `{
			"bytes json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bytes json quote": pfmt.Bytes([]byte(`"{"foo":"bar"}"`)...)},
		want:     `\"{\"foo\":\"bar\"}\"`,
		wantText: `\"{\"foo\":\"bar\"}\"`,
		wantJSON: `{
			"bytes json quote":"\"{\"foo\":\"bar\"}\""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty bytes": pfmt.Bytes([]byte{}...)},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var p []byte
			return map[string]json.Marshaler{"nil bytes": pfmt.Bytes(p...)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil bytes":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any bytes": pfmt.Any([]byte("Hello, Wörld!"))},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any empty bytes": pfmt.Any([]byte{})},
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty bytes":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect bytes": pfmt.Reflect([]byte("Hello, Wörld!"))},
		want:     "SGVsbG8sIFfDtnJsZCE=",
		wantText: "SGVsbG8sIFfDtnJsZCE=",
		wantJSON: `{
			"reflect bytes":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect empty bytes": pfmt.Reflect([]byte{})},
		want:     "",
		wantText: "",
		wantJSON: `{
			"reflect empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"any bytes pointer": pfmt.Any(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"any empty bytes pointer": pfmt.Any(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"reflect bytes pointer": pfmt.Reflect(&p)}
		}(),
		want:     "SGVsbG8sIFfDtnJsZCE=",
		wantText: "SGVsbG8sIFfDtnJsZCE=",
		wantJSON: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"reflect empty bytes pointer": pfmt.Reflect(&p)}
		}(),
		want: "",
		wantJSON: `{
			"reflect empty bytes pointer":""
		}`,
	},
}

func TestMarshalBytes(t *testing.T) {
	testMarshal(t, MarshalBytesTests)
}
