package day_six

import "testing"

func TestExampleOne(t *testing.T) {
	expectation := uint64(288)
	input := `Time:      7  15   30
Distance:  9  40  200`
	value := StarOne(input)

	if value != expectation {
		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
	}
}

func TestTwoExampleOne(t *testing.T) {
	expectation := uint64(71503)
	input := `Time:      71530
Distance:  940200`
	value := StarOne(input)

	if value != expectation {
		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
	}
}

func TestExampleTwo(t *testing.T) {
	expectation := uint64(71503)
	input := `Time:      7  15   30
Distance:  9  40  200`
	value := StarTwo(input)

	if value != expectation {
		t.Fatalf(`StarTwo(input) = %d, want %d, error`, value, expectation)
	}
}
