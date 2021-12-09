// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"strings"

	"github.com/pprint/pfmt/pencode"
)

// Strings returns stringer/JSON/text marshaler for the string slice type.
func Strings(s ...string) StringS { return StringS{s: s} }

type StringS struct{ s []string }

func (s StringS) String() string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := pencode.String(buf, strings.Join(s.s, " "))
	if err != nil {
		return ""
	}
	return buf.String()
}

func (s StringS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	err := pencode.String(&buf, strings.Join(s.s, " "))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s StringS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := stringV{V: v}.MarshalJSON()
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
