// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uintptrp returns stringer/JSON/text marshaler for the uintptr pointer type.
func Uintptrp(p *uintptr) UintptrP { return UintptrP{p: p} }

type UintptrP struct{ p *uintptr }

func (p UintptrP) String() string {
	if p.p == nil {
		return "null"
	}
	return uintptrV{V: *p.p}.String()
}

func (p UintptrP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p UintptrP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
