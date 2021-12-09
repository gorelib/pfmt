// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pprint/pfmt"
)

var MarshalDurationpsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of durations pointers": pfmt.Durationps(&d, &d2)}
		}(),
		want:     "42ns 42s",
		wantText: "42ns 42s",
		wantJSON: `{
			"slice of durations pointers":["42ns","42s"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil durations pointers": pfmt.Durationps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil durations pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without durations pointers": pfmt.Durationps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without durations pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of any duration pointers": pfmt.Anys(&d, &d2)}
		}(),
		want:     "42ns 42s",
		wantText: "42ns 42s",
		wantJSON: `{
			"slice of any duration pointers":["42ns","42s"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of reflect of duration pointers": pfmt.Reflects(&d, &d2)}
		}(),
		want:     "42ns 42s",
		wantText: "42ns 42s",
		wantJSON: `{
			"slice of reflect of duration pointers":[42,42000000000]
		}`,
	},
}

func TestMarshalDurationps(t *testing.T) {
	testMarshal(t, MarshalDurationpsTests)
}
