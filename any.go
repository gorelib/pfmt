// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"encoding"
	"encoding/json"
	"time"
)

// Any returns stringer/JSON/text marshaler for any type.
func Any(v interface{}) AnyV { return New().Any(v) }

// Any returns stringer/JSON/text marshaler for any type.
func (pretty Pretty) Any(v interface{}) AnyV {
	return AnyV{
		v:        v,
		prettier: pretty,
	}
}

type AnyV struct {
	v        interface{}
	prettier Pretty
}

func (v AnyV) String() string {
	if v.v == nil {
		return v.prettier.nil
	}

	switch x := v.v.(type) {
	case bool:
		return v.prettier.Bool(x).String()

	case *bool:
		return v.prettier.Boolp(x).String()

	case []*bool:
		return v.prettier.Boolps(x).String()

	case []byte:
		return v.prettier.Bytes(x).String()

	case *[]byte:
		return v.prettier.Bytesp(x).String()

	case [][]byte:
		return v.prettier.Bytess(x).String()

	case []*[]byte:
		return v.prettier.Bytesps(x).String()

	case complex128:
		return v.prettier.Complex128(x).String()

	case *complex128:
		return v.prettier.Complex128p(x).String()

	case []complex128:
		return v.prettier.Complex128s(x).String()

	case []*complex128:
		return v.prettier.Complex128ps(x).String()

	case complex64:
		return v.prettier.Complex64(x).String()

	case *complex64:
		return v.prettier.Complex64p(x).String()

	case []complex64:
		return v.prettier.Complex64s(x).String()

	case []*complex64:
		return v.prettier.Complex64ps(x).String()

	case error:
		return v.prettier.Err(x).String()

	case *error:
		return v.prettier.Errp(x).String()

	case []error:
		return v.prettier.Errs(x).String()

	case []*error:
		return v.prettier.Errps(x).String()

	case float32:
		return v.prettier.Float32(x).String()

	case *float32:
		return v.prettier.Float32p(x).String()

	case []float32:
		return v.prettier.Float32s(x).String()

	case []*float32:
		return v.prettier.Float32ps(x).String()

	case float64:
		return v.prettier.Float64(x).String()

	case *float64:
		return v.prettier.Float64p(x).String()

	case []float64:
		return v.prettier.Float64s(x).String()

	case []*float64:
		return v.prettier.Float64ps(x).String()

	case int:
		return v.prettier.Int(x).String()

	case *int:
		return v.prettier.Intp(x).String()

	case []int:
		return v.prettier.Ints(x).String()

	case []*int:
		return v.prettier.Intps(x).String()

	case int16:
		return v.prettier.Int16(x).String()

	case *int16:
		return v.prettier.Int16p(x).String()

	case []int16:
		return v.prettier.Int16s(x).String()

	case []*int16:
		return v.prettier.Int16ps(x).String()

	case int32:
		return v.prettier.Int32(x).String()

	case *int32:
		return v.prettier.Int32p(x).String()

	case []*int32:
		return v.prettier.Int32ps(x).String()

	case int64:
		return v.prettier.Int64(x).String()

	case *int64:
		return v.prettier.Int64p(x).String()

	case []int64:
		return v.prettier.Int64s(x).String()

	case []*int64:
		return v.prettier.Int64ps(x).String()

	case int8:
		return v.prettier.Int8(x).String()

	case *int8:
		return v.prettier.Int8p(x).String()

	case []int8:
		return v.prettier.Int8s(x).String()

	case []*int8:
		return v.prettier.Int8ps(x).String()

	case []rune:
		return v.prettier.Runes(x).String()

	case *[]rune:
		return v.prettier.Runesp(x).String()

	case []*[]rune:
		return v.prettier.Runesps(x).String()

	case string:
		return v.prettier.String(x).String()

	case *string:
		return v.prettier.Stringp(x).String()

	case []string:
		return v.prettier.Strings(x).String()

	case []*string:
		return v.prettier.Stringps(x).String()

	case uint:
		return v.prettier.Uint(x).String()

	case *uint:
		return v.prettier.Uintp(x).String()

	case []uint:
		return v.prettier.Uints(x).String()

	case []*uint:
		return v.prettier.Uintps(x).String()

	case uint16:
		return v.prettier.Uint16(x).String()

	case *uint16:
		return v.prettier.Uint16p(x).String()

	case []uint16:
		return v.prettier.Uint16s(x).String()

	case []*uint16:
		return v.prettier.Uint16ps(x).String()

	case uint32:
		return v.prettier.Uint32(x).String()

	case *uint32:
		return v.prettier.Uint32p(x).String()

	case []uint32:
		return v.prettier.Uint32s(x).String()

	case []*uint32:
		return v.prettier.Uint32ps(x).String()

	case uint64:
		return v.prettier.Uint64(x).String()

	case *uint64:
		return v.prettier.Uint64p(x).String()

	case []uint64:
		return v.prettier.Uint64s(x).String()

	case []*uint64:
		return v.prettier.Uint64ps(x).String()

	case uint8:
		return v.prettier.Uint8(x).String()

	case *uint8:
		return v.prettier.Uint8p(x).String()

	case []*uint8:
		return v.prettier.Uint8ps(x).String()

	case uintptr:
		return v.prettier.Uintptr(x).String()

	case *uintptr:
		return v.prettier.Uintptrp(x).String()

	case []uintptr:
		return v.prettier.Uintptrs(x).String()

	case []*uintptr:
		return v.prettier.Uintptrps(x).String()

	case time.Time:
		return v.prettier.Time(x).String()

	case *time.Time:
		return v.prettier.Timep(x).String()

	case []time.Time:
		return v.prettier.Times(x).String()

	case []*time.Time:
		return v.prettier.Timeps(x).String()

	case time.Duration:
		return v.prettier.Duration(x).String()

	case *time.Duration:
		return v.prettier.Durationp(x).String()

	case []time.Duration:
		return v.prettier.Durations(x).String()

	case []*time.Duration:
		return v.prettier.Durationps(x).String()

	case encoding.TextMarshaler:
		return v.prettier.Text(x).String()

	case []encoding.TextMarshaler:
		return v.prettier.Texts(x).String()

	case json.Marshaler:
		b, _ := x.MarshalJSON()
		return string(b)

	case []json.Marshaler:
		return v.prettier.JSONMarshalers(x).String()
	}

	return v.prettier.Reflect(v.v).String()
}

