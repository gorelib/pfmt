// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Timep returns stringer/JSON/text marshaler for the time pointer type.
func Timep(p *time.Time) TimeP { return TimeP{p: p} }

type TimeP struct{ p *time.Time }

func (p TimeP) String() string {
	if p.p == nil {
		return "null"
	}
	return timeV{V: *p.p}.String()
}

func (p TimeP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return timeV{V: *p.p}.MarshalText()
}

func (p TimeP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return timeV{V: *p.p}.MarshalJSON()
}
