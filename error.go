// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"

	"github.com/pprint/pfmt/pencode"
)

// Error returns stringer/JSON/text marshaler for the error type.
func Error(v error) errorV { return errorV{V: v} }

type errorV struct{ V error }

func (v errorV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v errorV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer

	err := pencode.String(&buf, v.V.Error())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v errorV) MarshalJSON() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
