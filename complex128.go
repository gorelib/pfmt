// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"fmt"
)

// Complex128 returns stringer/JSON/text marshaler for the complex128 type.
func Complex128(v complex128) Complex128V { return New().Complex128(v) }

// Complex128 returns stringer/JSON/text marshaler for the complex128 type.
func (pretty Pretty) Complex128(v complex128) Complex128V {
	return Complex128V{v: v}
}

type Complex128V struct {
	v complex128
}

func (v Complex128V) String() string {
	s := fmt.Sprintf("%g", v.v)
	return s[1 : len(s)-1]
}

func (v Complex128V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Complex128V) MarshalJSON() ([]byte, error) {
	return append([]byte(`"`), append([]byte(v.String()), []byte(`"`)...)...), nil
}
