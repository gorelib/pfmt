// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint64s returns stringer/JSON/text marshaler for the uint64 slice type.
func Uint64s(s []uint64) Uint64S { return New().Uint64s(s) }

// Uint64s returns stringer/JSON/text marshaler for the uint64 slice type.
func (pretty Pretty) Uint64s(s []uint64) Uint64S {
	return Uint64S{
		s:        s,
		prettier: pretty,
	}
}

type Uint64S struct {
	s        []uint64
	prettier Pretty
}

func (s Uint64S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Uint64S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Uint64(v).MarshalText()
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

func (s Uint64S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Uint64(v).MarshalJSON()
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
