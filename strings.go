// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"strings"

	"github.com/pfmt/pfmt/pencode"
)

// Strings returns stringer/JSON/text marshaler for the string slice type.
func Strings(s []string) StringS { return New().Strings(s) }

// Strings returns stringer/JSON/text marshaler for the string slice type.
func (pretty Pretty) Strings(s []string) StringS {
	return StringS{
		s:        s,
		prettier: pretty,
	}
}

type StringS struct {
	s        []string
	prettier Pretty
}

func (s StringS) String() string {
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	err := pencode.String(buf, strings.Join(s.s, " "))
	if err != nil {
		return s.prettier.empty
	}
	return buf.String()
}

func (s StringS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	err := pencode.String(&buf, strings.Join(s.s, " "))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s StringS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.String(v).MarshalJSON()
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
