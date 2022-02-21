// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"

	"github.com/pprint/pfmt/pencode"
)

// Errs returns stringer/JSON/text marshaler for the error slice type.
func Errs(s []error) ErrorS { return New().Errs(s) }

// Errors returns stringer/JSON/text marshaler for the error slice type.
func (pretty Pretty) Errs(s []error) ErrorS {
	return ErrorS{
		s:        s,
		prettier: pretty,
	}
}

type ErrorS struct {
	s        []error
	prettier Pretty
}

func (s ErrorS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s ErrorS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		if i != 0 {
			buf.WriteString(s.prettier.separator)
		}
		if v == nil {
			buf.WriteString(s.prettier.nil)
			continue
		}
		err := pencode.String(&buf, v.Error())
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s ErrorS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := s.prettier.Err(v).MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
