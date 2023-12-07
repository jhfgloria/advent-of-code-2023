package day_two

import (
	"regexp"
	"strconv"
	"strings"
)

var gameExpression = regexp.MustCompile(`Game (?P<id>\d+):`)
var blueExpression = regexp.MustCompile(`(?P<blue_count>\d+) blue`)
var redExpression = regexp.MustCompile(`(?P<red_count>\d+) red`)
var greenExpression = regexp.MustCompile(`(?P<green_count>\d+) green`)

var redCubes = uint8(12)
var greenCubes = uint8(13)
var blueCubes = uint8(14)

func StarOne(input string) uint16 {
	lines := strings.Split(input, "\n")
	result := 0

	for _, line := range lines {
		gameMatch := gameExpression.FindStringSubmatch(line)
		blueMatch := blueExpression.FindAllStringSubmatch(line, -1)
		redMatch := redExpression.FindAllStringSubmatch(line, -1)
		greenMatch := greenExpression.FindAllStringSubmatch(line, -1)

		if len(gameMatch) == 0 {
			break
		}

		if !overLimit(blueMatch, blueCubes) && !overLimit(greenMatch, greenCubes) && !overLimit(redMatch, redCubes) {
			gameID, _ := strconv.Atoi(gameMatch[1])
			result += gameID
		}
	}

	return uint16(result)
}

func StarTwo(input string) uint16 {
	lines := strings.Split(input, "\n")
	result := uint16(0)

	for _, line := range lines {
		gameMatch := gameExpression.FindStringSubmatch(line)
		blueMatch := blueExpression.FindAllStringSubmatch(line, -1)
		redMatch := redExpression.FindAllStringSubmatch(line, -1)
		greenMatch := greenExpression.FindAllStringSubmatch(line, -1)

		if len(gameMatch) == 0 {
			break
		}

		result += moreCubes(blueMatch) * moreCubes(redMatch) * moreCubes(greenMatch)
	}

	return result
}

func overLimit(list [][]string, limit uint8) bool {
	for _, v := range list {
		value, _ := strconv.Atoi(v[1])
		if uint8(value) > limit {
			return true
		}
	}
	return false
}

func moreCubes(list [][]string) uint16 {
	cubes := uint16(0)
	for _, v := range list {
		value, _ := strconv.Atoi(v[1])
		cubes = max(uint16(value), cubes)
	}
	return cubes
}
