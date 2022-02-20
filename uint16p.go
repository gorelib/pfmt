// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uint16p returns stringer/JSON/text marshaler for the uint16 pointer type.
func Uint16p(p *uint16) Uint16P { return New().Uint16p(p) }

// Uint16p returns stringer/JSON/text marshaler for the uint16 pointer type.
func (pretty Pretty) Uint16p(p *uint16) Uint16P {
	return Uint16P{
		p:        p,
		prettier: pretty,
	}
}

type Uint16P struct {
	p        *uint16
	prettier Pretty
}

func (p Uint16P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Uint16(*p.p).String()
}

func (p Uint16P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Uint16P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
