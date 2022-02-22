// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uintps returns stringer/JSON/text marshaler for the uint pointer slice type.
func Uintps(a []*uint) UintPS { return New().Uintps(a) }

// Uintps returns stringer/JSON/text marshaler for the uint pointer slice type.
func (pretty Pretty) Uintps(a []*uint) UintPS {
	return UintPS{
		a:        a,
		prettier: pretty,
	}
}

type UintPS struct {
	a        []*uint
	prettier Pretty
}

func (a UintPS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a UintPS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Uintp(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(a.prettier.separator)
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func (a UintPS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Uintp(p).MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		buf.Write(b)
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
