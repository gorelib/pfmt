// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/pprint/pfmt"
)

func TestSprint(t *testing.T) {
	var nilStringPtr *string
	var nilTimePtr *time.Time

	testInt := 42
	testStr := "Hello, World!"

	testTime, err := time.Parse(time.RFC3339, "2022-12-07T18:56:49Z")
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
			want: "2022-12-07T18:56:49Z",
		},
		{
			name: "time ptr",
			line: line(),
			args: []interface{}{&testTime},
			want: "2022-12-07T18:56:49Z",
		},
		{
			name: "nil time ptr",
			line: line(),
			args: []interface{}{nilTimePtr},
			want: "null",
		},
		{
			name: "array",
			line: line(),
			args: []interface{}{[3]int{1, 2, 3}},
			want: "[1 2 3]",
		},
		{
			name: "slice",
			line: line(),
			args: []interface{}{[]int{1, 2, 3}},
			want: "1 2 3",
		},
		{
			name: "struct",
			line: line(),
			args: []interface{}{http.Request{}},
			want: "http.Request{Method: URL:null Proto: ProtoMajor:0 ProtoMinor:0 Header:null Body:null GetBody:null ContentLength:0 TransferEncoding:null Close:false Host: Form:null PostForm:null MultipartForm:null Trailer:null RemoteAddr: RequestURI: TLS:null Cancel:null Response:null}",
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
