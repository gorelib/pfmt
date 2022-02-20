// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"encoding"
	"encoding/json"
)

// KV is a key-value pair.
type KV interface {
	encoding.TextMarshaler
	json.Marshaler
}

// KVFunc returns stringer/JSON/text marshaler for the custom function type.
func KVFunc(v func() KV) KVFuncV { return New().KVFunc(v) }

// KVFunc returns stringer/JSON/text marshaler for the custom function type.
func (pretty Pretty) KVFunc(v func() KV) KVFuncV {
	return KVFuncV{v: v}
}

type KVFuncV struct {
	v func() KV
}

func (v KVFuncV) String() string {
	b, _ := v.v().MarshalText()
	return string(b)
}

func (v KVFuncV) MarshalText() ([]byte, error) {
	return v.v().MarshalText()
}

func (v KVFuncV) MarshalJSON() ([]byte, error) {
	return v.v().MarshalJSON()
}
