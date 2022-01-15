// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pprint/pfmt"
)

var MarshalErrorsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"error slice": pfmt.Errors([]error{errors.New("something went wrong"), errors.New("we have a problem")})},
		want:     "something went wrong we have a problem",
		wantText: "something went wrong we have a problem",
		wantJSON: `{
			"error slice":["something went wrong","we have a problem"]
		}`,
	},
	{
		line:  line(),
		input: map[string]json.Marshaler{"nil errors": pfmt.Errors([]error{nil, nil})},
		wantJSON: `{
			"nil errors":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without errors": pfmt.Errors(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"without errors":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"slice of any errors": pfmt.Anys([]interface{}{err, err2})}
		}(),
		want:     "something went wrong we have a problem",
		wantText: "something went wrong we have a problem",
		wantJSON: `{
			"slice of any errors":["something went wrong","we have a problem"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"slice of reflect of errors": pfmt.Reflects([]interface{}{err, err2})}
		}(),
		want:     "{something went wrong} {we have a problem}",
		wantText: "{something went wrong} {we have a problem}",
		wantJSON: `{
			"slice of reflect of errors":[{},{}]
		}`,
	},
}

func TestMarshalErrors(t *testing.T) {
	testMarshal(t, MarshalErrorsTests)
}
