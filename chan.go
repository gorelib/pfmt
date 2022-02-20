// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Chan returns stringer/JSON/text marshaler for the channel.
func Chan(v interface{}) ChanV { return New().Chan(v) }

// Chan returns stringer/JSON/text marshaler for the channel.
func (pretty Pretty) Chan(v interface{}) ChanV {
	return ChanV{
		v:        v,
		prettier: pretty,
	}
}

type ChanV struct {
	v        interface{}
	prettier Pretty
}

func (v ChanV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v ChanV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Chan {
		return nil, errors.New("not channel")
	}
	if val.IsNil() {
		return []byte(v.prettier.nil), nil
	}

	return []byte(fmt.Sprint(v.v)), nil
}

func (v ChanV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}

	val := reflect.ValueOf(v.v)

	if val.Kind() != reflect.Chan {
		return nil, errors.New("not channel")
	}
	if val.IsNil() {
		return []byte("null"), nil
	}

	return json.Marshal(v.v)
}
