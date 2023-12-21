package day_five

import "testing"

func TestExampleOne(t *testing.T) {
	expectation := uint64(35)
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	value := StarOne(input)

	if value != expectation {
		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
	}
}

func TestTwoExampleOne(t *testing.T) {
	expectation := uint64(86)
	input := `seeds: 55

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	value := StarOne(input)

	if value != expectation {
		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
	}
}

func TestExampleTwo(t *testing.T) {
	expectation := uint64(46)
	input := `seeds: 79 14

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	value := StarTwo(input)

	if value != expectation {
		t.Fatalf(`StarTwo(input) = %d, want %d, error`, value, expectation)
	}
}

func TestTwoExampleTwo(t *testing.T) {
	expectation := uint64(18)
	// 10 - 20
	input := `seeds: 10 11

seed-to-soil map:
50 14 4
90 10 4
`
	value := StarTwo(input)

	if value != expectation {
		t.Fatalf(`StarTwo(input) = %d, want %d, error`, value, expectation)
	}
}
