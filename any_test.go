// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalAny(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any struct": pfmt.Any(Struct{Name: "John Doe", Age: 42})},
			want:     "pfmt_test.Struct{Name:John Doe Age:42}",
			wantText: "pfmt_test.Struct{Name:John Doe Age:42}",
			wantJSON: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				s := Struct{Name: "John Doe", Age: 42}
				return map[string]json.Marshaler{"any struct pointer": pfmt.Any(&s)}
			}(),
			want:     "pfmt_test.Struct{Name:John Doe Age:42}",
			wantText: "pfmt_test.Struct{Name:John Doe Age:42}",
			wantJSON: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any byte array": pfmt.Any([3]byte{'f', 'o', 'o'})},
			want:     "[102 111 111]",
			wantText: "[102 111 111]",
			wantJSON: `{
			"any byte array":[102,111,111]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				a := [3]byte{'f', 'o', 'o'}
				return map[string]json.Marshaler{"any byte array pointer": pfmt.Any(&a)}
			}(),
			want:     "[102 111 111]",
			wantText: "[102 111 111]",
			wantJSON: `{
			"any byte array pointer":[102,111,111]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var a *[3]byte
				return map[string]json.Marshaler{"any byte array pointer to nil": pfmt.Any(a)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"any byte array pointer to nil":null
		}`,
		},
	}

	testMarshal(t, tests)
}
