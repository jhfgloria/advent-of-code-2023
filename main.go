package main

import (
	"aoc/2023/day_five"
	"aoc/2023/requests"
	"fmt"
)

func main() {
	input := requests.GetExercise("53616c7465645f5ff687033561bd0e29fa45f5decbefb866b175133376a0981e9c76ee345e33947135828ccc716dff1d09ad3e53f648bbc5e31b358790966b65", 5)
	result := day_five.StarTwo(input)
	fmt.Println(result)
}
