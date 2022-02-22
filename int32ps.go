// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int32ps returns stringer/JSON/text marshaler for the int32 pointer slice type.
func Int32ps(a []*int32) Int32PS { return New().Int32ps(a) }

// Int32ps returns stringer/JSON/text marshaler for the int32 pointer slice type.
func (pretty Pretty) Int32ps(a []*int32) Int32PS {
	return Int32PS{
		a:        a,
		prettier: pretty,
	}
}

type Int32PS struct {
	a        []*int32
	prettier Pretty
}

func (a Int32PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Int32PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Int32p(p).MarshalText()
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

func (a Int32PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Int32p(p).MarshalJSON()
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
