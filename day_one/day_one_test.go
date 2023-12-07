package day_one

import "testing"

func TestExampleOne(t *testing.T) {
	expectation := uint16(142)
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	value := StarOne(input)

	if value != expectation {
		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
	}
}

func TestExampleTwo(t *testing.T) {
	expectation := uint16(281)
	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`
	value := StarTwo(input)

	if value != expectation {
		t.Fatalf(`StarTwo(input) = %d, want %d, error`, value, expectation)
	}
}
