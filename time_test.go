// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pfmt/pfmt"
)

func TestTimeMarshalTest(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"time": time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC)},
			want:     "1970-01-01 02:03:04.000000042 +0000 UTC",
			wantText: "1970-01-01T02:03:04.000000042Z",
			wantJSON: `{
			"time":"1970-01-01T02:03:04.000000042Z"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any time": pfmt.Any(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC))},
			want:     `1970-01-01 02:03:04.000000042 +0000 UTC`,
			wantText: `1970-01-01T02:03:04.000000042Z`,
			wantJSON: `{
			"any time":"1970-01-01T02:03:04.000000042Z"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect time": pfmt.Reflect(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC))},
			want:     "1970-01-01 02:03:04.000000042 +0000 UTC",
			wantText: "1970-01-01 02:03:04.000000042 +0000 UTC",
			wantJSON: `{
			"reflect time":"1970-01-01T02:03:04.000000042Z"
		}`,
		},
	}

	testMarshal(t, tests)
}
