// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Boolps returns stringer/JSON/text marshaler for slice of bool pointers type.
func Boolps(s []*bool) BoolPS { return New().Boolps(s) }

// Boolps returns stringer/JSON/text marshaler for slice of bool pointers type.
func (pretty Pretty) Boolps(s []*bool) BoolPS {
	return BoolPS{
		s:        s,
		prettier: pretty,
	}
}

type BoolPS struct {
	s        []*bool
	prettier Pretty
}

func (s BoolPS) String() string {
	t, _ := s.MarshalText()
	return string(t)
}

func (s BoolPS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range s.s {
		b, err := s.prettier.Boolp(p).MarshalText()
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

func (s BoolPS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		b, err := Boolp(p).MarshalJSON()
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
