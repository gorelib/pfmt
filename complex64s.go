// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Complex64s returns stringer/JSON/text marshaler for the complex64 type.
func Complex64s(s []complex64) Complex64S { return New().Complex64s(s) }

// Complex64s returns stringer/JSON/text marshaler for the complex64 type.
func (pretty Pretty) Complex64s(s []complex64) Complex64S {
	return Complex64S{
		s:        s,
		prettier: pretty,
	}
}

type Complex64S struct {
	s        []complex64
	prettier Pretty
}

func (s Complex64S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Complex64S) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Complex64(v).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(s.prettier.separator)
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func (s Complex64S) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Complex64(v).MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		buf.Write(b)
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
