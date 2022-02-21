// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Dummie returns stringer/JSON/text marshaler for the dummie type.
func Dummie(v interface{}) DummieV { return New().Dummie(v) }

// Dummie returns stringer/JSON/text marshaler for the dummie type.
func (pretty Pretty) Dummie(v interface{}) DummieV {
	return DummieV{
		v:        v,
		prettier: pretty,
	}
}

type DummieV struct {
	v        interface{}
	prettier Pretty
}

func (v DummieV) String() string {
	if v.v == nil {
		return v.prettier.nil
	}
	return fmt.Sprint(v.v)
}

func (v DummieV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v DummieV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}
	p, err := json.Marshal(v.v)
	if _, ok := err.(*json.UnsupportedTypeError); ok {
		return []byte(reflect.TypeOf(v.v).String()), nil
	}
	return p, err
}
