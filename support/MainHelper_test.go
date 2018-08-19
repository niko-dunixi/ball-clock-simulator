package support

import (
	"testing"
	"strings"
)

func TestGetOneUserInput(t *testing.T) {
	inputs, _ := GetValidUserInput(strings.NewReader("123"))
	assertIntSlicesEqual(t, inputs, []int{123})
}

func TestGetTwoUserInput(t *testing.T) {
	inputs, _ := GetValidUserInput(strings.NewReader("123 456"))
	assertIntSlicesEqual(t, inputs, []int{123, 456})
}

func TestGetInvalidEmptyInput(t *testing.T) {
	_, err := GetValidUserInput(strings.NewReader(""))
	if err != InvalidInputCount {
		t.FailNow()
	}
}

func TestGetInvalidUserInput(t *testing.T) {
	_, err := GetValidUserInput(strings.NewReader("BadInput"))
	if err != InvalidInputString {
		t.FailNow()
	}
}

func TestGetThreeInvalidUserInput(t *testing.T) {
	_, err := GetValidUserInput(strings.NewReader("123 456 7"))
	if err != InvalidInputCount {
		fmtPrintlnErr(err)
		t.FailNow()
	}
}

func TestGetBelowRangeInvalidUserInput(t *testing.T) {
	_, err := GetValidUserInput(strings.NewReader("26"))
	if err != InvalidBallCount {
		fmtPrintlnErr(err)
		t.FailNow()
	}
}

func TestGetAboveRangeInvalidUserInput(t *testing.T) {
	_, err := GetValidUserInput(strings.NewReader("128"))
	if err != InvalidBallCount {
		fmtPrintlnErr(err)
		t.FailNow()
	}
}
