// =====================================================================================================================
// == LICENSE:       Copyright (c) 2024 Kevin De Coninck
// ==
// ==                Permission is hereby granted, free of charge, to any person
// ==                obtaining a copy of this software and associated documentation
// ==                files (the "Software"), to deal in the Software without
// ==                restriction, including without limitation the rights to use,
// ==                copy, modify, merge, publish, distribute, sublicense, and/or sell
// ==                copies of the Software, and to permit persons to whom the
// ==                Software is furnished to do so, subject to the following
// ==                conditions:
// ==
// ==                The above copyright notice and this permission notice shall be
// ==                included in all copies or substantial portions of the Software.
// ==
// ==                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// ==                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// ==                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// ==                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// ==                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// ==                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// ==                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// ==                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Verify and measure the performance of the public API of the "assert" package.
package assert_test

import (
	"fmt"
	"testing"

	"github.com/kdeconinck/auditr/internal/pkg/assert"
)

// Wraps the testing.TB struct and add a field for storing the failure message.
type tbStub struct {
	testing.TB
	failureMsg string
}

// Fatalf flags tb as failed and formats args using fmt.Sprintf and stores the result in t.
func (tb *tbStub) Fatalf(format string, args ...any) {
	tb.failureMsg = fmt.Sprintf(format, args...)
}

// UT: Compare 2 values for equality.
func TestEqual(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		gotInput, wantInput any
		nameInput           string
		want                string
	}{
		"When `got` and `want` are NOT equal.": {
			gotInput: false, wantInput: true,
			nameInput: "Compare `false` against `true`",
			want:      "Compare `false` against `true` = false, want true",
		},
		"When `got` and `want` are equal.": {
			gotInput: true, wantInput: true,
			nameInput: "Compare `true` against `true`",
		},
		"When comparing `got` against `nil`.": {
			gotInput: true, wantInput: nil,
			nameInput: "Compare `true` against `<nil>`",
			want:      "Compare `true` against `<nil>` = true, want <nil>",
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			tc := tc     // Rebind the `tc` variable. Required to support parallel exceution.
			t.Parallel() // Enable parallel execution.

			// ARRANGE.
			testingTB := &tbStub{TB: t}

			// ACT.
			assert.Equal(testingTB, tc.gotInput, tc.wantInput, tc.nameInput)

			// ASSERT.
			if testingTB.failureMsg != tc.want {
				t.Fatalf("Failure message = \"%s\", want \"%s\"", testingTB.failureMsg, tc.want)
			}
		})
	}

	t.Run("When `got` and `want` are NOT equal and using a custom message", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// ARRANGE.
		testingTB := &tbStub{TB: t}

		// ACT."Test"
		assert.Equal(testingTB, false, true, "", "UT Failed: \"Compare `false` against `true`\" - got %t, want %t.",
			false, true)

		// ASSERT.
		if testingTB.failureMsg != "UT Failed: \"Compare `false` against `true`\" - got false, want true." {
			t.Fatalf("Failure message = \"%s\", want \"%s\"", testingTB.failureMsg, "UT Failed: \"Compare `false` "+
				"against `true`\" - got false, want true.")
		}
	})
}
