// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
)

// Anys returns stringer/JSON/text marshaler for the slice of any type.
func Anys(s []interface{}) AnyS { return AnyS{s: s} }

type AnyS struct{ s []interface{} }

func (s AnyS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s AnyS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := Any(v).MarshalText()
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

func (s AnyS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := Any(v).MarshalJSON()
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
