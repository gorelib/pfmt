// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gorelib/pfmt"
)

func TestMarshalDuration(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"duration": pfmt.Duration(42 * time.Nanosecond)},
			want:     "42ns",
			wantText: "42ns",
			wantJSON: `{
			"duration":"42ns"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any duration": pfmt.Any(42 * time.Nanosecond)},
			want:     "42ns",
			wantText: "42ns",
			wantJSON: `{
			"any duration":"42ns"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect duration": pfmt.Reflect(42 * time.Nanosecond)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect duration":42
		}`,
		},
	}

	testMarshal(t, tests)
}
