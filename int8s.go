// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int8s returns stringer/JSON/text marshaler for the int8 slice type.
func Int8s(s []int8) Int8S { return New().Int8s(s) }

// Int8s returns stringer/JSON/text marshaler for the int8 slice type.
func (pretty Pretty) Int8s(s []int8) Int8S {
	return Int8S{
		s:        s,
		prettier: pretty,
	}
}

type Int8S struct {
	s        []int8
	prettier Pretty
}

func (s Int8S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Int8S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Int8(v).MarshalText()
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

func (s Int8S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Int8(v).MarshalJSON()
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
