// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Stringp returns stringer/JSON/text marshaler for the string pointer type.
func Stringp(p *string) StringP { return StringP{p: p} }

type StringP struct{ p *string }

func (p StringP) String() string {
	if p.p == nil {
		return "null"
	}
	return stringV{V: *p.p}.String()
}

func (p StringP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return stringV{V: *p.p}.MarshalText()
}

func (p StringP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return stringV{V: *p.p}.MarshalJSON()
}
