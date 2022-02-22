// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
)

// Reflects returns stringer/JSON/text marshaler uses reflection for the slice of some type.
func Reflects(s []interface{}) ReflectS { return New().Reflects(s) }

// Reflects returns stringer/JSON/text marshaler uses reflection for the slice of some type.
func (pretty Pretty) Reflects(s []interface{}) ReflectS {
	return ReflectS{
		s:        s,
		prettier: pretty,
	}
}

type ReflectS struct {
	s        []interface{}
	prettier Pretty
}

func (s ReflectS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ReflectS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Reflect(v).MarshalText()
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

func (s ReflectS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Reflect(v).MarshalJSON()
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
