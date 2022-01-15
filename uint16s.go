// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
)

// Uint16s returns stringer/JSON/text marshaler for the uint16 slice type.
func Uint16s(s []uint16) Uint16S { return Uint16S{s: s} }

type Uint16S struct{ s []uint16 }

func (s Uint16S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Uint16S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := uint16V{V: v}.MarshalText()
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

func (s Uint16S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := uint16V{V: v}.MarshalJSON()
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
