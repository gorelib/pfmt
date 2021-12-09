// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int16p returns stringer/JSON/text marshaler for the int16 pointer type.
func Int16p(p *int16) Int16P { return Int16P{p: p} }

type Int16P struct{ p *int16 }

func (p Int16P) String() string {
	if p.p == nil {
		return "null"
	}
	return int16V{V: *p.p}.String()
}

func (p Int16P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int16P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
