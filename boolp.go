// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Boolp returns stringer/JSON/text marshaler for the bool pointer type.
func Boolp(p *bool) BoolP { return New().Boolp(p) }

// Boolp returns stringer/JSON/text marshaler for the bool pointer type.
func (pretty Pretty) Boolp(p *bool) BoolP {
	return BoolP{
		p:        p,
		prettier: pretty,
	}
}

type BoolP struct {
	p        *bool
	prettier Pretty
}

func (p BoolP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p BoolP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Bool(*p.p).MarshalText()
}

func (p BoolP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.MarshalText()
}
