// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"

	"github.com/pfmt/pfmt/pencode"
)

// Runes returns stringer/JSON/text marshaler for the rune slice type.
func Runes(s []rune) RuneS { return New().Runes(s) }

// Runes returns stringer/JSON/text marshaler for the rune slice type.
func (pretty Pretty) Runes(s []rune) RuneS {
	return RuneS{
		s:        s,
		prettier: pretty,
	}
}

type RuneS struct {
	s        []rune
	prettier Pretty
}

func (s RuneS) String() string {
	if s.s == nil {
		return s.prettier.nil
	}
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	err := pencode.Runes(buf, s.s)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (s RuneS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer

	err := pencode.Runes(&buf, s.s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s RuneS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}

	b, err := s.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
