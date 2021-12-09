// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Intp returns stringer/JSON/text marshaler for the int pointer type.
func Intp(p *int) IntP { return IntP{p: p} }

type IntP struct{ p *int }

func (p IntP) String() string {
	if p.p == nil {
		return "null"
	}
	return intV{V: *p.p}.String()
}

func (p IntP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p IntP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