func (v AnyV) MarshalText() ([]byte, error) {
	switch x := v.v.(type) {
	case bool:
		return v.prettier.Bool(x).MarshalText()

	case *bool:
		return v.prettier.Boolp(x).MarshalText()

	case []*bool:
		return v.prettier.Boolps(x).MarshalText()

	case []byte:
		return v.prettier.Bytes(x).MarshalText()

	case *[]byte:
		return v.prettier.Bytesp(x).MarshalText()

	case [][]byte:
		return v.prettier.Bytess(x).MarshalText()

	case []*[]byte:
		return v.prettier.Bytesps(x).MarshalText()

	case complex128:
		return v.prettier.Complex128(x).MarshalText()

	case *complex128:
		return v.prettier.Complex128p(x).MarshalText()

	case []complex128:
		return v.prettier.Complex128s(x).MarshalText()

	case []*complex128:
		return v.prettier.Complex128ps(x).MarshalText()

	case complex64:
		return v.prettier.Complex64(x).MarshalText()

	case *complex64:
		return v.prettier.Complex64p(x).MarshalText()

	case []complex64:
		return v.prettier.Complex64s(x).MarshalText()

	case []*complex64:
		return v.prettier.Complex64ps(x).MarshalText()

	case error:
		return v.prettier.Err(x).MarshalText()

	case *error:
		return v.prettier.Errp(x).MarshalText()

	case []error:
		return v.prettier.Errs(x).MarshalText()

	case []*error:
		return v.prettier.Errps(x).MarshalText()

	case float32:
		return v.prettier.Float32(x).MarshalText()

	case *float32:
		return v.prettier.Float32p(x).MarshalText()

	case []float32:
		return v.prettier.Float32s(x).MarshalText()

	case []*float32:
		return v.prettier.Float32ps(x).MarshalText()

	case float64:
		return v.prettier.Float64(x).MarshalText()

	case *float64:
		return v.prettier.Float64p(x).MarshalText()

	case []float64:
		return v.prettier.Float64s(x).MarshalText()

	case []*float64:
		return v.prettier.Float64ps(x).MarshalText()

	case int:
		return v.prettier.Int(x).MarshalText()

	case *int:
		return v.prettier.Intp(x).MarshalText()

	case []int:
		return v.prettier.Ints(x).MarshalText()

	case []*int:
		return v.prettier.Intps(x).MarshalText()

	case int16:
		return v.prettier.Int16(x).MarshalText()

	case *int16:
		return v.prettier.Int16p(x).MarshalText()

	case []int16:
		return v.prettier.Int16s(x).MarshalText()

	case []*int16:
		return v.prettier.Int16ps(x).MarshalText()

	case int32:
		return v.prettier.Int32(x).MarshalText()

	case *int32:
		return v.prettier.Int32p(x).MarshalText()

	case []*int32:
		return v.prettier.Int32ps(x).MarshalText()

	case int64:
		return v.prettier.Int64(x).MarshalText()

	case *int64:
		return v.prettier.Int64p(x).MarshalText()

	case []int64:
		return v.prettier.Int64s(x).MarshalText()

	case []*int64:
		return v.prettier.Int64ps(x).MarshalText()

	case int8:
		return v.prettier.Int8(x).MarshalText()

	case *int8:
		return v.prettier.Int8p(x).MarshalText()

	case []int8:
		return v.prettier.Int8s(x).MarshalText()

	case []*int8:
		return v.prettier.Int8ps(x).MarshalText()

	case []rune:
		return v.prettier.Runes(x).MarshalText()

	case *[]rune:
		return v.prettier.Runesp(x).MarshalText()

	case []*[]rune:
		return v.prettier.Runesps(x).MarshalText()

	case string:
		return v.prettier.String(x).MarshalText()

	case *string:
		return v.prettier.Stringp(x).MarshalText()

	case []string:
		return v.prettier.Strings(x).MarshalText()

	case []*string:
		return v.prettier.Stringps(x).MarshalText()

	case uint:
		return v.prettier.Uint(x).MarshalText()

	case *uint:
		return v.prettier.Uintp(x).MarshalText()

	case []uint:
		return v.prettier.Uints(x).MarshalText()

	case []*uint:
		return v.prettier.Uintps(x).MarshalText()

	case uint16:
		return v.prettier.Uint16(x).MarshalText()

	case *uint16:
		return v.prettier.Uint16p(x).MarshalText()

	case []uint16:
		return v.prettier.Uint16s(x).MarshalText()

	case []*uint16:
		return v.prettier.Uint16ps(x).MarshalText()

	case uint32:
		return v.prettier.Uint32(x).MarshalText()

	case *uint32:
		return v.prettier.Uint32p(x).MarshalText()

	case []uint32:
		return v.prettier.Uint32s(x).MarshalText()

	case []*uint32:
		return v.prettier.Uint32ps(x).MarshalText()

	case uint64:
		return v.prettier.Uint64(x).MarshalText()

	case *uint64:
		return v.prettier.Uint64p(x).MarshalText()

	case []uint64:
		return v.prettier.Uint64s(x).MarshalText()

	case []*uint64:
		return v.prettier.Uint64ps(x).MarshalText()

	case uint8:
		return v.prettier.Uint8(x).MarshalText()

	case *uint8:
		return v.prettier.Uint8p(x).MarshalText()

	case []*uint8:
		return v.prettier.Uint8ps(x).MarshalText()

	case uintptr:
		return v.prettier.Uintptr(x).MarshalText()

	case *uintptr:
		return v.prettier.Uintptrp(x).MarshalText()

	case []uintptr:
		return v.prettier.Uintptrs(x).MarshalText()

	case []*uintptr:
		return v.prettier.Uintptrps(x).MarshalText()

	case time.Time:
		return v.prettier.Time(x).MarshalText()

	case *time.Time:
		return v.prettier.Timep(x).MarshalText()

	case []time.Time:
		return v.prettier.Times(x).MarshalText()

	case []*time.Time:
		return v.prettier.Timeps(x).MarshalText()

	case time.Duration:
		return v.prettier.Duration(x).MarshalText()

	case *time.Duration:
		return v.prettier.Durationp(x).MarshalText()

	case []time.Duration:
		return v.prettier.Durations(x).MarshalText()

	case []*time.Duration:
		return v.prettier.Durationps(x).MarshalText()

	case encoding.TextMarshaler:
		return v.prettier.Text(x).MarshalText()

	case []encoding.TextMarshaler:
		return v.prettier.Texts(x).MarshalText()

	case json.Marshaler:
		return x.MarshalJSON()

	case []json.Marshaler:
		return v.prettier.JSONMarshalers(x).MarshalText()

	default:
		return v.prettier.Reflect(x).MarshalText()
	}
}

