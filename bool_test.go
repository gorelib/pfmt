// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/gorelib/pfmt"
)

func TestMarshalBool(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"bool true": pfmt.Bool(true)},
			want:     "true",
			wantText: "true",
			wantJSON: `{
			"bool true":true
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"bool false": pfmt.Bool(false)},
			want:     "false",
			wantText: "false",
			wantJSON: `{
			"bool false":false
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any bool false": pfmt.Any(false)},
			want:     "false",
			wantText: "false",
			wantJSON: `{
			"any bool false":false
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect bool false": pfmt.Reflect(false)},
			want:     "false",
			wantText: "false",
			wantJSON: `{
			"reflect bool false":false
		}`,
		},
	}

	testMarshal(t, tests)
}
