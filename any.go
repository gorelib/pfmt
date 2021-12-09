// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"encoding"
	"encoding/json"
	"time"
)

// Any returns stringer/JSON/text marshaler for any type.
func Any(v interface{}) AnyV { return AnyV{V: v} }

type AnyV struct{ V interface{} }

func (v AnyV) String() string {
	switch x := v.V.(type) {
	case bool:
		return Bool(x).String()
	case *bool:
		return Boolp(x).String()
	case []*bool:
		return Boolps(x...).String()
	case []byte:
		return Bytes(x...).String()
	case *[]byte:
		return Bytesp(x).String()
	case [][]byte:
		return Bytess(x...).String()
	case []*[]byte:
		return Bytesps(x...).String()
	case complex128:
		return Complex128(x).String()
	case *complex128:
		return Complex128p(x).String()
	case []complex128:
		return Complex128s(x...).String()
	case []*complex128:
		return Complex128ps(x...).String()
	case complex64:
		return Complex64(x).String()
	case *complex64:
		return Complex64p(x).String()
	case []complex64:
		return Complex64s(x...).String()
	case []*complex64:
		return Complex64ps(x...).String()
	case error:
		return Error(x).String()
	case *error:
		return Errorp(x).String()
	case []error:
		return Errors(x...).String()
	case []*error:
		return Errorps(x...).String()
	case float32:
		return Float32(x).String()
	case *float32:
		return Float32p(x).String()
	case []float32:
		return Float32s(x...).String()
	case []*float32:
		return Float32ps(x...).String()
	case float64:
		return Float64(x).String()
	case *float64:
		return Float64p(x).String()
	case []float64:
		return Float64s(x...).String()
	case []*float64:
		return Float64ps(x...).String()
	case int:
		return Int(x).String()
	case *int:
		return Intp(x).String()
	case []int:
		return Ints(x...).String()
	case []*int:
		return Intps(x...).String()
	case int16:
		return Int16(x).String()
	case *int16:
		return Int16p(x).String()
	case []int16:
		return Int16s(x...).String()
	case []*int16:
		return Int16ps(x...).String()
	case int32:
		return Int32(x).String()
	case *int32:
		return Int32p(x).String()
	case []*int32:
		return Int32ps(x...).String()
	case int64:
		return Int64(x).String()
	case *int64:
		return Int64p(x).String()
	case []int64:
		return Int64s(x...).String()
	case []*int64:
		return Int64ps(x...).String()
	case int8:
		return Int8(x).String()
	case *int8:
		return Int8p(x).String()
	case []int8:
		return Int8s(x...).String()
	case []*int8:
		return Int8ps(x...).String()
	case []rune:
		return Runes(x...).String()
	case *[]rune:
		return Runesp(x).String()
	case []*[]rune:
		return Runesps(x...).String()
	case string:
		return String(x).String()
	case *string:
		return Stringp(x).String()
	case []string:
		return Strings(x...).String()
	case []*string:
		return Stringps(x...).String()
	case uint:
		return Uint(x).String()
	case *uint:
		return Uintp(x).String()
	case []uint:
		return Uints(x...).String()
	case []*uint:
		return Uintps(x...).String()
	case uint16:
		return Uint16(x).String()
	case *uint16:
		return Uint16p(x).String()
	case []uint16:
		return Uint16s(x...).String()
	case []*uint16:
		return Uint16ps(x...).String()
	case uint32:
		return Uint32(x).String()
	case *uint32:
		return Uint32p(x).String()
	case []uint32:
		return Uint32s(x...).String()
	case []*uint32:
		return Uint32ps(x...).String()
	case uint64:
		return Uint64(x).String()
	case *uint64:
		return Uint64p(x).String()
	case []uint64:
		return Uint64s(x...).String()
	case []*uint64:
		return Uint64ps(x...).String()
	case uint8:
		return Uint8(x).String()
	case *uint8:
		return Uint8p(x).String()
	case []*uint8:
		return Uint8ps(x...).String()
	case uintptr:
		return Uintptr(x).String()
	case *uintptr:
		return Uintptrp(x).String()
	case []uintptr:
		return Uintptrs(x...).String()
	case []*uintptr:
		return Uintptrps(x...).String()
	case time.Time:
		return Time(x).String()
	case *time.Time:
		return Timep(x).String()
	case []time.Time:
		return Times(x...).String()
	case []*time.Time:
		return Timeps(x...).String()
	case time.Duration:
		return Duration(x).String()
	case *time.Duration:
		return Durationp(x).String()
	case []time.Duration:
		return Durations(x...).String()
	case []*time.Duration:
		return Durationps(x...).String()
	case encoding.TextMarshaler:
		return Text(x).String()
	case []encoding.TextMarshaler:
		return Texts(x...).String()
	case json.Marshaler:
		b, _ := x.MarshalJSON()
		return string(b)
	case []json.Marshaler:
		return JSONMarshalers(x...).String()
	default:
		return Reflect(x).String()
	}
}

