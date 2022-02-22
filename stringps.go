// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Stringps returns stringer/JSON/text marshaler for the string pointer slice type.
func Stringps(a []*string) StringPS { return New().Stringps(a) }

// Stringps returns stringer/JSON/text marshaler for the string pointer slice type.
func (pretty Pretty) Stringps(a []*string) StringPS {
	return StringPS{
		a:        a,
		prettier: pretty,
	}
}

type StringPS struct {
	a        []*string
	prettier Pretty
}

func (a StringPS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a StringPS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Stringp(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(a.prettier.separator)
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func (a StringPS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Stringp(p).MarshalJSON()
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
