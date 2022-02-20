// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"fmt"
)

// Complex64 returns stringer/JSON/text marshaler for the complex64 type.
func Complex64(v complex64) Complex64V { return New().Complex64(v) }

// Complex64 returns stringer/JSON/text marshaler for the complex64 type.
func (pretty Pretty) Complex64(v complex64) Complex64V {
	return Complex64V{v: v}
}

type Complex64V struct {
	v complex64
}

func (v Complex64V) String() string {
	s := fmt.Sprintf("%g", v.v)
	return s[1 : len(s)-1]
}

func (v Complex64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Complex64V) MarshalJSON() ([]byte, error) {
	return append([]byte(`"`), append([]byte(v.String()), []byte(`"`)...)...), nil
}
