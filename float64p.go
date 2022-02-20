// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Float64p returns stringer/JSON/text marshaler for the  float64 pointer type.
func Float64p(p *float64) Float64P { return New().Float64p(p) }

// Float64p returns stringer/JSON/text marshaler for the  float64 pointer type.
func (pretty Pretty) Float64p(p *float64) Float64P {
	return Float64P{
		p:        p,
		prettier: pretty,
	}
}

type Float64P struct {
	p        *float64
	prettier Pretty
}

func (p Float64P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Float64(*p.p).String()
}

func (p Float64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Float64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
