// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalError(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"error": pfmt.Error(errors.New("something went wrong"))},
			want:     "something went wrong",
			wantText: "something went wrong",
			wantJSON: `{
			"error":"something went wrong"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil error": pfmt.Error(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil error":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any error": pfmt.Any(errors.New("something went wrong"))},
			want:     "something went wrong",
			wantText: "something went wrong",
			wantJSON: `{
			"any error":"something went wrong"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect error": pfmt.Reflect(errors.New("something went wrong"))},
			want:     "{something went wrong}",
			wantText: "{something went wrong}",
			wantJSON: `{
			"reflect error":{}
		}`,
		},
	}

	testMarshal(t, tests)
}
