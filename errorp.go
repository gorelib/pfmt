// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Errorp returns stringer/JSON/text marshaler for the error pointer type.
func Errorp(p *error) ErrorP { return ErrorP{p: p} }

type ErrorP struct{ p *error }

func (p ErrorP) String() string {
	if p.p == nil {
		return "null"
	}
	return Error(*p.p).String()
}

func (p ErrorP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Error(*p.p).MarshalText()
}

func (p ErrorP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Error(*p.p).MarshalJSON()
}
