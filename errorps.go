// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
)

// Errorps returns stringer/JSON/text marshaler for the slice of error pointers type.
func Errorps(s ...*error) errorPS { return errorPS{S: s} }

type errorPS struct{ S []*error }

func (s errorPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s errorPS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	var tail bool

	for _, p := range s.S {
		if p == nil {
			continue
		}
		b, err := errorP{P: p}.MarshalText()
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

func (s errorPS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.S {
		b, err := errorP{P: p}.MarshalJSON()
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
