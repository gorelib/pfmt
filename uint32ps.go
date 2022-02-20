// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint32ps returns stringer/JSON/text marshaler for the uint32 pointer slice type.
func Uint32ps(a []*uint32) Uint32PS { return New().Uint32ps(a) }

// Uint32ps returns stringer/JSON/text marshaler for the uint32 pointer slice type.
func (pretty Pretty) Uint32ps(a []*uint32) Uint32PS {
	return Uint32PS{
		a:        a,
		prettier: pretty,
	}
}

type Uint32PS struct {
	a        []*uint32
	prettier Pretty
}

func (a Uint32PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Uint32PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Uint32p(p).MarshalText()
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

func (a Uint32PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Uint32p(p).MarshalJSON()
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
