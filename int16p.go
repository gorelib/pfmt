// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int16p returns stringer/JSON/text marshaler for the int16 pointer type.
func Int16p(p *int16) Int16P { return New().Int16p(p) }

// Int16p returns stringer/JSON/text marshaler for the int16 pointer type.
func (pretty Pretty) Int16p(p *int16) Int16P {
	return Int16P{
		p:        p,
		prettier: pretty,
	}
}

type Int16P struct {
	p        *int16
	prettier Pretty
}

func (p Int16P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Int16(*p.p).String()
}

func (p Int16P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int16P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
