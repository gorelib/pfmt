// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uint64p returns stringer/JSON/text marshaler for the uint64 pointer type.
func Uint64p(p *uint64) Uint64P { return Uint64P{p: p} }

type Uint64P struct{ p *uint64 }

func (p Uint64P) String() string {
	if p.p == nil {
		return "null"
	}
	return uint64V{V: *p.p}.String()
}

func (p Uint64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Uint64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
