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
)

// Reflect returns stringer/JSON/text marshaler uses reflection.
func Reflect(v interface{}) ReflectV { return ReflectV{v: v} }

type ReflectV struct{ v interface{} }

func (v ReflectV) String() string {
	if v.v == nil {
		return "null"
	}

	val := reflect.ValueOf(v.v)

	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if val.IsNil() {
			return "null"

		} else if val.Kind() == reflect.Ptr {
			return ReflectV{v: val.Elem().Interface()}.String()

		} else if val.Kind() == reflect.Slice && val.Type().Elem().Kind() == reflect.Uint8 { // Byte slice.
			buf := bufPool.Get().(*bytes.Buffer)
			buf.Reset()
			defer bufPool.Put(buf)

			p := val.Bytes()
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			_, _ = enc.Write(p)
			enc.Close()

			return buf.String()
		}
	}

	return fmt.Sprint(v.v)
}

func (v ReflectV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v ReflectV) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.v)
}
