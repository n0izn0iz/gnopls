// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import "iter"

func _(seq1 iter.Seq[int], seq2 iter.Seq2[int, int]) {
	// range-over-func is (once again) consistent with other types (#65236)
	for _ = range "" { // want "simplify range expression"
	}
	for _ = range seq1 { // want `simplify range expression`
	}
	for _, v := range seq2 { // silence
		_ = v
	}
	for _, _ = range seq2 { // want `simplify range expression`
	}
}
