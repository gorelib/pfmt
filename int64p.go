// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Int64p returns stringer/JSON/text marshaler for the int64 pointer type.
func Int64p(p *int64) Int64P { return Int64P{p: p} }

type Int64P struct{ p *int64 }

func (p Int64P) String() string {
	if p.p == nil {
		return "null"
	}
	return int64V{V: *p.p}.String()
}

func (p Int64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p Int64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
