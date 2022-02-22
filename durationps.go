// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"time"
)

// Durationps returns stringer/JSON/text marshaler for the time duration pointer slice type.
func Durationps(s []*time.Duration) DurationPS { return New().Durationps(s) }

// Durationps returns stringer/JSON/text marshaler for the time duration pointer slice type.
func (pretty Pretty) Durationps(s []*time.Duration) DurationPS {
	return DurationPS{
		s:        s,
		prettier: pretty,
	}
}

type DurationPS struct {
	s        []*time.Duration
	prettier Pretty
}

func (s DurationPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s DurationPS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range s.s {
		b, err := s.prettier.Durationp(p).MarshalText()
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

func (s DurationPS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.s {
		b, err := s.prettier.Durationp(p).MarshalJSON()
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
