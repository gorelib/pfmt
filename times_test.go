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

var MarshalTimesTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"time slice": pfmt.Times(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		want:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantJSON: `{
			"time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any time slice": pfmt.Anys(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		want:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantJSON: `{
			"any time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect time slice": pfmt.Reflects(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		want:     "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		wantText: "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		wantJSON: `{
			"reflect time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
}

func TestMarshalTimes(t *testing.T) {
	testMarshal(t, MarshalTimesTests)
}
