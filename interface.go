// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"errors"
	"fmt"
	"reflect"
)

// Interface returns stringer/JSON/text marshaler for the interface.
func Interface(v interface{}) InterfaceV { return New().Interface(v) }

// Interface returns stringer/JSON/text marshaler for the interface.
func (pretty Pretty) Interface(v interface{}) InterfaceV {
	return InterfaceV{
		v:        v,
		prettier: pretty,
	}
}

type InterfaceV struct {
	v        interface{}
	prettier Pretty
}

func (v InterfaceV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v InterfaceV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Interface {
		return nil, errors.New("not interface")
	}
	if val.IsNil() {
		return []byte(v.prettier.nil), nil
	}

	return []byte(fmt.Sprint(v.v)), nil
}

func (v InterfaceV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
