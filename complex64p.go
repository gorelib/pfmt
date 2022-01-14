// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Complex64p returns stringer/JSON/text marshaler for the complex64 pointer type.
func Complex64p(p *complex64) Complex64P { return Complex64P{p: p} }

type Complex64P struct{ p *complex64 }

func (p Complex64P) String() string {
	if p.p == nil {
		return "null"
	}
	return Complex64(*p.p).String()
}

func (p Complex64P) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Complex64(*p.p).MarshalText()
}

func (p Complex64P) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Complex64(*p.p).MarshalJSON()
}
