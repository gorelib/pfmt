// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Bool returns stringer/JSON/text marshaler for the bool type.
func Bool(v bool) BoolV { return New().Bool(v) }

// Bool returns stringer/JSON/text marshaler for the bool type.
func (pretty Pretty) Bool(v bool) BoolV {
	return BoolV{
		v:        v,
		prettier: pretty,
	}
}

type BoolV struct {
	v        bool
	prettier Pretty
}

func (v BoolV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v BoolV) MarshalText() ([]byte, error) {
	if v.v {
		return []byte(v.prettier.true), nil
	}
	return []byte(v.prettier.false), nil
}

func (v BoolV) MarshalJSON() ([]byte, error) {
	if v.v {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}
