package pfmt

import (
	"encoding"
	"encoding/json"
	"fmt"
)

// Println it returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, errno error) {
	return fmt.Println(Anys(a...))
}

// Sprint return string.
func Sprint(a ...interface{}) string {
	return Anys(a...).String()
}

// Formatter returns string.
func Formatter(v interface{}) fmt.Formatter {
	return Fmt{v: v}
}

func (fm Fmt) Format(f fmt.State, c rune) {
	_, err := f.Write([]byte(Any(fm.v).String()))
	if err != nil {
		_, _ = f.Write([]byte(err.Error()))
	}
}

type Fmt struct {
	v interface{}
}

// KV is a key-value pair.
type KV interface {
	encoding.TextMarshaler
	json.Marshaler
}
