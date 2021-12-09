// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Complex128p returns stringer/JSON/text marshaler for the complex128 pointer type.
func Complex128p(p *complex128) Complex128P { return Complex128P{p: p} }

type Complex128P struct{ p *complex128 }

func (p Complex128P) String() string {
	if p.p == nil {
		return "null"
	}
	return Complex128(*p.p).String()
}

func (p Complex128P) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Complex128(*p.p).MarshalText()
}

func (p Complex128P) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Complex128(*p.p).MarshalJSON()
}
