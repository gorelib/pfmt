// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"time"
)

// Timeps returns stringer/JSON/text marshaler for the time pointer slice type.
func Timeps(a []*time.Time) TimePS { return New().Timeps(a) }

// Timeps returns stringer/JSON/text marshaler for the time pointer slice type.
func (pretty Pretty) Timeps(a []*time.Time) TimePS {
	return TimePS{
		a:        a,
		prettier: pretty,
	}
}

type TimePS struct {
	a        []*time.Time
	prettier Pretty
}

func (a TimePS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a TimePS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Timep(p).MarshalText()
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

func (a TimePS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Timep(p).MarshalJSON()
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
