// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding"
)

// Texts returns stringer/JSON/text marshaler for the text marshaler slice type.
func Texts(s []encoding.TextMarshaler) TextS { return New().Texts(s) }

// Texts returns stringer/JSON/text marshaler for the text marshaler slice type.
func (pretty Pretty) Texts(s []encoding.TextMarshaler) TextS {
	return TextS{
		s:        s,
		prettier: pretty,
	}
}

type TextS struct {
	s        []encoding.TextMarshaler
	prettier Pretty
}

func (s TextS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s TextS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Text(v).MarshalText()
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

func (s TextS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Text(v).MarshalJSON()
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
