// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"fmt"
)

// JSON returns stringer/JSON/text marshaler for the KV slice type.
func JSON(s []KV) JSONV { return New().JSON(s) }

// JSON returns stringer/JSON/text marshaler for the KV slice type.
func (pretty Pretty) JSON(s []KV) JSONV {
	return JSONV{
		s:        s,
		prettier: pretty,
	}
}

type JSONV struct {
	s        []KV
	prettier Pretty
}

func (s JSONV) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s JSONV) MarshalText() ([]byte, error) {
	if s.s == nil {
		return nil, nil
	}
	var buf bytes.Buffer
	for i, j := range s.s {
		if i != 0 {
			buf.WriteString(s.prettier.separator)
		}
		if j == nil {
			buf.WriteString(s.prettier.nil)
			continue
		}
		k, err := j.MarshalText()
		if err != nil {
			return nil, err
		}
		if k == nil {
			continue
		}
		v, err := j.MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(append(k, append([]byte(s.prettier.separator), v...)...))
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s JSONV) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, j := range s.s {
		if i != 0 {
			buf.WriteString(",")
		}
		if j == nil {
			buf.WriteString(`"":""`)
			continue
		}
		k, err := j.MarshalText()
		if err != nil {
			return nil, err
		}
		if k == nil {
			k = []byte("")
		}
		v, err := j.MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(append([]byte(`"`), append(k, append([]byte(`":`), v...)...)...))
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("}")
	fmt.Println(buf.String())
	return buf.Bytes(), nil
}
