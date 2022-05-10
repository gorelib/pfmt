// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding"

	"github.com/gorelib/pfmt/pencode"
)

// Text returns stringer/JSON/text marshaler for the encoding.TextMarshaler type.
func Text(v encoding.TextMarshaler) TextV { return New().Text(v) }

// Text returns stringer/JSON/text marshaler for the encoding.TextMarshaler type.
func (pretty Pretty) Text(v encoding.TextMarshaler) TextV {
	return TextV{
		v:        v,
		prettier: pretty,
	}
}

type TextV struct {
	v        encoding.TextMarshaler
	prettier Pretty
}

func (v TextV) String() string {
	if v.v == nil {
		return v.prettier.empty
	}
	b, err := v.v.MarshalText()
	if err != nil {
		return v.prettier.empty
	}
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	err = pencode.Bytes(buf, b)
	if err != nil {
		return v.prettier.empty
	}
	return buf.String()
}

func (v TextV) MarshalText() ([]byte, error) {
	if v.v == nil {
		return nil, nil
	}
	b, err := v.v.MarshalText()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	err = pencode.Bytes(&buf, b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (v TextV) MarshalJSON() ([]byte, error) {
	if v.v == nil {
		return []byte("null"), nil
	}
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
