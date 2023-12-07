package day_three

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var symbolExpression = regexp.MustCompile(`[^\w\.]`)
var partNumberExpression = regexp.MustCompile(`\d`)

func countAsPartNumber(lines []string, line []rune, number int, previousLineIndex int, nextLineIndex int, previousCharacterIndex int, nextCharacterIndex int, start int, end int) uint16 {
	if previousCharacterIndex >= 0 {
		if symbolExpression.MatchString(string(line[previousCharacterIndex])) {
			return uint16(number)
		}
	}

	if nextCharacterIndex < len(line) {
		if symbolExpression.MatchString(string(line[nextCharacterIndex])) {
			return uint16(number)
		}
	}

	if previousLineIndex > 0 {
		for i := max(0, start-1); i < min(end+1, len(line)-1); i++ {
			if symbolExpression.MatchString(string(lines[previousLineIndex][i])) {
				return uint16(number)
			}
		}
	}

	if nextLineIndex < len(lines)-1 {
		for i := max(0, start-1); i < min(end+1, len(line)-1); i++ {
			if symbolExpression.MatchString(string(lines[nextLineIndex][i])) {
				return uint16(number)
			}
		}
	}

	return 0
}

func StarOne(input string) uint16 {
	lines := strings.Split(input, "\n")
	result := uint16(0)

	for lineIndex, line := range lines {
		fmt.Println(line)

		strNumber := ""
		start := -1

		for charIndex, character := range line {
			if partNumberExpression.MatchString(string(character)) {
				strNumber += string(character)
				if start == -1 {
					start = charIndex
				}
			}
			if strNumber != "" && (!partNumberExpression.MatchString(string(character)) || charIndex == len(line)-1) {
				number, _ := strconv.Atoi(strNumber)
				previousCharacterIndex := start - 1
				nextCharacterIndex := charIndex
				previousLineIndex := lineIndex - 1
				nextLineIndex := lineIndex + 1

				result += countAsPartNumber(lines, []rune(line), number, previousLineIndex, nextLineIndex, previousCharacterIndex, nextCharacterIndex, start, charIndex)

				strNumber = ""
				start = -1
			}
		}
	}

	return result
}
