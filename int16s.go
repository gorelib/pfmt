// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int16s returns stringer/JSON/text marshaler for the int16 slice type.
func Int16s(s []int16) Int16S { return New().Int16s(s) }

// Int16s returns stringer/JSON/text marshaler for the int16 slice type.
func (pretty Pretty) Int16s(s []int16) Int16S {
	return Int16S{
		s:        s,
		prettier: pretty,
	}
}

type Int16S struct {
	s        []int16
	prettier Pretty
}

func (s Int16S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Int16S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Int16(v).MarshalText()
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

func (s Int16S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Int16(v).MarshalJSON()
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
