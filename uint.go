// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Uint returns stringer/JSON/text marshaler for the uint type.
func Uint(v uint) UintV { return New().Uint(v) }

// Uint returns stringer/JSON/text marshaler for the uint type.
func (pretty Pretty) Uint(v uint) UintV {
	return UintV{V: v}
}

type UintV struct {
	V uint
}

func (v UintV) String() string {
	return strconv.FormatUint(uint64(v.V), 10)
}

func (v UintV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v UintV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
