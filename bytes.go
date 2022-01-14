// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"sync"

	"github.com/pprint/pfmt/pencode"
)

var bufPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// Bytes returns stringer/JSON/text marshaler for slice of bytes type.
func Bytes(s ...byte) ByteS { return ByteS{s: s} }

type ByteS struct{ s []byte }

func (s ByteS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ByteS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	err := pencode.Bytes(&buf, s.s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s ByteS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	b, err := s.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
