// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uint64p returns stringer/JSON/text marshaler for the uint64 pointer type.
func Uint64p(p *uint64) Uint64P { return New().Uint64p(p) }

// Uint64p returns stringer/JSON/text marshaler for the uint64 pointer type.
func (pretty Pretty) Uint64p(p *uint64) Uint64P {
	return Uint64P{
		p:        p,
		prettier: pretty,
	}
}

type Uint64P struct {
	p        *uint64
	prettier Pretty
}

func (p Uint64P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Uint64(*p.p).String()
}

func (p Uint64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Uint64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
