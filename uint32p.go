// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uint32p returns stringer/JSON/text marshaler for the uint32 pointer type.
func Uint32p(p *uint32) Uint32P { return New().Uint32p(p) }

// Uint32p returns stringer/JSON/text marshaler for the uint32 pointer type.
func (pretty Pretty) Uint32p(p *uint32) Uint32P {
	return Uint32P{
		p:        p,
		prettier: pretty,
	}
}

type Uint32P struct {
	p        *uint32
	prettier Pretty
}

func (p Uint32P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Uint32(*p.p).String()
}

func (p Uint32P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Uint32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
