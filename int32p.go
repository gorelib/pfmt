// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int32p returns stringer/JSON/text marshaler for the int32 pointer type.
func Int32p(p *int32) Int32P { return New().Int32p(p) }

// Int32p returns stringer/JSON/text marshaler for the int32 pointer type.
func (pretty Pretty) Int32p(p *int32) Int32P {
	return Int32P{
		p:        p,
		prettier: pretty,
	}
}

type Int32P struct {
	p        *int32
	prettier Pretty
}

func (p Int32P) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Int32(*p.p).String()
}

func (p Int32P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
