// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Float32p returns stringer/JSON/text marshaler for the float32 pointer type.
func Float32p(p *float32) Float32P { return New().Float32p(p) }

// Float32p returns stringer/JSON/text marshaler for the float32 pointer type.
func (pretty Pretty) Float32p(p *float32) Float32P {
	return Float32P{
		p:        p,
		prettier: pretty,
	}
}

type Float32P struct {
	p        *float32
	prettier Pretty
}

func (p Float32P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Float32(*p.p).String()
}

func (p Float32P) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Float32(*p.p).MarshalText()
}

func (p Float32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
