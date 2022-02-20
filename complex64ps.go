// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Complex64ps returns stringer/JSON/text marshaler for the slice of complex64 pointers type.
func Complex64ps(s []*complex64) Complex64PS { return New().Complex64ps(s) }

// Complex64ps returns stringer/JSON/text marshaler for the slice of complex64 pointers type.
func (pretty Pretty) Complex64ps(s []*complex64) Complex64PS {
	return Complex64PS{
		s:        s,
		prettier: pretty,
	}
}

type Complex64PS struct {
	s        []*complex64
	prettier Pretty
}

func (s Complex64PS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Complex64PS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range s.s {
		b, err := s.prettier.Complex64p(p).MarshalText()
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

func (s Complex64PS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		b, err := s.prettier.Complex64p(p).MarshalJSON()
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
