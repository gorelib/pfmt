// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Bytesps returns stringer/JSON/text marshaler for slice of pointers to byte slice type.
func Bytesps(a []*[]byte) ByteSPS { return New().Bytesps(a) }

// Bytesps returns stringer/JSON/text marshaler for slice of pointers to byte slice type.
func (pretty Pretty) Bytesps(a []*[]byte) ByteSPS {
	return ByteSPS{
		a:        a,
		prettier: pretty,
	}
}

type ByteSPS struct {
	a        []*[]byte
	prettier Pretty
}

func (a ByteSPS) String() string {
	t, _ := a.MarshalText()
	return string(t)
}

func (a ByteSPS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Bytesp(p).MarshalText()
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

func (a ByteSPS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Bytesp(p).MarshalJSON()
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
