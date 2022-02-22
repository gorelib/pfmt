// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int64s returns stringer/JSON/text marshaler for the int64 slice type.
func Int64s(s []int64) Int64S { return New().Int64s(s) }

// Int64s returns stringer/JSON/text marshaler for the int64 slice type.
func (pretty Pretty) Int64s(s []int64) Int64S {
	return Int64S{
		s:        s,
		prettier: pretty,
	}
}

type Int64S struct {
	s        []int64
	prettier Pretty
}

func (s Int64S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Int64S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Int64(v).MarshalText()
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

func (s Int64S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Int64(v).MarshalJSON()
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
