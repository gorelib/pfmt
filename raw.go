// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Raw returns stringer/JSON/text marshaler for the raw byte slice.
func Raw(v []byte) RawV { return RawV{V: v} }

type RawV struct{ V []byte }

func (v RawV) String() string {
	if v.V == nil {
		return "null"
	}
	return string(v.V)
}

func (v RawV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}
	return v.V, nil
}

func (v RawV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