func (v AnyV) MarshalJSON() ([]byte, error) {
	switch x := v.v.(type) {
	case bool:
		return v.prettier.Bool(x).MarshalJSON()

	case *bool:
		return v.prettier.Boolp(x).MarshalJSON()

	case []*bool:
		return v.prettier.Boolps(x).MarshalJSON()

	case []byte:
		return v.prettier.Bytes(x).MarshalJSON()

	case *[]byte:
		return v.prettier.Bytesp(x).MarshalJSON()

	case [][]byte:
		return v.prettier.Bytess(x).MarshalJSON()

	case []*[]byte:
		return v.prettier.Bytesps(x).MarshalJSON()

	case complex128:
		return v.prettier.Complex128(x).MarshalJSON()

	case *complex128:
		return v.prettier.Complex128p(x).MarshalJSON()

	case []complex128:
		return v.prettier.Complex128s(x).MarshalJSON()

	case []*complex128:
		return v.prettier.Complex128ps(x).MarshalJSON()

	case complex64:
		return v.prettier.Complex64(x).MarshalJSON()

	case *complex64:
		return v.prettier.Complex64p(x).MarshalJSON()

	case []complex64:
		return v.prettier.Complex64s(x).MarshalJSON()

	case []*complex64:
		return v.prettier.Complex64ps(x).MarshalJSON()

	case error:
		return v.prettier.Err(x).MarshalJSON()

	case *error:
		return v.prettier.Errp(x).MarshalJSON()

	case []error:
		return v.prettier.Errs(x).MarshalJSON()

	case []*error:
		return v.prettier.Errps(x).MarshalJSON()

	case float32:
		return v.prettier.Float32(x).MarshalJSON()

	case *float32:
		return v.prettier.Float32p(x).MarshalJSON()

	case []float32:
		return v.prettier.Float32s(x).MarshalJSON()

	case []*float32:
		return v.prettier.Float32ps(x).MarshalJSON()

	case float64:
		return v.prettier.Float64(x).MarshalJSON()

	case *float64:
		return v.prettier.Float64p(x).MarshalJSON()

	case []float64:
		return v.prettier.Float64s(x).MarshalJSON()

	case []*float64:
		return v.prettier.Float64ps(x).MarshalJSON()

	case int:
		return v.prettier.Int(x).MarshalJSON()

	case *int:
		return v.prettier.Intp(x).MarshalJSON()

	case []int:
		return v.prettier.Ints(x).MarshalJSON()

	case []*int:
		return v.prettier.Intps(x).MarshalJSON()

	case int16:
		return v.prettier.Int16(x).MarshalJSON()

	case *int16:
		return v.prettier.Int16p(x).MarshalJSON()

	case []int16:
		return v.prettier.Int16s(x).MarshalJSON()

	case []*int16:
		return v.prettier.Int16ps(x).MarshalJSON()

	case int32:
		return v.prettier.Int32(x).MarshalJSON()

	case *int32:
		return v.prettier.Int32p(x).MarshalJSON()

	case []*int32:
		return v.prettier.Int32ps(x).MarshalJSON()

	case int64:
		return v.prettier.Int64(x).MarshalJSON()

	case *int64:
		return v.prettier.Int64p(x).MarshalJSON()

	case []int64:
		return v.prettier.Int64s(x).MarshalJSON()

	case []*int64:
		return v.prettier.Int64ps(x).MarshalJSON()

	case int8:
		return v.prettier.Int8(x).MarshalJSON()

	case *int8:
		return v.prettier.Int8p(x).MarshalJSON()

	case []int8:
		return v.prettier.Int8s(x).MarshalJSON()

	case []*int8:
		return v.prettier.Int8ps(x).MarshalJSON()

	case []rune:
		return v.prettier.Runes(x).MarshalJSON()

	case *[]rune:
		return v.prettier.Runesp(x).MarshalJSON()

	case []*[]rune:
		return v.prettier.Runesps(x).MarshalJSON()

	case string:
		return v.prettier.String(x).MarshalJSON()

	case *string:
		return v.prettier.Stringp(x).MarshalJSON()

	case []string:
		return v.prettier.Strings(x).MarshalJSON()

	case []*string:
		return v.prettier.Stringps(x).MarshalJSON()

	case uint:
		return v.prettier.Uint(x).MarshalJSON()

	case *uint:
		return v.prettier.Uintp(x).MarshalJSON()

	case []uint:
		return v.prettier.Uints(x).MarshalJSON()

	case []*uint:
		return v.prettier.Uintps(x).MarshalJSON()

	case uint16:
		return v.prettier.Uint16(x).MarshalJSON()

	case *uint16:
		return v.prettier.Uint16p(x).MarshalJSON()

	case []uint16:
		return v.prettier.Uint16s(x).MarshalJSON()

	case []*uint16:
		return v.prettier.Uint16ps(x).MarshalJSON()

	case uint32:
		return v.prettier.Uint32(x).MarshalJSON()

	case *uint32:
		return v.prettier.Uint32p(x).MarshalJSON()

	case []uint32:
		return v.prettier.Uint32s(x).MarshalJSON()

	case []*uint32:
		return v.prettier.Uint32ps(x).MarshalJSON()

	case uint64:
		return v.prettier.Uint64(x).MarshalJSON()

	case *uint64:
		return v.prettier.Uint64p(x).MarshalJSON()

	case []uint64:
		return v.prettier.Uint64s(x).MarshalJSON()

	case []*uint64:
		return v.prettier.Uint64ps(x).MarshalJSON()

	case uint8:
		return v.prettier.Uint8(x).MarshalJSON()

	case *uint8:
		return v.prettier.Uint8p(x).MarshalJSON()

	case []*uint8:
		return v.prettier.Uint8ps(x).MarshalJSON()

	case uintptr:
		return v.prettier.Uintptr(x).MarshalJSON()

	case *uintptr:
		return v.prettier.Uintptrp(x).MarshalJSON()

	case []uintptr:
		return v.prettier.Uintptrs(x).MarshalJSON()

	case []*uintptr:
		return v.prettier.Uintptrps(x).MarshalJSON()

	case time.Time:
		return v.prettier.Time(x).MarshalJSON()

	case *time.Time:
		return v.prettier.Timep(x).MarshalJSON()

	case []time.Time:
		return v.prettier.Times(x).MarshalJSON()

	case []*time.Time:
		return v.prettier.Timeps(x).MarshalJSON()

	case time.Duration:
		return v.prettier.Duration(x).MarshalJSON()

	case *time.Duration:
		return v.prettier.Durationp(x).MarshalJSON()

	case []time.Duration:
		return v.prettier.Durations(x).MarshalJSON()

	case []*time.Duration:
		return v.prettier.Durationps(x).MarshalJSON()

	case encoding.TextMarshaler:
		return v.prettier.Text(x).MarshalJSON()

	case []encoding.TextMarshaler:
		return v.prettier.Texts(x).MarshalJSON()

	case json.Marshaler:
		return x.MarshalJSON()

	case []json.Marshaler:
		return v.prettier.JSONMarshalers(x).MarshalJSON()

	default:
		return v.prettier.Reflect(x).MarshalJSON()
	}
}
