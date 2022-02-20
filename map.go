// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"errors"
	"fmt"
	"reflect"
)

// Map returns stringer/JSON/text marshaler for the map.
func Map(v interface{}) MapV { return New().Map(v) }

// Map returns stringer/JSON/text marshaler for the map.
func (pretty Pretty) Map(v interface{}) MapV {
	return MapV{
		v:        v,
		prettier: pretty,
	}
}

type MapV struct {
	v        interface{}
	prettier Pretty
}

func (v MapV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v MapV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Map {
		return nil, errors.New("not map")
	}
	if val.IsNil() {
		return []byte(v.prettier.nil), nil
	}

	return []byte(fmt.Sprint(v.v)), nil
}

func (v MapV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
