// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Complex64p returns stringer/JSON/text marshaler for the complex64 pointer type.
func Complex64p(p *complex64) Complex64P { return New().Complex64p(p) }

// Complex64p returns stringer/JSON/text marshaler for the complex64 pointer type.
func (pretty Pretty) Complex64p(p *complex64) Complex64P {
	return Complex64P{
		p:        p,
		prettier: pretty,
	}
}

type Complex64P struct {
	p        *complex64
	prettier Pretty
}

func (p Complex64P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Complex64(*p.p).String()
}

func (p Complex64P) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Complex64(*p.p).MarshalText()
}

func (p Complex64P) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.Complex64(*p.p).MarshalJSON()
}
