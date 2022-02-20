// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Uintptr returns stringer/JSON/text marshaler for the uintptr type.
func Uintptr(v uintptr) UintptrV { return New().Uintptr(v) }

// Uintptr returns stringer/JSON/text marshaler for the uintptr type.
func (pretty Pretty) Uintptr(v uintptr) UintptrV {
	return UintptrV{v: v}
}

type UintptrV struct {
	v uintptr
}

func (v UintptrV) String() string {
	return strconv.FormatUint(uint64(v.v), 10)
}

func (v UintptrV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v UintptrV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
