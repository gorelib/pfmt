// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Stringp returns stringer/JSON/text marshaler for the string pointer type.
func Stringp(p *string) StringP { return New().Stringp(p) }

// Stringp returns stringer/JSON/text marshaler for the string pointer type.
func (pretty Pretty) Stringp(p *string) StringP {
	return StringP{
		p:        p,
		prettier: pretty,
	}
}

type StringP struct {
	p        *string
	prettier Pretty
}

func (p StringP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.String(*p.p).String()
}

func (p StringP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.String(*p.p).MarshalText()
}

func (p StringP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.String(*p.p).MarshalJSON()
}
