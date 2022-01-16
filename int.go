// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Int returns stringer/JSON/text marshaler for the int type.
func Int(v int) IntV { return IntV{V: v} }

type IntV struct{ V int }

func (v IntV) String() string {
	return strconv.Itoa(v.V)
}

func (v IntV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v IntV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
