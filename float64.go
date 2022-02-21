// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "fmt"

// Float64 returns stringer/JSON/text marshaler for the float64 type.
func Float64(v float64) Float64V { return New().Float64(v) }

// Float64 returns stringer/JSON/text marshaler for the float64 type.
func (Pretty) Float64(v float64) Float64V { return Float64V{v: v} }

type Float64V struct {
	v float64
}

func (v Float64V) String() string {
	return fmt.Sprint(v.v)
}

func (v Float64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Float64V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
