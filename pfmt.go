// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"fmt"
	"sync"
)

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, errno error) {
	return fmt.Println(Anys(a))
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...interface{}) string {
	return Anys(a).String()
}

// Prettier.
type Prettier interface {
	fmt.Stringer
	KV
}

func New(opts ...Option) Pretty {
	pretty := Pretty{
		separator: " ",
		nil:       "null",
		true:      "true",
		false:     "false",
		empty:     "",
		stack:     3,
	}
	for _, opt := range opts {
		opt(&pretty)
	}
	return pretty
}

// Pretty.
type Pretty struct {
	separator string
	nil       string
	true      string
	false     string
	empty     string
	stack     int
}

// Formatter returns formatter.
func Formatter(v interface{}) fmt.Formatter {
	return Fmt{v: v}
}

// Format method controls how State and rune are interpreted,
// and may call Sprint(f) or Fprint(f) etc. to generate its output.
func (fm Fmt) Format(f fmt.State, c rune) {
	_, err := f.Write([]byte(Any(fm.v).String()))
	if err != nil {
		_, _ = f.Write([]byte(err.Error()))
	}
}

// Formatter implements fmt formatter interface.
type Fmt struct {
	v interface{}
}

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// Option changes prettier configuration.
type Option func(*Pretty)

// WithSeparator sets separator.
func WithSeparator(separator string) Option {
	return func(pretty *Pretty) { pretty.separator = separator }
}

// WithNil sets nil.
func WithNil(null string) Option {
	return func(pretty *Pretty) { pretty.nil = null }
}

// WithTrue sets true.
func WithTrue(t string) Option {
	return func(pretty *Pretty) { pretty.true = t }
}

// WithFalse sets false.
func WithFalse(f string) Option {
	return func(pretty *Pretty) { pretty.false = f }
}

// WithEmpty sets empty.
func WithEmpty(empty string) Option {
	return func(pretty *Pretty) { pretty.empty = empty }
}

// WithStackDepth sets stack depth.
func WithStackDepth(depth int) Option {
	return func(pretty *Pretty) { pretty.stack = depth }
}
