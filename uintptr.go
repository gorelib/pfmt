// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"strconv"
)

// Uintptr returns stringer/JSON/text marshaler for the uintptr type.
func Uintptr(v uintptr) uintptrV { return uintptrV{V: v} }

type uintptrV struct{ V uintptr }

func (v uintptrV) String() string {
	return strconv.FormatUint(uint64(v.V), 10)
}

func (v uintptrV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v uintptrV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
