// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int64p returns stringer/JSON/text marshaler for the int64 pointer type.
func Int64p(p *int64) Int64P { return New().Int64p(p) }

// Int64p returns stringer/JSON/text marshaler for the int64 pointer type.
func (pretty Pretty) Int64p(p *int64) Int64P {
	return Int64P{
		p:        p,
		prettier: pretty,
	}
}

type Int64P struct {
	p        *int64
	prettier Pretty
}

func (p Int64P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Int64(*p.p).String()
}

func (p Int64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
