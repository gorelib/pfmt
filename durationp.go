// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Durationp returns stringer/JSON/text marshaler for the  time duration pointer type.
func Durationp(p *time.Duration) DurationP { return DurationP{p: p} }

type DurationP struct{ p *time.Duration }

func (p DurationP) String() string {
	if p.p == nil {
		return "null"
	}
	return Duration(*p.p).String()
}

func (p DurationP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Duration(*p.p).MarshalText()
}

func (p DurationP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return Duration(*p.p).MarshalJSON()
}
