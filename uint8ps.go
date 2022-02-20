// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint8ps returns stringer/JSON/text marshaler for the uint8 pointer slice type.
func Uint8ps(a []*uint8) Uint8PS { return New().Uint8ps(a) }

// Uint8ps returns stringer/JSON/text marshaler for the uint8 pointer slice type.
func (pretty Pretty) Uint8ps(a []*uint8) Uint8PS {
	return Uint8PS{
		a:        a,
		prettier: pretty,
	}
}

type Uint8PS struct {
	a        []*uint8
	prettier Pretty
}

func (a Uint8PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Uint8PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Uint8p(p).MarshalText()
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

func (a Uint8PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Uint8p(p).MarshalJSON()
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
