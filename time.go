// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Time returns stringer/JSON/text marshaler for the time type.
func Time(v time.Time) TimeV { return New().Time(v) }

// Time returns stringer/JSON/text marshaler for the time type.
func (pretty Pretty) Time(v time.Time) TimeV {
	return TimeV{v: v}
}

type TimeV struct {
	v time.Time
}

func (v TimeV) String() string {
	return v.v.String()
}

func (v TimeV) MarshalText() ([]byte, error) {
	return v.v.MarshalText()
}

func (v TimeV) MarshalJSON() ([]byte, error) {
	return v.v.MarshalJSON()
}
