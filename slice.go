// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
)

// Slice returns stringer/JSON/text marshaler for the slice type.
func Slice(v interface{}) SliceV { return New().Slice(v) }

// Slice returns stringer/JSON/text marshaler for the slice type.
func (pretty Pretty) Slice(v interface{}) SliceV {
	return SliceV{
		v:        v,
		prettier: pretty,
	}
}

type SliceV struct {
	v        interface{}
	prettier Pretty
}

func (v SliceV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v SliceV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Slice {
		return nil, errors.New("not slice")
	}
	if val.IsNil() {
		return []byte(v.prettier.nil), nil
	}

	num := val.Len()
	values := make([]interface{}, 0, num)

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		if elem.CanInterface() {
			values = append(values, elem.Interface())
		} else {
			num--
		}
	}

	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	buf.WriteString("[")

	for i := 0; i < num; i++ {
		if i != 0 {
			buf.WriteString(v.prettier.separator)
		}
		buf.WriteString(v.prettier.Reflect(values[i]).String())
	}

	buf.WriteString("]")

	p := make([]byte, len(buf.Bytes()))
	copy(p, buf.Bytes())

	return p, nil
}

func (v SliceV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Slice {
		return nil, errors.New("not slice")
	}
	if val.IsNil() {
		return []byte("null"), nil
	}

	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	buf.WriteString("[")

	num := 0

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		if !elem.CanInterface() {
			continue
		}

		if num != 0 {
			buf.WriteString(",")
		}

		num++
		it := elem.Interface()
		var j []byte

		if it == nil {
			return []byte("null"), nil

		} else if marsh, ok := it.(json.Marshaler); ok {
			var err error
			j, err = marsh.MarshalJSON()
			if err != nil {
				return nil, err
			}

		} else {
			var err error
			j, err := v.prettier.Reflect(it).MarshalJSON()
			if err != nil {
				return nil, err
			}
			buf.Write(j)
		}

		buf.Write(j)
	}

	buf.WriteString("]")

	p := make([]byte, len(buf.Bytes()))
	copy(p, buf.Bytes())

	return p, nil
}
