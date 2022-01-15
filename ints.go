// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
)

// Ints returns stringer/JSON/text marshaler for the int slice type.
func Ints(s []int) IntS { return IntS{s: s} }

type IntS struct{ s []int }

func (s IntS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s IntS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := intV{V: v}.MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(" ")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s IntS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := intV{V: v}.MarshalJSON()
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
