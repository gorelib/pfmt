// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uintp returns stringer/JSON/text marshaler for the uint pointer type.
func Uintp(p *uint) UintP { return New().Uintp(p) }

// Uintp returns stringer/JSON/text marshaler for the uint pointer type.
func (pretty Pretty) Uintp(p *uint) UintP {
	return UintP{
		p:        p,
		prettier: pretty,
	}
}

type UintP struct {
	p        *uint
	prettier Pretty
}

func (p UintP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Uint(*p.p).String()
}

func (p UintP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p UintP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
