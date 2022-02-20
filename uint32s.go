// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint32s returns stringer/JSON/text marshaler for the uint32 slice type.
func Uint32s(s []uint32) Uint32S { return New().Uint32s(s) }

// Uint32s returns stringer/JSON/text marshaler for the uint32 slice type.
func (pretty Pretty) Uint32s(s []uint32) Uint32S {
	return Uint32S{
		s:        s,
		prettier: pretty,
	}
}

type Uint32S struct {
	s        []uint32
	prettier Pretty
}

func (s Uint32S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Uint32S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Uint32(v).MarshalText()
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

func (s Uint32S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Uint32(v).MarshalJSON()
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
