// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Struct returns stringer/JSON/text marshaler for the struct type.
func Struct(v interface{}) StructV { return New().Struct(v) }

// Struct returns stringer/JSON/text marshaler for the struct type.
func (pretty Pretty) Struct(v interface{}) StructV {
	return StructV{
		v:        v,
		prettier: pretty,
	}
}

type StructV struct {
	v        interface{}
	prettier Pretty
}

func (v StructV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v StructV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Struct {
		return nil, errors.New("not struct")
	}

	num := val.NumField()
	names := make([]string, 0, num)
	values := make([]interface{}, 0, num)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if !field.CanInterface() {
			num--
			continue
		}

		names = append(names, val.Type().Field(i).Name)
		values = append(values, field.Interface())
	}

	if num == 0 {
		return []byte(fmt.Sprint(v.v)), nil
	}

	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	_, err := buf.WriteString(val.Type().String() + "{")
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
		_, err = buf.WriteString(names[i] + ":")
		if err != nil {
			return nil, err
		}
		_, err = buf.WriteString(v.prettier.Reflect(values[i]).String())
		if err != nil {
			return nil, err
		}
	}

	_, err = buf.WriteString("}")
	if err != nil {
		return nil, err
	}

	p := make([]byte, len(buf.Bytes()))
	copy(p, buf.Bytes())

	return p, nil
}

func (v StructV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Struct {
		return nil, errors.New("not struct")
	}

	if marsh, ok := v.v.(json.Marshaler); ok {
		return marsh.MarshalJSON()
	}

	return json.Marshal(v.v)
}
