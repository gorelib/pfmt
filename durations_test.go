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

var MarshalDurationsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of durations": pfmt.Durations(42*time.Nanosecond, 42*time.Second)},
		want:     "42ns 42s",
		wantText: "42ns 42s",
		wantJSON: `{
			"slice of durations":["42ns","42s"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without durations": pfmt.Durations()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without durations":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of any durations": pfmt.Anys(d, d2)}
		}(),
		want:     "42ns 42s",
		wantText: "42ns 42s",
		wantJSON: `{
			"slice of any durations":["42ns","42s"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of reflect of durations": pfmt.Reflects(d, d2)}
		}(),
		want:     "42ns 42s",
		wantText: "42ns 42s",
		wantJSON: `{
			"slice of reflect of durations":[42,42000000000]
		}`,
	},
}

func TestMarshalDurations(t *testing.T) {
	testMarshal(t, MarshalDurationsTests)
}
