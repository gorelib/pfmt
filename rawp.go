// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Rawp returns stringer/JSON/text marshaler for the raw byte slice pointer.
func Rawp(p *[]byte) RawP { return RawP{p: p} }

type RawP struct{ p *[]byte }

func (p RawP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p RawP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return RawV{V: *p.p}.MarshalText()
}

func (p RawP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return RawV{V: *p.p}.MarshalJSON()
}
