// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"time"
)

// Durationps returns stringer/JSON/text marshaler for the time duration pointer slice type.
func Durationps(s []*time.Duration) DurationPS { return DurationPS{s: s} }

type DurationPS struct{ s []*time.Duration }

func (s DurationPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s DurationPS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range s.s {
		b, err := Durationp(p).MarshalText()
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

func (s DurationPS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		b, err := Durationp(p).MarshalJSON()
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
