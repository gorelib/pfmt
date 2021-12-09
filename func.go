// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

// Func returns stringer/JSON/text marshaler for the custom function type.
func Func(v func() KV) funcV { return funcV{V: v} }

type funcV struct{ V func() KV }

func (v funcV) String() string {
	b, _ := v.V().MarshalText()
	return string(b)
}

func (v funcV) MarshalText() ([]byte, error) {
	return v.V().MarshalText()
}

func (v funcV) MarshalJSON() ([]byte, error) {
	return v.V().MarshalJSON()
}
