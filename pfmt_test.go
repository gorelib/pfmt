package pfmt_test

import (
	"testing"
	"time"

	"github.com/pprint/pfmt"
)

func TestSprint(t *testing.T) {
	var nilStringPtr *string
	var nilTimePtr *time.Time

	testInt := 42
	testStr := "Hello, World!"

	testTime, err := time.Parse(time.RFC3339, "2021-12-07T18:56:49Z")
	if err != nil {
		t.Fatalf("time parse: %s", err)
	}

	tests := []struct {
		name string
		line string
		args []interface{}
		want string
	}{
		{
			name: "nil",
			line: line(),
			args: []interface{}{nil},
			want: "null",
		},
		{
			name: "int",
			line: line(),
			args: []interface{}{testInt},
			want: "42",
		},
		{
			name: "string",
			line: line(),
			args: []interface{}{testStr},
			want: "Hello, World!",
		},
		{
			name: "string ptr",
			line: line(),
			args: []interface{}{&testStr},
			want: "Hello, World!",
		},
		{
			name: "nil string ptr",
			line: line(),
			args: []interface{}{nilStringPtr},
			want: "null",
		},
		{
			name: "time",
			line: line(),
			args: []interface{}{testTime},
			want: "2021-12-07T18:56:49Z",
		},
		{
			name: "time ptr",
			line: line(),
			args: []interface{}{&testTime},
			want: "2021-12-07T18:56:49Z",
		},
		{
			name: "nil time ptr",
			line: line(),
			args: []interface{}{nilTimePtr},
			want: "null",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+tt.name, func(t *testing.T) {
			t.Parallel()
			got := pfmt.Sprint(tt.args...)
			if got != tt.want {
				t.Errorf("\nwant: %s\ngot:  %s\ntest: %s", tt.want, got, tt.line)
			}
		})
	}
}
