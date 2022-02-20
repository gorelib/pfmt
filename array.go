// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"errors"
	"reflect"
)

// Array returns stringer/JSON/text marshaler for the array type.
func Array(v interface{}) ArrayV { return New().Array(v) }

// Array returns stringer/JSON/text marshaler for the array type.
func (pretty Pretty) Array(v interface{}) ArrayV {
	return ArrayV{
		v:        v,
		prettier: pretty,
	}
}

type ArrayV struct {
	v        interface{}
	prettier Pretty
}

func (v ArrayV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v ArrayV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Array {
		return nil, errors.New("not array")
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

func (v ArrayV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
