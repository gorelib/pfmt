// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"strconv"
)

// Int8 returns stringer/JSON/text marshaler for the int8 type.
func Int8(v int8) int8V { return int8V{V: v} }

type int8V struct{ V int8 }

func (v int8V) String() string {
	return strconv.Itoa(int(v.V))
}

func (v int8V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v int8V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
