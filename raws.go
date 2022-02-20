// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Raws returns stringer/JSON/text marshaler for the slice of byte slice type.
func Raws(s [][]byte) RawS { return New().Raws(s) }

// Raws returns stringer/JSON/text marshaler for the slice of byte slice type.
func (pretty Pretty) Raws(s [][]byte) RawS {
	return RawS{
		s:        s,
		prettier: pretty,
	}
}

type RawS struct {
	s        [][]byte
	prettier Pretty
}

func (s RawS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s RawS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Raw(v).MarshalText()
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

func (s RawS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Raw(v).MarshalJSON()
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
