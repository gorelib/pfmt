// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Errp returns stringer/JSON/text marshaler for the error pointer type.
func Errp(p *error) ErrorP { return New().Errp(p) }

// Errp returns stringer/JSON/text marshaler for the error pointer type.
func (pretty Pretty) Errp(p *error) ErrorP {
	return ErrorP{
		p:        p,
		prettier: pretty,
	}
}

type ErrorP struct {
	p        *error
	prettier Pretty
}

func (p ErrorP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return Err(*p.p).String()
}

func (p ErrorP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Err(*p.p).MarshalText()
}

func (p ErrorP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.Err(*p.p).MarshalJSON()
}
