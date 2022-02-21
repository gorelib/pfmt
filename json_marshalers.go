// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"encoding/json"
)

// JSONMarshalers returns stringer/JSON/text marshaler for the JSON marshaler slice type.
func JSONMarshalers(s []json.Marshaler) JSONMarshalerS { return New().JSONMarshalers(s) }

// JSONMarshalers returns stringer/JSON/text marshaler for the JSON marshaler slice type.
func (pretty Pretty) JSONMarshalers(s []json.Marshaler) JSONMarshalerS {
	return JSONMarshalerS{
		s:        s,
		prettier: pretty,
	}
}

type JSONMarshalerS struct {
	s        []json.Marshaler
	prettier Pretty
}

func (s JSONMarshalerS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s JSONMarshalerS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	return s.MarshalJSON()
}

func (s JSONMarshalerS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		if i != 0 {
			buf.WriteString(",")
		}
		if v == nil {
			buf.WriteString("null")
			continue
		}
		b, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
