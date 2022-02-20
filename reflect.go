// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"unicode"
)

// Reflect returns stringer/JSON/text marshaler uses reflection.
func Reflect(v interface{}) Prettier { return New().Reflect(v) }

// Reflect returns stringer/JSON/text marshaler uses reflection.
func (pretty Pretty) Reflect(v interface{}) Prettier {
	pretty.stack--
	if pretty.stack == 0 {
		return pretty.Dummie(v)
	}
	return ReflectV{
		v:        v,
		prettier: pretty,
	}
}

type ReflectV struct {
	v        interface{}
	prettier Pretty
}

func (v ReflectV) String() string {
	if v.v == nil {
		return v.prettier.nil
	}

	val := reflect.ValueOf(v.v)

	switch val.Kind() {
	case reflect.Bool:
		return v.prettier.Bool(val.Bool()).String()

	case reflect.Int:
		return v.prettier.Int(int(val.Int())).String()

	case reflect.Int8:
		return v.prettier.Int8(int8(val.Int())).String()

	case reflect.Int16:
		return v.prettier.Int16(int16(val.Int())).String()

	case reflect.Int32:
		return v.prettier.Int32(int32(val.Int())).String()

	case reflect.Int64:
		return v.prettier.Int64(int64(val.Int())).String()

	case reflect.Uint:
		return v.prettier.Uint(uint(val.Uint())).String()

	case reflect.Uint8:
		return v.prettier.Uint8(uint8(val.Uint())).String()

	case reflect.Uint16:
		return v.prettier.Uint16(uint16(val.Uint())).String()

	case reflect.Uint32:
		return v.prettier.Uint32(uint32(val.Uint())).String()

	case reflect.Uint64:
		return v.prettier.Uint64(uint64(val.Uint())).String()

	case reflect.Uintptr:
		return v.prettier.Uintptr(uintptr(val.Uint())).String()

	case reflect.Float32:
		return v.prettier.Float32(float32(val.Float())).String()

	case reflect.Float64:
		return v.prettier.Float64(float64(val.Float())).String()

	case reflect.Complex64:
		return v.prettier.Complex64(complex64(val.Complex())).String()

	case reflect.Complex128:
		return v.prettier.Complex128(complex128(val.Complex())).String()

	case reflect.Array:
		return v.prettier.Array(v.v).String()

	case reflect.Chan:
		return v.prettier.Chan(v.v).String()

	case reflect.Func:
		return v.prettier.Func(v.v).String()

	case reflect.Interface:
		return v.prettier.Interface(v.v).String()

	case reflect.Map:
		return v.prettier.Map(v.v).String()

	case reflect.Ptr:
		if val.IsNil() {
			return v.prettier.nil
		}
		return v.prettier.Reflect(val.Elem().Interface()).String()

	case reflect.Slice:
		if val.IsNil() {
			return v.prettier.nil

		} else if val.Kind() == reflect.Slice && val.Type().Elem().Kind() == reflect.Uint8 { // Byte slice.
			buf := pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer pool.Put(buf)

			p := val.Bytes()
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			_, _ = enc.Write(p)
			enc.Close()

			return buf.String()
		}

		return v.prettier.Slice(v.v).String()

	case reflect.String:
		return v.prettier.String(val.String()).String()

	case reflect.Struct:
		if val.Type().Name() == "" || !unicode.IsUpper(rune(val.Type().Name()[0])) {
			return fmt.Sprint(v.v)
		}
		return v.prettier.Struct(v.v).String()

	case reflect.UnsafePointer:
		return fmt.Sprint(v.v)
	}

	return fmt.Sprint(v.v)
}

func (v ReflectV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v ReflectV) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.v)
}
