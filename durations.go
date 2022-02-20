// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"time"
)

// Durations returns stringer/JSON/text marshaler for time duration slice type.
func Durations(s []time.Duration) DurationS { return New().Durations(s) }

// Durations returns stringer/JSON/text marshaler for time duration slice type.
func (pretty Pretty) Durations(s []time.Duration) DurationS {
	return DurationS{
		s:        s,
		prettier: pretty,
	}
}

type DurationS struct {
	s        []time.Duration
	prettier Pretty
}

func (s DurationS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s DurationS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Duration(v).MarshalText()
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

func (s DurationS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Duration(v).MarshalJSON()
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
