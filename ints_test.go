// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalIntsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int slice": pfmt.Ints(123, 321)},
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"int slice":[123,321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int": pfmt.Ints()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without int":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int = 123, 321
			return map[string]json.Marshaler{"slice of any int": pfmt.Anys(i, i2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of any int":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int = 123, 321
			return map[string]json.Marshaler{"slice of int reflects": pfmt.Reflects(i, i2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of int reflects":[123,321]
		}`,
	},
}

func TestMarshalInts(t *testing.T) {
	testMarshal(t, MarshalIntsTests)
}