func (v AnyV) MarshalText() ([]byte, error) {
	switch x := v.V.(type) {
	case bool:
		return Bool(x).MarshalText()
	case *bool:
		return Boolp(x).MarshalText()
	case []*bool:
		return Boolps(x...).MarshalText()
	case []byte:
		return Bytes(x...).MarshalText()
	case *[]byte:
		return Bytesp(x).MarshalText()
	case [][]byte:
		return Bytess(x...).MarshalText()
	case []*[]byte:
		return Bytesps(x...).MarshalText()
	case complex128:
		return Complex128(x).MarshalText()
	case *complex128:
		return Complex128p(x).MarshalText()
	case []complex128:
		return Complex128s(x...).MarshalText()
	case []*complex128:
		return Complex128ps(x...).MarshalText()
	case complex64:
		return Complex64(x).MarshalText()
	case *complex64:
		return Complex64p(x).MarshalText()
	case []complex64:
		return Complex64s(x...).MarshalText()
	case []*complex64:
		return Complex64ps(x...).MarshalText()
	case error:
		return Error(x).MarshalText()
	case *error:
		return Errorp(x).MarshalText()
	case []error:
		return Errors(x...).MarshalText()
	case []*error:
		return Errorps(x...).MarshalText()
	case float32:
		return Float32(x).MarshalText()
	case *float32:
		return Float32p(x).MarshalText()
	case []float32:
		return Float32s(x...).MarshalText()
	case []*float32:
		return Float32ps(x...).MarshalText()
	case float64:
		return Float64(x).MarshalText()
	case *float64:
		return Float64p(x).MarshalText()
	case []float64:
		return Float64s(x...).MarshalText()
	case []*float64:
		return Float64ps(x...).MarshalText()
	case int:
		return Int(x).MarshalText()
	case *int:
		return Intp(x).MarshalText()
	case []int:
		return Ints(x...).MarshalText()
	case []*int:
		return Intps(x...).MarshalText()
	case int16:
		return Int16(x).MarshalText()
	case *int16:
		return Int16p(x).MarshalText()
	case []int16:
		return Int16s(x...).MarshalText()
	case []*int16:
		return Int16ps(x...).MarshalText()
	case int32:
		return Int32(x).MarshalText()
	case *int32:
		return Int32p(x).MarshalText()
	case []*int32:
		return Int32ps(x...).MarshalText()
	case int64:
		return Int64(x).MarshalText()
	case *int64:
		return Int64p(x).MarshalText()
	case []int64:
		return Int64s(x...).MarshalText()
	case []*int64:
		return Int64ps(x...).MarshalText()
	case int8:
		return Int8(x).MarshalText()
	case *int8:
		return Int8p(x).MarshalText()
	case []int8:
		return Int8s(x...).MarshalText()
	case []*int8:
		return Int8ps(x...).MarshalText()
	case []rune:
		return Runes(x...).MarshalText()
	case *[]rune:
		return Runesp(x).MarshalText()
	case []*[]rune:
		return Runesps(x...).MarshalText()
	case string:
		return String(x).MarshalText()
	case *string:
		return Stringp(x).MarshalText()
	case []string:
		return Strings(x...).MarshalText()
	case []*string:
		return Stringps(x...).MarshalText()
	case uint:
		return Uint(x).MarshalText()
	case *uint:
		return Uintp(x).MarshalText()
	case []uint:
		return Uints(x...).MarshalText()
	case []*uint:
		return Uintps(x...).MarshalText()
	case uint16:
		return Uint16(x).MarshalText()
	case *uint16:
		return Uint16p(x).MarshalText()
	case []uint16:
		return Uint16s(x...).MarshalText()
	case []*uint16:
		return Uint16ps(x...).MarshalText()
	case uint32:
		return Uint32(x).MarshalText()
	case *uint32:
		return Uint32p(x).MarshalText()
	case []uint32:
		return Uint32s(x...).MarshalText()
	case []*uint32:
		return Uint32ps(x...).MarshalText()
	case uint64:
		return Uint64(x).MarshalText()
	case *uint64:
		return Uint64p(x).MarshalText()
	case []uint64:
		return Uint64s(x...).MarshalText()
	case []*uint64:
		return Uint64ps(x...).MarshalText()
	case uint8:
		return Uint8(x).MarshalText()
	case *uint8:
		return Uint8p(x).MarshalText()
	case []*uint8:
		return Uint8ps(x...).MarshalText()
	case uintptr:
		return Uintptr(x).MarshalText()
	case *uintptr:
		return Uintptrp(x).MarshalText()
	case []uintptr:
		return Uintptrs(x...).MarshalText()
	case []*uintptr:
		return Uintptrps(x...).MarshalText()
	case time.Time:
		return Time(x).MarshalText()
	case *time.Time:
		return Timep(x).MarshalText()
	case []time.Time:
		return Times(x...).MarshalText()
	case []*time.Time:
		return Timeps(x...).MarshalText()
	case time.Duration:
		return Duration(x).MarshalText()
	case *time.Duration:
		return Durationp(x).MarshalText()
	case []time.Duration:
		return Durations(x...).MarshalText()
	case []*time.Duration:
		return Durationps(x...).MarshalText()
	case encoding.TextMarshaler:
		return Text(x).MarshalText()
	case []encoding.TextMarshaler:
		return Texts(x...).MarshalText()
	case json.Marshaler:
		return x.MarshalJSON()
	case []json.Marshaler:
		return JSONMarshalers(x...).MarshalText()
	default:
		return Reflect(x).MarshalText()
	}
}

