// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalIntTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int": pfmt.Int(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int": pfmt.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int": pfmt.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"int pointer": pfmt.Intp(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"any int pointer": pfmt.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"reflect int pointer": pfmt.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int pointer":42
		}`,
	},
}

func TestMarshalInt(t *testing.T) {
	testMarshal(t, MarshalIntTests)
}