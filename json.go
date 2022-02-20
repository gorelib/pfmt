// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// JSON returns stringer/JSON/text marshaler for the KV slice type.
func JSON(s ...KV) JSONV { return JSONV{s: s} }

type JSONV struct {
	s []KV
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
	var tail bool
	for _, s := range s.s {
		if s == nil {
			continue
		}
		k, err := s.MarshalText()
		if err != nil {
			return nil, err
		}
		if k == nil {
			continue
		}
		v, err := s.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if tail {
			buf.WriteString(" ")
		}
		_, err = buf.Write(append(k, append([]byte(" "), v...)...))
		if err != nil {
			return nil, err
		}
		tail = true
	}
	return buf.Bytes(), nil
}

func (s JSONV) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("{")
	var tail bool
	for _, s := range s.s {
		if s == nil {
			continue
		}
		k, err := s.MarshalText()
		if err != nil {
			return nil, err
		}
		if k == nil {
			continue
		}
		v, err := s.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if tail {
			buf.WriteString(",")
		}
		_, err = buf.Write(append([]byte(`"`), append(k, append([]byte(`":`), v...)...)...))
		if err != nil {
			return nil, err
		}
		tail = true
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}
