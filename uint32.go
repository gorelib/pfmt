// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Uint32 returns stringer/JSON/text marshaler for the uint32 type.
func Uint32(v uint32) Uint32V { return New().Uint32(v) }

// Uint32 returns stringer/JSON/text marshaler for the uint32 type.
func (Pretty) Uint32(v uint32) Uint32V {
	return Uint32V{v: v}
}

type Uint32V struct {
	v uint32
}

func (v Uint32V) String() string {
	return strconv.FormatUint(uint64(v.v), 10)
}

func (v Uint32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Uint32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
