// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Float32s returns stringer/JSON/text marshaler for the float32 slice type.
func Float32s(s []float32) Float32S { return New().Float32s(s) }

// Float32s returns stringer/JSON/text marshaler for the float32 slice type.
func (pretty Pretty) Float32s(s []float32) Float32S {
	return Float32S{
		s:        s,
		prettier: pretty,
	}
}

type Float32S struct {
	s        []float32
	prettier Pretty
}

func (s Float32S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Float32S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Float32(v).MarshalText()
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

func (s Float32S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Float32(v).MarshalJSON()
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
