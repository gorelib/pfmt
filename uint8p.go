// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uint8p returns stringer/JSON/text marshaler for the uint8 pointer type.
func Uint8p(p *uint8) Uint8P { return Uint8P{p: p} }

type Uint8P struct{ p *uint8 }

func (p Uint8P) String() string {
	if p.p == nil {
		return "null"
	}
	return Uint8V{V: *p.p}.String()
}

func (p Uint8P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Uint8P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
