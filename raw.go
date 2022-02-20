// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Raw returns stringer/JSON/text marshaler for the raw byte slice.
func Raw(v []byte) RawV { return New().Raw(v) }

// Raw returns stringer/JSON/text marshaler for the raw byte slice.
func (pretty Pretty) Raw(v []byte) RawV {
	return RawV{
		v:        v,
		prettier: pretty,
	}
}

type RawV struct {
	v        []byte
	prettier Pretty
}

func (v RawV) String() string {
	if v.v == nil {
		return v.prettier.nil
	}
	return string(v.v)
}

func (v RawV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return []byte(v.prettier.nil), nil
	}
	return v.v, nil
}

func (v RawV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
