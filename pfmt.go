package pfmt

import (
	"encoding"
	"encoding/json"
	"fmt"
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

// KV is a key-value pair.
type KV interface {
	encoding.TextMarshaler
	json.Marshaler
}
