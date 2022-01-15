// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
)

// Rawps returns stringer/JSON/text marshaler for the slice of byte slice pointers type.
func Rawps(s []*[]byte) RawPS { return RawPS{s: s} }

type RawPS struct{ s []*[]byte }

func (s RawPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s RawPS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range s.s {
		b, err := Rawp(p).MarshalText()
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

func (s RawPS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		b, err := Rawp(p).MarshalJSON()
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
