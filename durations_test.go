// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pprint/pfmt"
)

func TestMarshalDurations(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of durations": pfmt.Durations([]time.Duration{42 * time.Nanosecond, 42 * time.Second})},
			want:     "42ns 42s",
			wantText: "42ns 42s",
			wantJSON: `{
			"slice of durations":["42ns","42s"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice without durations": pfmt.Durations(nil)},
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
				return map[string]json.Marshaler{"slice of any durations": pfmt.Anys([]interface{}{d, d2})}
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
				return map[string]json.Marshaler{"slice of reflect of durations": pfmt.Reflects([]interface{}{d, d2})}
			}(),
			want:     "42ns 42s",
			wantText: "42ns 42s",
			wantJSON: `{
			"slice of reflect of durations":[42,42000000000]
		}`,
		},
	}

	testMarshal(t, tests)
}
