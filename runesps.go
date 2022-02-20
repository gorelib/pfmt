// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Runesps returns stringer/JSON/text marshaler for slice of pointers to rune slice type.
func Runesps(a []*[]rune) RuneSPS { return New().Runesps(a) }

// Runesps returns stringer/JSON/text marshaler for slice of pointers to rune slice type.
func (pretty Pretty) Runesps(a []*[]rune) RuneSPS {
	return RuneSPS{
		a:        a,
		prettier: pretty,
	}
}

type RuneSPS struct {
	a        []*[]rune
	prettier Pretty
}

func (a RuneSPS) String() string {
	t, _ := a.MarshalText()
	return string(t)
}

func (a RuneSPS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Runesp(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(a.prettier.separator)
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (a RuneSPS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Runesp(p).MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
