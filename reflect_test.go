// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
)

func TestMarshalReflect(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"struct reflect": pfmt.Reflect(Struct{Name: "John Doe", Age: 42})},
			want:     "pfmt_test.Struct{Name:John Doe Age:42}",
			wantText: "pfmt_test.Struct{Name:John Doe Age:42}",
			wantJSON: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				s := Struct{Name: "John Doe", Age: 42}
				return map[string]json.Marshaler{"struct reflect pointer": pfmt.Reflect(&s)}
			}(),
			want:     "pfmt_test.Struct{Name:John Doe Age:42}",
			wantText: "pfmt_test.Struct{Name:John Doe Age:42}",
			wantJSON: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect byte array": pfmt.Reflect([3]byte{'f', 'o', 'o'})},
			want:     "[102 111 111]",
			wantText: "[102 111 111]",
			wantJSON: `{
			"reflect byte array":[102,111,111]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				a := [3]byte{'f', 'o', 'o'}
				return map[string]json.Marshaler{"reflect byte array pointer": pfmt.Reflect(&a)}
			}(),
			want:     "[102 111 111]",
			wantText: "[102 111 111]",
			wantJSON: `{
			"reflect byte array pointer":[102,111,111]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var a *[3]byte
				return map[string]json.Marshaler{"reflect byte array pointer to nil": pfmt.Reflect(a)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"reflect byte array pointer to nil":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect untyped nil": pfmt.Reflect(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"reflect untyped nil":null
		}`,
		},
	}

	testMarshal(t, tests)
}
