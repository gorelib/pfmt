// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Bytesp returns stringer/JSON/text marshaler for the pointer to byte slice type.
func Bytesp(p *[]byte) ByteSP { return ByteSP{p: p} }

type ByteSP struct{ p *[]byte }

func (p ByteSP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p ByteSP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Bytes(*p.p...).MarshalText()
}

func (p ByteSP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Bytes(*p.p...).MarshalJSON()
}
