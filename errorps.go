// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Errps returns stringer/JSON/text marshaler for the slice of error pointers type.
func Errps(s []*error) ErrorPS { return New().Errps(s) }

// Errorps returns stringer/JSON/text marshaler for the slice of error pointers type.
func (pretty Pretty) Errps(s []*error) ErrorPS {
	return ErrorPS{
		s:        s,
		prettier: pretty,
	}
}

type ErrorPS struct {
	s        []*error
	prettier Pretty
}

func (s ErrorPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ErrorPS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range s.s {
		if i != 0 {
			buf.WriteString(s.prettier.separator)
		}
		if p == nil {
			buf.WriteString(s.prettier.nil)
			continue
		}
		b, err := s.prettier.Errp(p).MarshalText()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s ErrorPS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := s.prettier.Errp(p).MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
