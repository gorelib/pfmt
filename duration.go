// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"time"
)

// Duration returns stringer/JSON/text marshaler for the time duration type.
func Duration(v time.Duration) DurationV { return DurationV{v: v} }

type DurationV struct{ v time.Duration }

func (v DurationV) String() string {
	return v.v.String()
}

func (v DurationV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v DurationV) MarshalJSON() ([]byte, error) {
	return append([]byte(`"`), append([]byte(v.String()), []byte(`"`)...)...), nil
}
