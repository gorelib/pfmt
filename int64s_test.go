// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalInt64sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int64 slice": pfmt.Int64s(123, 321)},
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"int64 slice":[123,321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int64": pfmt.Int64s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without int64":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int64 = 123, 321
			return map[string]json.Marshaler{"slice of any int64": pfmt.Anys(i, i2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of any int64":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int64 = 123, 321
			return map[string]json.Marshaler{"slice of int64 reflects": pfmt.Reflects(i, i2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of int64 reflects":[123,321]
		}`,
	},
}

func TestMarshalInt64s(t *testing.T) {
	testMarshal(t, MarshalInt64sTests)
}