func (v AnyV) MarshalJSON() ([]byte, error) {
	switch x := v.V.(type) {
	case bool:
		return Bool(x).MarshalJSON()
	case *bool:
		return Boolp(x).MarshalJSON()
	case []*bool:
		return Boolps(x...).MarshalJSON()
	case []byte:
		return Bytes(x...).MarshalJSON()
	case *[]byte:
		return Bytesp(x).MarshalJSON()
	case [][]byte:
		return Bytess(x...).MarshalJSON()
	case []*[]byte:
		return Bytesps(x...).MarshalJSON()
	case complex128:
		return Complex128(x).MarshalJSON()
	case *complex128:
		return Complex128p(x).MarshalJSON()
	case []complex128:
		return Complex128s(x...).MarshalJSON()
	case []*complex128:
		return Complex128ps(x...).MarshalJSON()
	case complex64:
		return Complex64(x).MarshalJSON()
	case *complex64:
		return Complex64p(x).MarshalJSON()
	case []complex64:
		return Complex64s(x...).MarshalJSON()
	case []*complex64:
		return Complex64ps(x...).MarshalJSON()
	case error:
		return Error(x).MarshalJSON()
	case *error:
		return Errorp(x).MarshalJSON()
	case []error:
		return Errors(x...).MarshalJSON()
	case []*error:
		return Errorps(x...).MarshalJSON()
	case float32:
		return Float32(x).MarshalJSON()
	case *float32:
		return Float32p(x).MarshalJSON()
	case []float32:
		return Float32s(x...).MarshalJSON()
	case []*float32:
		return Float32ps(x...).MarshalJSON()
	case float64:
		return Float64(x).MarshalJSON()
	case *float64:
		return Float64p(x).MarshalJSON()
	case []float64:
		return Float64s(x...).MarshalJSON()
	case []*float64:
		return Float64ps(x...).MarshalJSON()
	case int:
		return Int(x).MarshalJSON()
	case *int:
		return Intp(x).MarshalJSON()
	case []int:
		return Ints(x...).MarshalJSON()
	case []*int:
		return Intps(x...).MarshalJSON()
	case int16:
		return Int16(x).MarshalJSON()
	case *int16:
		return Int16p(x).MarshalJSON()
	case []int16:
		return Int16s(x...).MarshalJSON()
	case []*int16:
		return Int16ps(x...).MarshalJSON()
	case int32:
		return Int32(x).MarshalJSON()
	case *int32:
		return Int32p(x).MarshalJSON()
	case []*int32:
		return Int32ps(x...).MarshalJSON()
	case int64:
		return Int64(x).MarshalJSON()
	case *int64:
		return Int64p(x).MarshalJSON()
	case []int64:
		return Int64s(x...).MarshalJSON()
	case []*int64:
		return Int64ps(x...).MarshalJSON()
	case int8:
		return Int8(x).MarshalJSON()
	case *int8:
		return Int8p(x).MarshalJSON()
	case []int8:
		return Int8s(x...).MarshalJSON()
	case []*int8:
		return Int8ps(x...).MarshalJSON()
	case []rune:
		return Runes(x...).MarshalJSON()
	case *[]rune:
		return Runesp(x).MarshalJSON()
	case []*[]rune:
		return Runesps(x...).MarshalJSON()
	case string:
		return String(x).MarshalJSON()
	case *string:
		return Stringp(x).MarshalJSON()
	case []string:
		return Strings(x...).MarshalJSON()
	case []*string:
		return Stringps(x...).MarshalJSON()
	case uint:
		return Uint(x).MarshalJSON()
	case *uint:
		return Uintp(x).MarshalJSON()
	case []uint:
		return Uints(x...).MarshalJSON()
	case []*uint:
		return Uintps(x...).MarshalJSON()
	case uint16:
		return Uint16(x).MarshalJSON()
	case *uint16:
		return Uint16p(x).MarshalJSON()
	case []uint16:
		return Uint16s(x...).MarshalJSON()
	case []*uint16:
		return Uint16ps(x...).MarshalJSON()
	case uint32:
		return Uint32(x).MarshalJSON()
	case *uint32:
		return Uint32p(x).MarshalJSON()
	case []uint32:
		return Uint32s(x...).MarshalJSON()
	case []*uint32:
		return Uint32ps(x...).MarshalJSON()
	case uint64:
		return Uint64(x).MarshalJSON()
	case *uint64:
		return Uint64p(x).MarshalJSON()
	case []uint64:
		return Uint64s(x...).MarshalJSON()
	case []*uint64:
		return Uint64ps(x...).MarshalJSON()
	case uint8:
		return Uint8(x).MarshalJSON()
	case *uint8:
		return Uint8p(x).MarshalJSON()
	case []*uint8:
		return Uint8ps(x...).MarshalJSON()
	case uintptr:
		return Uintptr(x).MarshalJSON()
	case *uintptr:
		return Uintptrp(x).MarshalJSON()
	case []uintptr:
		return Uintptrs(x...).MarshalJSON()
	case []*uintptr:
		return Uintptrps(x...).MarshalJSON()
	case time.Time:
		return Time(x).MarshalJSON()
	case *time.Time:
		return Timep(x).MarshalJSON()
	case []time.Time:
		return Times(x...).MarshalJSON()
	case []*time.Time:
		return Timeps(x...).MarshalJSON()
	case time.Duration:
		return Duration(x).MarshalJSON()
	case *time.Duration:
		return Durationp(x).MarshalJSON()
	case []time.Duration:
		return Durations(x...).MarshalJSON()
	case []*time.Duration:
		return Durationps(x...).MarshalJSON()
	case encoding.TextMarshaler:
		return Text(x).MarshalJSON()
	case []encoding.TextMarshaler:
		return Texts(x...).MarshalJSON()
	case json.Marshaler:
		return x.MarshalJSON()
	case []json.Marshaler:
		return JSONMarshalers(x...).MarshalJSON()
	default:
		return Reflect(x).MarshalJSON()
	}
}