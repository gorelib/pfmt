// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uintptrs returns stringer/JSON/text marshaler for the uintptr slice type.
func Uintptrs(s []uintptr) UintptrS { return New().Uintptrs(s) }

// Uintptrs returns stringer/JSON/text marshaler for the uintptr slice type.
func (pretty Pretty) Uintptrs(s []uintptr) UintptrS {
	return UintptrS{
		s:        s,
		prettier: pretty,
	}
}

type UintptrS struct {
	s        []uintptr
	prettier Pretty
}

func (s UintptrS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s UintptrS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Uintptr(v).MarshalText()
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

func (s UintptrS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Uintptr(v).MarshalJSON()
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
