package pfmt

import (
	"fmt"
)

// Println it returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, errno error) {
	return fmt.Println(Sprint(a...))
}

// Sprint return string.
func Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}
