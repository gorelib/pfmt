// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding"

	"github.com/pprint/pfmt/pencode"
)

// Text returns stringer/JSON/text marshaler for the encoding.TextMarshaler type.
func Text(v encoding.TextMarshaler) textV { return textV{V: v} }

type textV struct{ V encoding.TextMarshaler }

func (v textV) String() string {
	if v.V == nil {
		return ""
	}
	b, err := v.V.MarshalText()
	if err != nil {
		return ""
	}
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err = pencode.Bytes(buf, b)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (v textV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return nil, nil
	}
	b, err := v.V.MarshalText()
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

func (v textV) MarshalJSON() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
