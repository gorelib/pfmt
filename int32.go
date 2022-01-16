// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Int32 returns stringer/JSON/text marshaler for the int32 type.
func Int32(v int32) Int32V { return Int32V{V: v} }

type Int32V struct{ V int32 }

func (v Int32V) String() string {
	return strconv.Itoa(int(v.V))
}

func (v Int32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Int32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
