// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Int64 returns stringer/JSON/text marshaler for the int64 type.
func Int64(v int64) Int64V { return New().Int64(v) }

// Int64 returns stringer/JSON/text marshaler for the int64 type.
func (Pretty) Int64(v int64) Int64V {
	return Int64V{v: v}
}

type Int64V struct {
	v int64
}

func (v Int64V) String() string {
	return strconv.FormatInt(int64(v.v), 10)
}

func (v Int64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Int64V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
