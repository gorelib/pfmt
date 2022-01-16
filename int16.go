// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Int16 returns stringer/JSON/text marshaler for the int16 type.
func Int16(v int16) Int16V { return Int16V{V: v} }

type Int16V struct{ V int16 }

func (v Int16V) String() string {
	return strconv.Itoa(int(v.V))
}

func (v Int16V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Int16V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
