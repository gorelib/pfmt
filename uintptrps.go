// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uintptrps returns stringer/JSON/text marshaler for the slice of a pointer to a uintptr type.
func Uintptrps(a []*uintptr) UintptrPS { return New().Uintptrps(a) }

// Uintptrps returns stringer/JSON/text marshaler for the slice of a pointer to a uintptr type.
func (pretty Pretty) Uintptrps(a []*uintptr) UintptrPS {
	return UintptrPS{
		a:        a,
		prettier: pretty,
	}
}

type UintptrPS struct {
	a        []*uintptr
	prettier Pretty
}

func (a UintptrPS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a UintptrPS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Uintptrp(p).MarshalText()
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

func (a UintptrPS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Uintptrp(p).MarshalJSON()
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
