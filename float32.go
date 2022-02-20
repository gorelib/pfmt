// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"fmt"
)

// Float32 returns stringer/JSON/text marshaler for the float32 type.
func Float32(v float32) Float32V { return New().Float32(v) }

// Float32 returns stringer/JSON/text marshaler for the float32 type.
func (Pretty) Float32(v float32) Float32V {
	return Float32V{v: v}
}

type Float32V struct {
	v float32
}

func (v Float32V) String() string {
	return fmt.Sprint(v.v)
}

func (v Float32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Float32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
