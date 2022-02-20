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

func TestDurationpMarshal(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				d := 42 * time.Nanosecond
				return map[string]json.Marshaler{"duration pointer": pfmt.Durationp(&d)}
			}(),
			want:     "42ns",
			wantText: "42ns",
			wantJSON: `{
			"duration pointer":"42ns"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil duration pointer": pfmt.Durationp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil duration pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				d := 42 * time.Nanosecond
				return map[string]json.Marshaler{"any duration pointer": pfmt.Any(&d)}
			}(),
			want:     "42ns",
			wantText: "42ns",
			wantJSON: `{
			"any duration pointer":"42ns"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				d := 42 * time.Nanosecond
				return map[string]json.Marshaler{"reflect duration pointer": pfmt.Reflect(&d)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect duration pointer":42
		}`,
		},
	}

	testMarshal(t, tests)
}
