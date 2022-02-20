// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Runesp returns stringer/JSON/text marshaler for the rune pointer slice type.
func Runesp(p *[]rune) RuneSP { return New().Runesp(p) }

// Runesp returns stringer/JSON/text marshaler for the rune pointer slice type.
func (pretty Pretty) Runesp(p *[]rune) RuneSP {
	return RuneSP{
		p:        p,
		prettier: pretty,
	}
}

type RuneSP struct {
	p        *[]rune
	prettier Pretty
}

func (p RuneSP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Runes(*p.p).String()
}

func (p RuneSP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Runes(*p.p).MarshalText()
}

func (p RuneSP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.Runes(*p.p).MarshalJSON()
}
