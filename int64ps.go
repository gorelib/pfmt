// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int64ps returns stringer/JSON/text marshaler for the int64 pointer slice type.
func Int64ps(a []*int64) Int64PS { return Int64PS{a: a} }

type Int64PS struct{ a []*int64 }

func (a Int64PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Int64PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := Int64p(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(" ")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (a Int64PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := Int64p(p).MarshalJSON()
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
