package pfmt

import (
	"encoding"
	"encoding/json"
	"fmt"
	"time"
)

// Formatter returns string.
func Formatter(v interface{}) string {
	return Sprint(v)
}

// Println it returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, errno error) {
	return fmt.Println(Sprint(a...))
}

// Sprint return string.
func Sprint(a ...interface{}) string {
	for i := 0; i < len(a); i++ {
		if a[i] == nil {
			continue
		}

		switch v := a[i].(type) {
		case *string:
			a[i] = "*" + *v

		case *time.Time:
			a[i] = "*" + v.String()
		}
	}
	return fmt.Sprint(a...)
}

// KV is a key-value pair.
type KV interface {
	encoding.TextMarshaler
	json.Marshaler
}
