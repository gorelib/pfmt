// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Bytess returns stringer/JSON/text marshaler for the slice of byte slice type.
func Bytess(s [][]byte) ByteSS { return New().Bytess(s) }

// Bytess returns stringer/JSON/text marshaler for the slice of byte slice type.
func (pretty Pretty) Bytess(s [][]byte) ByteSS {
	return ByteSS{
		s:        s,
		prettier: pretty,
	}
}

type ByteSS struct {
	s        [][]byte
	prettier Pretty
}

func (s ByteSS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ByteSS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, a := range s.s {
		b, err := s.prettier.Bytes(a).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(s.prettier.separator)
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s ByteSS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, a := range s.s {
		b, err := s.prettier.Bytes(a).MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
