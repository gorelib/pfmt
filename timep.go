// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Timep returns stringer/JSON/text marshaler for the time pointer type.
func Timep(p *time.Time) TimeP { return New().Timep(p) }

// Timep returns stringer/JSON/text marshaler for the time pointer type.
func (pretty Pretty) Timep(p *time.Time) TimeP {
	return TimeP{
		p:        p,
		prettier: pretty,
	}
}

type TimeP struct {
	p        *time.Time
	prettier Pretty
}

func (p TimeP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Time(*p.p).String()
}

func (p TimeP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Time(*p.p).MarshalText()
}

func (p TimeP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.Time(*p.p).MarshalJSON()
}
