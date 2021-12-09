// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalBoolsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bools true false": pfmt.Bools(true, false)},
		want:     "true false",
		wantText: "true false",
		wantJSON: `{
			"bools true false":[true,false]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without bools": pfmt.Bools()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"without bools":[]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any bools": pfmt.Anys(true, false)},
		want:     "true false",
		wantText: "true false",
		wantJSON: `{
			"any bools":[true, false]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflects bools": pfmt.Reflects(true, false)},
		want:     "true false",
		wantText: "true false",
		wantJSON: `{
			"reflects bools":[true, false]
		}`,
	},
}

func TestMarshalBools(t *testing.T) {
	testMarshal(t, MarshalBoolsTests)
}
