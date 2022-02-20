// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Complex128p returns stringer/JSON/text marshaler for the complex128 pointer type.
func Complex128p(p *complex128) Complex128P { return New().Complex128p(p) }

// Complex128p returns stringer/JSON/text marshaler for the complex128 pointer type.
func (pretty Pretty) Complex128p(p *complex128) Complex128P {
	return Complex128P{
		p:        p,
		prettier: pretty,
	}
}

type Complex128P struct {
	p        *complex128
	prettier Pretty
}

func (p Complex128P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return Complex128(*p.p).String()
}

func (p Complex128P) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Complex128(*p.p).MarshalText()
}

func (p Complex128P) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.Complex128(*p.p).MarshalJSON()
}
