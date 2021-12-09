// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int8p returns stringer/JSON/text marshaler for the int8 pointer type.
func Int8p(p *int8) Int8P { return Int8P{p: p} }

type Int8P struct{ p *int8 }

func (p Int8P) String() string {
	if p.p == nil {
		return "null"
	}
	return int8V{V: *p.p}.String()
}

func (p Int8P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int8P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
