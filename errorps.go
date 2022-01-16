// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Errorps returns stringer/JSON/text marshaler for the slice of error pointers type.
func Errorps(s []*error) ErrorPS { return ErrorPS{s: s} }

type ErrorPS struct{ s []*error }

func (s ErrorPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ErrorPS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	var tail bool

	for _, p := range s.s {
		if p == nil {
			continue
		}
		b, err := Errorp(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if tail {
			buf.WriteString(" ")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
		tail = true
	}
	return buf.Bytes(), nil
}

func (s ErrorPS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		b, err := Errorp(p).MarshalJSON()
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
