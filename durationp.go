// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Durationp returns stringer/JSON/text marshaler for the  time duration pointer type.
func Durationp(p *time.Duration) DurationP { return New().Durationp(p) }

// Durationp returns stringer/JSON/text marshaler for the  time duration pointer type.
func (pretty Pretty) Durationp(p *time.Duration) DurationP {
	return DurationP{
		p:        p,
		prettier: pretty,
	}
}

type DurationP struct {
	p        *time.Duration
	prettier Pretty
}

func (p DurationP) String() string {
	if p.p == nil {
		return p.prettier.nil
	}
	return p.prettier.Duration(*p.p).String()
}

func (p DurationP) MarshalText() ([]byte, error) {
	if p.p == nil {
		return []byte(p.prettier.nil), nil
	}
	return p.prettier.Duration(*p.p).MarshalText()
}

func (p DurationP) MarshalJSON() ([]byte, error) {
	if p.p == nil {
		return []byte("null"), nil
	}
	return p.prettier.Duration(*p.p).MarshalJSON()
}
