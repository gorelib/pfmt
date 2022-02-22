// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Float64ps returns stringer/JSON/text marshaler for the float64 pointer slice type.
func Float64ps(a []*float64) Float64PS { return New().Float64ps(a) }

// Float64ps returns stringer/JSON/text marshaler for the float64 pointer slice type.
func (pretty Pretty) Float64ps(a []*float64) Float64PS {
	return Float64PS{
		a:        a,
		prettier: pretty,
	}
}

type Float64PS struct {
	a        []*float64
	prettier Pretty
}

func (a Float64PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Float64PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Float64p(p).MarshalText()
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

func (a Float64PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Float64p(p).MarshalJSON()
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
