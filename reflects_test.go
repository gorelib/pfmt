// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt_test

import (
	"testing"
)

// FIXME: test missing!!!
var MarshalReflectsTests = []marshalTests{}

func TestMarshalReflects(t *testing.T) {
	testMarshal(t, MarshalReflectsTests)
}
