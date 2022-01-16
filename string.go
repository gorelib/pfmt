// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"

	"github.com/pprint/pfmt/pencode"
)

// String returns stringer/JSON/text marshaler for the string type.
func String(v string) StringV { return StringV{V: v} }

type StringV struct{ V string }

func (v StringV) String() string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := pencode.String(buf, v.V)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (v StringV) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	err := pencode.String(&buf, v.V)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (v StringV) MarshalJSON() ([]byte, error) {
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
