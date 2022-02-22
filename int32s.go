// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int32s returns stringer/JSON/text marshaler for the int32 slice type.
func Int32s(s []int32) Int32S { return New().Int32s(s) }

// Int32s returns stringer/JSON/text marshaler for the int32 slice type.
func (pretty Pretty) Int32s(s []int32) Int32S {
	return Int32S{
		s:        s,
		prettier: pretty,
	}
}

type Int32S struct {
	s        []int32
	prettier Pretty
}

func (s Int32S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Int32S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Int32(v).MarshalText()
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

func (s Int32S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Int32(v).MarshalJSON()
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
