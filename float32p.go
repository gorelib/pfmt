// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Float32p returns stringer/JSON/text marshaler for the float32 pointer type.
func Float32p(p *float32) Float32P { return Float32P{p: p} }

type Float32P struct{ p *float32 }

func (p Float32P) String() string {
	if p.p == nil {
		return "null"
	}
	return Float32V{V: *p.p}.String()
}

func (p Float32P) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Float32V{V: *p.p}.MarshalText()
}

func (p Float32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
