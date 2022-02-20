// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Uint8 returns stringer/JSON/text marshaler for the uint8 type.
func Uint8(v uint8) Uint8V { return New().Uint8(v) }

// Uint8 returns stringer/JSON/text marshaler for the uint8 type.
func (Pretty) Uint8(v uint8) Uint8V {
	return Uint8V{v: v}
}

type Uint8V struct {
	v uint8
}

func (v Uint8V) String() string {
	return strconv.FormatUint(uint64(v.v), 10)
}

func (v Uint8V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Uint8V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
