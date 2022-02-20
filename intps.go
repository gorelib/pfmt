// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Intps returns stringer/JSON/text marshaler for the int pointer slice type.
func Intps(a []*int) IntPS { return New().Intps(a) }

// Intps returns stringer/JSON/text marshaler for the int pointer slice type.
func (pretty Pretty) Intps(a []*int) IntPS {
	return IntPS{
		a:        a,
		prettier: pretty,
	}
}

type IntPS struct {
	a        []*int
	prettier Pretty
}

func (a IntPS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a IntPS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Intp(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(a.prettier.separator)
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (a IntPS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Intp(p).MarshalJSON()
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
