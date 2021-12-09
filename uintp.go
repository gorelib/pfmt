// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uintp returns stringer/JSON/text marshaler for the uint pointer type.
func Uintp(p *uint) UintP { return UintP{p: p} }

type UintP struct{ p *uint }

func (p UintP) String() string {
	if p.p == nil {
		return "null"
	}
	return uintV{V: *p.p}.String()
}

func (p UintP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p UintP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
