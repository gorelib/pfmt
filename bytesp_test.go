// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/gorelib/pfmt"
)

func TestMarshalBytesp(t *testing.T) {
	tests := []marshalTest{
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []byte("Hello, Wörld!")
				return map[string]json.Marshaler{"bytes pointer": pfmt.Bytesp(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"bytes pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []byte{}
				return map[string]json.Marshaler{"empty bytes pointer": pfmt.Bytesp(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty bytes pointer":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil bytes pointer": pfmt.Bytesp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil bytes pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []byte("Hello, Wörld!")
				return map[string]json.Marshaler{"any bytes pointer": pfmt.Any(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []byte{}
				return map[string]json.Marshaler{"any empty bytes pointer": pfmt.Any(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty bytes pointer":""
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []byte("Hello, Wörld!")
				return map[string]json.Marshaler{"reflect bytes pointer": pfmt.Reflect(&p)}
			}(),
			want:     "SGVsbG8sIFfDtnJsZCE=",
			wantText: "SGVsbG8sIFfDtnJsZCE=",
			wantJSON: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []byte{}
				return map[string]json.Marshaler{"reflect empty bytes pointer": pfmt.Reflect(&p)}
			}(),
			want: "",
			wantJSON: `{
			"reflect empty bytes pointer":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"complex128": pfmt.Complex128(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"complex128":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any complex128": pfmt.Any(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"any complex128":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect complex128": pfmt.Reflect(complex(1, 23))},
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"reflect complex128":"1+23i"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex128 = complex(1, 23)
				return map[string]json.Marshaler{"complex128 pointer": pfmt.Complex128p(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"complex128 pointer":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil complex128 pointer": pfmt.Complex128p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil complex128 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex128 = complex(1, 23)
				return map[string]json.Marshaler{"any complex128 pointer": pfmt.Any(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"any complex128 pointer":"1+23i"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex128 = complex(1, 23)
				return map[string]json.Marshaler{"reflect complex128 pointer": pfmt.Reflect(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"reflect complex128 pointer":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"complex64": pfmt.Complex64(complex(3, 21))},
			want:     "3+21i",
			wantText: "3+21i",
			wantJSON: `{
			"complex64":"3+21i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any complex64": pfmt.Any(complex(3, 21))},
			want:     "3+21i",
			wantText: "3+21i",
			wantJSON: `{
			"any complex64":"3+21i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect complex64": pfmt.Reflect(complex(3, 21))},
			want:     "3+21i",
			wantText: "3+21i",
			wantJSON: `{
			"reflect complex64":"3+21i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"error": pfmt.Err(errors.New("something went wrong"))},
			want:     "something went wrong",
			wantText: "something went wrong",
			wantJSON: `{
			"error":"something went wrong"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil error": pfmt.Err(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil error":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"errors": pfmt.Errs([]error{errors.New("something went wrong"), errors.New("wrong")})},
			want:     "something went wrong wrong",
			wantText: "something went wrong wrong",
			wantJSON: `{
			"errors":["something went wrong","wrong"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil errors": pfmt.Errs([]error{nil, nil})},
			want:     "null null",
			wantText: "null null",
			wantJSON: `{
			"nil errors":[null,null]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without errors": pfmt.Errs(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"without errors":null
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
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex64 = complex(1, 23)
				return map[string]json.Marshaler{"complex64 pointer": pfmt.Complex64p(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"complex64 pointer":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil complex64 pointer": pfmt.Complex64p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil complex64 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex64 = complex(1, 23)
				return map[string]json.Marshaler{"any complex64 pointer": pfmt.Any(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"any complex64 pointer":"1+23i"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var c complex64 = complex(1, 23)
				return map[string]json.Marshaler{"reflect complex64 pointer": pfmt.Reflect(&c)}
			}(),
			want:     "1+23i",
			wantText: "1+23i",
			wantJSON: `{
			"reflect complex64 pointer":"1+23i"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"float32": pfmt.Float32(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"float32":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"high precision float32": pfmt.Float32(0.123456789)},
			want:     "0.12345679",
			wantText: "0.12345679",
			wantJSON: `{
			"high precision float32":0.123456789
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"zero float32": pfmt.Float32(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"zero float32":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any float32": pfmt.Any(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"any float32":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any zero float32": pfmt.Any(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"any zero float32":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect float32": pfmt.Reflect(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"reflect float32":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect zero float32": pfmt.Reflect(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"reflect zero float32":0
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float32 = 4.2
				return map[string]json.Marshaler{"float32 pointer": pfmt.Float32p(&f)}
			}(),
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"float32 pointer":4.2
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float32 = 0.123456789
				return map[string]json.Marshaler{"high precision float32 pointer": pfmt.Float32p(&f)}
			}(),
			want:     "0.12345679",
			wantText: "0.12345679",
			wantJSON: `{
			"high precision float32 pointer":0.123456789
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"float32 nil pointer": pfmt.Float32p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"float32 nil pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float32 = 4.2
				return map[string]json.Marshaler{"any float32 pointer": pfmt.Any(&f)}
			}(),
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"any float32 pointer":4.2
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float32 = 4.2
				return map[string]json.Marshaler{"reflect float32 pointer": pfmt.Reflect(&f)}
			}(),
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"reflect float32 pointer":4.2
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f *float32
				return map[string]json.Marshaler{"reflect float32 pointer to nil": pfmt.Reflect(f)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"reflect float32 pointer to nil":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"float64": pfmt.Float64(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"float64":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"high precision float64": pfmt.Float64(0.123456789)},
			want:     "0.123456789",
			wantText: "0.123456789",
			wantJSON: `{
			"high precision float64":0.123456789
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"zero float64": pfmt.Float64(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"zero float64":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any float64": pfmt.Any(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"any float64":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any zero float64": pfmt.Any(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"any zero float64":0
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect float64": pfmt.Reflect(4.2)},
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"reflect float64":4.2
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect zero float64": pfmt.Reflect(0)},
			want:     "0",
			wantText: "0",
			wantJSON: `{
			"reflect zero float64":0
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float64 = 4.2
				return map[string]json.Marshaler{"float64 pointer": pfmt.Float64p(&f)}
			}(),
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"float64 pointer":4.2
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float64 = 0.123456789
				return map[string]json.Marshaler{"high precision float64 pointer": pfmt.Float64p(&f)}
			}(),
			want:     "0.123456789",
			wantText: "0.123456789",
			wantJSON: `{
			"high precision float64 pointer":0.123456789
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"float64 nil pointer": pfmt.Float64p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"float64 nil pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float64 = 4.2
				return map[string]json.Marshaler{"any float64 pointer": pfmt.Any(&f)}
			}(),
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"any float64 pointer":4.2
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f float64 = 4.2
				return map[string]json.Marshaler{"reflect float64 pointer": pfmt.Reflect(&f)}
			}(),
			want:     "4.2",
			wantText: "4.2",
			wantJSON: `{
			"reflect float64 pointer":4.2
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var f *float64
				return map[string]json.Marshaler{"reflect float64 pointer to nil": pfmt.Reflect(f)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"reflect float64 pointer to nil":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int": pfmt.Int(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int = 42
				return map[string]json.Marshaler{"int pointer": pfmt.Intp(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int = 42
				return map[string]json.Marshaler{"any int pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int = 42
				return map[string]json.Marshaler{"reflect int pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int16": pfmt.Int16(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int16":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int16": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int16":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int16": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int16":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int16 = 42
				return map[string]json.Marshaler{"int16 pointer": pfmt.Int16p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int16 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int16 = 42
				return map[string]json.Marshaler{"any int16 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int16 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int16 = 42
				return map[string]json.Marshaler{"reflect int16 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int16 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int32": pfmt.Int32(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int32":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int32": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int32":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int32": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int32":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int32 = 42
				return map[string]json.Marshaler{"int32 pointer": pfmt.Int32p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int32 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int32 = 42
				return map[string]json.Marshaler{"any int32 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int32 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int32 = 42
				return map[string]json.Marshaler{"reflect int32 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int32 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int64": pfmt.Int64(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int64": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int64": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int64":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int64 = 42
				return map[string]json.Marshaler{"int64 pointer": pfmt.Int64p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int64 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int64 = 42
				return map[string]json.Marshaler{"any int64 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int64 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int64 = 42
				return map[string]json.Marshaler{"reflect int64 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int64 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"int8": pfmt.Int8(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int8":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any int8": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int8":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect int8": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int8":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int8 = 42
				return map[string]json.Marshaler{"int8 pointer": pfmt.Int8p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"int8 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int8 = 42
				return map[string]json.Marshaler{"any int8 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any int8 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i int8 = 42
				return map[string]json.Marshaler{"reflect int8 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect int8 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"runes": pfmt.Runes([]rune("Hello, Wörld!"))},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"runes":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty runes": pfmt.Runes([]rune{})},
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty runes":""
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var p []rune
				return map[string]json.Marshaler{"nil runes": pfmt.Runes(p)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil runes":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"rune slice with zero rune": pfmt.Runes([]rune{rune(0)})},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"rune slice with zero rune":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any runes": pfmt.Any([]rune("Hello, Wörld!"))},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any runes":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any empty runes": pfmt.Any([]rune{})},
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty runes":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any rune slice with zero rune": pfmt.Any([]rune{rune(0)})},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"any rune slice with zero rune":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect runes": pfmt.Reflect([]rune("Hello, Wörld!"))},
			want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
			wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
			wantJSON: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect empty runes": pfmt.Reflect([]rune{})},
			want:     "[]",
			wantText: "[]",
			wantJSON: `{
			"reflect empty runes":[]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect rune slice with zero rune": pfmt.Reflect([]rune{rune(0)})},
			want:     "[0]",
			wantText: "[0]",
			wantJSON: `{
			"reflect rune slice with zero rune":[0]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []rune("Hello, Wörld!")
				return map[string]json.Marshaler{"runes pointer": pfmt.Runesp(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"runes pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []rune{}
				return map[string]json.Marshaler{"empty runes pointer": pfmt.Runesp(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty runes pointer":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil runes pointer": pfmt.Runesp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil runes pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []rune("Hello, Wörld!")
				return map[string]json.Marshaler{"any runes pointer": pfmt.Any(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []rune{}
				return map[string]json.Marshaler{"any empty runes pointer": pfmt.Any(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty runes pointer":""
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []rune("Hello, Wörld!")
				return map[string]json.Marshaler{"reflect runes pointer": pfmt.Reflect(&p)}
			}(),
			want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
			wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
			wantJSON: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := []rune{}
				return map[string]json.Marshaler{"reflect empty runes pointer": pfmt.Reflect(&p)}
			}(),
			want:     "[]",
			wantText: "[]",
			wantJSON: `{
			"reflect empty runes pointer":[]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"string": pfmt.String("Hello, Wörld!")},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"string":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty string": pfmt.String("")},
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty string":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"string with zero byte": pfmt.String(string(byte(0)))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"string with zero byte":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"strings": pfmt.Strings([]string{"Hello, Wörld!", "Hello, World!"})},
			want:     "Hello, Wörld! Hello, World!",
			wantText: "Hello, Wörld! Hello, World!",
			wantJSON: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty strings": pfmt.Strings([]string{"", ""})},
			want:     " ",
			wantText: " ",
			wantJSON: `{
			"empty strings":["",""]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"strings with zero byte": pfmt.Strings([]string{string(byte(0)), string(byte(0))})},
			want:     "\\u0000 \\u0000",
			wantText: "\\u0000 \\u0000",
			wantJSON: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without strings": pfmt.Strings(nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"without strings":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any string": pfmt.Any("Hello, Wörld!")},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any string":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any empty string": pfmt.Any("")},
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty string":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any string with zero byte": pfmt.Any(string(byte(0)))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"any string with zero byte":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect string": pfmt.Reflect("Hello, Wörld!")},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"reflect string":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect empty string": pfmt.Reflect("")},
			want:     "",
			wantText: "",
			wantJSON: `{
			"reflect empty string":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect string with zero byte": pfmt.Reflect(string(byte(0)))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"reflect string with zero byte":"\u0000"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := "Hello, Wörld!"
				return map[string]json.Marshaler{"string pointer": pfmt.Stringp(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"string pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := ""
				return map[string]json.Marshaler{"empty string pointer": pfmt.Stringp(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty string pointer":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil string pointer": pfmt.Stringp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil string pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := "Hello, Wörld!"
				return map[string]json.Marshaler{"any string pointer": pfmt.Any(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"any string pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := ""
				return map[string]json.Marshaler{"any empty string pointer": pfmt.Any(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"any empty string pointer":""
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := "Hello, Wörld!"
				return map[string]json.Marshaler{"reflect string pointer": pfmt.Reflect(&p)}
			}(),
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				p := ""
				return map[string]json.Marshaler{"reflect empty string pointer": pfmt.Reflect(&p)}
			}(),
			want:     "",
			wantText: "",
			wantJSON: `{
			"reflect empty string pointer":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"text": pfmt.Text(pfmt.String("Hello, Wörld!"))},
			want:     "Hello, Wörld!",
			wantText: "Hello, Wörld!",
			wantJSON: `{
			"text":"Hello, Wörld!"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"empty text": pfmt.Text(pfmt.String(""))},
			want:     "",
			wantText: "",
			wantJSON: `{
			"empty text":""
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"text with zero byte": pfmt.Text(pfmt.String(string(byte(0))))},
			want:     "\\u0000",
			wantText: "\\u0000",
			wantJSON: `{
			"text with zero byte":"\u0000"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint": pfmt.Uint(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint = 42
				return map[string]json.Marshaler{"uint pointer": pfmt.Uintp(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint pointer": pfmt.Uintp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint = 42
				return map[string]json.Marshaler{"any uint pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint = 42
				return map[string]json.Marshaler{"reflect uint pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint16": pfmt.Uint16(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint16":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint16": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint16":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint16": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint16":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint16 = 42
				return map[string]json.Marshaler{"uint16 pointer": pfmt.Uint16p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint16 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint16 pointer": pfmt.Uint16p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"uint16 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint16 = 42
				return map[string]json.Marshaler{"any uint16 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint16 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint16 = 42
				return map[string]json.Marshaler{"reflect uint16 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint16 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i *uint16
				return map[string]json.Marshaler{"reflect uint16 pointer to nil": pfmt.Reflect(i)}
			}(),
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"reflect uint16 pointer to nil":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint32": pfmt.Uint32(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint32":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint32": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint32":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint32": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint32":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint32 = 42
				return map[string]json.Marshaler{"uint32 pointer": pfmt.Uint32p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint32 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint32 pointer": pfmt.Uint32p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint32 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint32 = 42
				return map[string]json.Marshaler{"any uint32 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint32 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint32 = 42
				return map[string]json.Marshaler{"reflect uint32 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint32 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint64": pfmt.Uint64(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint64": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint64":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint64": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint64":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint64 = 42
				return map[string]json.Marshaler{"uint64 pointer": pfmt.Uint64p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint64 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint64 pointer": pfmt.Uint64p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint64 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint64 = 42
				return map[string]json.Marshaler{"any uint64 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint64 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint64 = 42
				return map[string]json.Marshaler{"reflect uint64 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint64 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uint8": pfmt.Uint8(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint8":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uint8": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint8":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uint8": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint8":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint8 = 42
				return map[string]json.Marshaler{"uint8 pointer": pfmt.Uint8p(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uint8 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uint8 pointer": pfmt.Uint8p(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uint8 pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint8 = 42
				return map[string]json.Marshaler{"any uint8 pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uint8 pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uint8 = 42
				return map[string]json.Marshaler{"reflect uint8 pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uint8 pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"uintptr": pfmt.Uintptr(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uintptr":42
		}`,
		},
		// FIXME: use var x uintptr = 42
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any uintptr": pfmt.Any(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uintptr":42
		}`,
		},
		// FIXME: use var x uintptr = 42
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect uintptr": pfmt.Reflect(42)},
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uintptr":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uintptr = 42
				return map[string]json.Marshaler{"uintptr pointer": pfmt.Uintptrp(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"uintptr pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"nil uintptr pointer": pfmt.Uintptrp(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"nil uintptr pointer":null
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uintptr = 42
				return map[string]json.Marshaler{"any uintptr pointer": pfmt.Any(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"any uintptr pointer":42
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var i uintptr = 42
				return map[string]json.Marshaler{"reflect uintptr pointer": pfmt.Reflect(&i)}
			}(),
			want:     "42",
			wantText: "42",
			wantJSON: `{
			"reflect uintptr pointer":42
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"time": time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)},
			want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
			wantText: "1970-01-01T00:00:00.000000042Z",
			wantJSON: `{
			"time":"1970-01-01T00:00:00.000000042Z"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"any time": pfmt.Any(time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC))},
			want:     `1970-01-01 00:00:00.000000042 +0000 UTC`,
			wantText: `1970-01-01T00:00:00.000000042Z`,
			wantJSON: `{
			"any time":"1970-01-01T00:00:00.000000042Z"
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"reflect time": pfmt.Reflect(time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC))},
			want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
			wantText: "1970-01-01 00:00:00.000000042 +0000 UTC",
			wantJSON: `{
			"reflect time":"1970-01-01T00:00:00.000000042Z"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return map[string]json.Marshaler{"time pointer": &t}
			}(),
			want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
			wantText: "1970-01-01T00:00:00.000000042Z",
			wantJSON: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				var t time.Time
				return map[string]json.Marshaler{"nil time pointer": t}
			}(),
			want:     "0001-01-01 00:00:00 +0000 UTC",
			wantText: "0001-01-01T00:00:00Z",
			wantJSON: `{
			"nil time pointer":"0001-01-01T00:00:00Z"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return map[string]json.Marshaler{"any time pointer": pfmt.Any(&t)}
			}(),
			want:     `1970-01-01 00:00:00.000000042 +0000 UTC`,
			wantText: `1970-01-01T00:00:00.000000042Z`,
			wantJSON: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
		},
		{
			line: line(),
			input: func() map[string]json.Marshaler {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return map[string]json.Marshaler{"reflect time pointer": pfmt.Reflect(&t)}
			}(),
			want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
			wantText: "1970-01-01 00:00:00.000000042 +0000 UTC",
			wantJSON: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
		},
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
			input:    map[string]json.Marshaler{"raw json": pfmt.Raw([]byte(`{"foo":"bar"}`))},
			want:     `{"foo":"bar"}`,
			wantText: `{"foo":"bar"}`,
			wantJSON: `{
			"raw json":{"foo":"bar"}
		}`,
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw malformed json object": pfmt.Raw([]byte(`xyz{"foo":"bar"}`))},
			want:      `xyz{"foo":"bar"}`,
			wantText:  `xyz{"foo":"bar"}`,
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw malformed json key/value": pfmt.Raw([]byte(`{"foo":"bar""}`))},
			want:      `{"foo":"bar""}`,
			wantText:  `{"foo":"bar""}`,
			wantError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
		},
		{
			line:      line(),
			input:     map[string]json.Marshaler{"raw json with unescaped null byte": pfmt.Raw(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...))},
			want:      "{\"foo\":\"\u0000xyz\"}",
			wantText:  "{\"foo\":\"\u0000xyz\"}",
			wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"raw nil": pfmt.Raw(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"raw nil":null
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
			input:    map[string]json.Marshaler{"any untyped nil": pfmt.Any(nil)},
			want:     "null",
			wantText: "null",
			wantJSON: `{
			"any untyped nil":null
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
