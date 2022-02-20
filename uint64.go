// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Uint64 returns stringer/JSON/text marshaler for the uint64 type.
func Uint64(v uint64) Uint64V { return New().Uint64(v) }

// Uint64 returns stringer/JSON/text marshaler for the uint64 type.
func (Pretty) Uint64(v uint64) Uint64V {
	return Uint64V{v: v}
}

type Uint64V struct {
	v uint64
}

func (v Uint64V) String() string {
	return strconv.FormatUint(v.v, 10)
}

func (v Uint64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Uint64V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
