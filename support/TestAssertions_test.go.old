package support

import (
	"testing"
	"fmt"
)

func assertIntSlicesEqual(t *testing.T, actual, expected []int) {
	// Verify that either both or neither are nil
	if (actual == nil) != (expected == nil) {
		t.FailNow()
	}
	// Verify they have the same size, no need for
	// an O(n) operation if we can know it's fundamentally
	// not equivalent
	if len(actual) != len(expected) {
		t.FailNow()
	}
	// Verify each item within are equal
	for i := range actual {
		if actual[i] != expected[i] {
			t.FailNow()
		}
	}
}

func fmtPrintlnErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Attempted to log error, but it was nil")
	}
}
