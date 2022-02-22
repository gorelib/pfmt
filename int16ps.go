// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int16ps returns stringer/JSON/text marshaler for the int16 pointer slice type.
func Int16ps(a []*int16) Int16PS { return New().Int16ps(a) }

// Int16ps returns stringer/JSON/text marshaler for the int16 pointer slice type.
func (pretty Pretty) Int16ps(a []*int16) Int16PS {
	return Int16PS{
		a:        a,
		prettier: pretty,
	}
}

type Int16PS struct {
	a        []*int16
	prettier Pretty
}

func (a Int16PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Int16PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Int16p(p).MarshalText()
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

func (a Int16PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Int16p(p).MarshalJSON()
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
