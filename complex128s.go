// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Complex128s returns stringer/JSON/text marshaler for the complex128 slice type.
func Complex128s(s []complex128) Complex128S { return New().Complex128s(s) }

// Complex128s returns stringer/JSON/text marshaler for the complex128 slice type.
func (pretty Pretty) Complex128s(s []complex128) Complex128S {
	return Complex128S{
		s:        s,
		prettier: pretty,
	}
}

type Complex128S struct {
	s        []complex128
	prettier Pretty
}

func (s Complex128S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Complex128S) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Complex128(v).MarshalText()
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

func (s Complex128S) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Complex128(v).MarshalJSON()
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
