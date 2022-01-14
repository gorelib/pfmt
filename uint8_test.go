// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalUint8Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint8": pfmt.Uint8(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint8":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint8": pfmt.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint8":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint8": pfmt.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint8":42
		}`,
	},
}

func TestMarshalUint8(t *testing.T) {
	testMarshal(t, MarshalUint8Tests)
}
