// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
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

	num := val.Len()
	values := make([]interface{}, 0, num)

	for i := 0; i < val.Len(); i++ {
		field := val.Index(i)
		if field.CanInterface() {
			values = append(values, field.Interface())
		} else {
			num--
		}
	}

	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	_, err := buf.WriteString("[")
	if err != nil {
		return nil, err
	}

	for i := 0; i < num; i++ {
		if i != 0 {
			_, err = buf.WriteString(v.prettier.separator)
			if err != nil {
				return nil, err
			}
		}
		_, err = buf.WriteString(v.prettier.Reflect(values[i]).String())
		if err != nil {
			return nil, err
		}
	}

	_, err = buf.WriteString("]")
	if err != nil {
		return nil, err
	}

	p := make([]byte, len(buf.Bytes()))
	copy(p, buf.Bytes())

	return p, nil
}

func (v SliceV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
