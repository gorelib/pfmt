// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Runesp returns stringer/JSON/text marshaler for the rune pointer slice type.
func Runesp(p *[]rune) RuneSP { return RuneSP{p: p} }

type RuneSP struct{ p *[]rune }

func (p RuneSP) String() string {
	if p.p == nil {
		return "null"
	}
	return Runes(*p.p).String()
}

func (p RuneSP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Runes(*p.p).MarshalText()
}

func (p RuneSP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Runes(*p.p).MarshalJSON()
}
