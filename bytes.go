// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"

	"github.com/gorelib/pfmt/pencode"
)

// Bytes returns stringer/JSON/text marshaler for slice of bytes type.
func Bytes(s []byte) ByteS { return New().Bytes(s) }

// Bytes returns stringer/JSON/text marshaler for slice of bytes type.
func (pretty Pretty) Bytes(s []byte) ByteS {
	return ByteS{
		s:        s,
		prettier: pretty,
	}
}

type ByteS struct {
	s        []byte
	prettier Pretty
}

func (s ByteS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ByteS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
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
