// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint16ps returns stringer/JSON/text marshaler for the uint16 pointer slice type.
func Uint16ps(a []*uint16) Uint16PS { return New().Uint16ps(a) }

// Uint16ps returns stringer/JSON/text marshaler for the uint16 pointer slice type.
func (pretty Pretty) Uint16ps(a []*uint16) Uint16PS {
	return Uint16PS{
		a:        a,
		prettier: pretty,
	}
}

type Uint16PS struct {
	a        []*uint16
	prettier Pretty
}

func (a Uint16PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Uint16PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Uint16p(p).MarshalText()
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

func (a Uint16PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Uint16p(p).MarshalJSON()
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
