// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package useany_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/gnoverse/gnopls/internal/analysis/useany"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, useany.Analyzer, "a")
}
