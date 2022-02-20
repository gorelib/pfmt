// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Uintptrp returns stringer/JSON/text marshaler for the uintptr pointer type.
func Uintptrp(p *uintptr) UintptrP { return New().Uintptrp(p) }

// Uintptrp returns stringer/JSON/text marshaler for the uintptr pointer type.
func (pretty Pretty) Uintptrp(p *uintptr) UintptrP {
	return UintptrP{
		p:        p,
		prettier: pretty,
	}
}

type UintptrP struct {
	p        *uintptr
	prettier Pretty
}

func (p UintptrP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Uintptr(*p.p).String()
}

func (p UintptrP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p UintptrP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
