package day_three

import "testing"

func TestExampleOne(t *testing.T) {
	expectation := uint16(4400)
	input := `467..114..
...+......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$./....
.664.598..
........1$
..........
........$1
..........
1@........
..........
=1...-....
......35..
`
	value := StarOne(input)

	if value != expectation {
		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
	}
}

// ..............423....688..934............970................................95.728..........896...113..................153..972.............
// ...122..................*.....*..........................919..509*..........&...@.........../...........................+.......*...........
// ....+..........259....698..373.992.52.674.........................781...22........130.584.....-...%399.......777.................266........

// func TestExampleTwo(t *testing.T) {
// 	expectation := uint16(2286)
// 	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// 	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// 	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// 	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// 	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
// 	value := StarTwo(input)

// 	if value != expectation {
// 		t.Fatalf(`StarOne(input) = %d, want %d, error`, value, expectation)
// 	}
// }
