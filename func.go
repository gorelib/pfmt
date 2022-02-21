// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"errors"
	"fmt"
	"reflect"
)

// Func returns stringer/JSON/text marshaler for the function.
func Func(v interface{}) FuncV { return New().Func(v) }

// Func returns stringer/JSON/text marshaler for the function.
func (pretty Pretty) Func(v interface{}) FuncV {
	return FuncV{
		v:        v,
		prettier: pretty,
	}
}

type FuncV struct {
	v        interface{}
	prettier Pretty
}

func (v FuncV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v FuncV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Func {
		return nil, errors.New("not function")
	}
	if val.IsNil() {
		return []byte(v.prettier.nil), nil
	}

	return []byte(fmt.Sprint(v.v)), nil
}

func (v FuncV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Func {
		return nil, errors.New("not function")
	}
	if val.IsNil() {
		return []byte("null"), nil
	}

	return []byte(reflect.TypeOf(v.v).String()), nil
}
