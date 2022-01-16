// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Time returns stringer/JSON/text marshaler for the time type.
func Time(v time.Time) TimeV { return TimeV{V: v} }

type TimeV struct{ V time.Time }

func (v TimeV) String() string {
	return v.V.String()
}

func (v TimeV) MarshalText() ([]byte, error) {
	return v.V.MarshalText()
}

func (v TimeV) MarshalJSON() ([]byte, error) {
	return v.V.MarshalJSON()
}
