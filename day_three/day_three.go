package day_three

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var symbolExpression = regexp.MustCompile(`[^\w\.]`)
var partNumberExpression = regexp.MustCompile(`\d`)

func adjacentSymbol(regex *regexp.Regexp, lines []string, line []rune, number int, lineIndex int, previousCharacterIndex int, nextCharacterIndex int, start int, end int) string {
	if previousCharacterIndex >= 0 {
		if regex.MatchString(string(line[previousCharacterIndex])) {
			return fmt.Sprintf("%d:%d", lineIndex, previousCharacterIndex)
		}
	}

	if nextCharacterIndex < len(line) {
		if regex.MatchString(string(line[nextCharacterIndex])) {
			return fmt.Sprintf("%d:%d", lineIndex, nextCharacterIndex)
		}
	}

	if (lineIndex - 1) > 0 {
		for i := max(0, start-1); i < min(end+2, len(line)-1); i++ {
			if regex.MatchString(string(lines[(lineIndex - 1)][i])) {
				return fmt.Sprintf("%d:%d", lineIndex-1, i)
			}
		}
	}

	if (lineIndex+1) <= len(lines)-1 && lines[(lineIndex+1)] != "" {
		for i := max(0, start-1); i < min(end+2, len(line)-1); i++ {
			if regex.MatchString(string(lines[(lineIndex + 1)][i])) {
				return fmt.Sprintf("%d:%d", lineIndex+1, i)
			}
		}
	}

	return ""
}

func StarOne(input string) uint32 {
	lines := strings.Split(input, "\n")
	result := uint32(0)

	for lineIndex, line := range lines {
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

				if adjacentSymbol(symbolExpression, lines, []rune(line), number, lineIndex, previousCharacterIndex, nextCharacterIndex, start, charIndex-1) != "" {
					result += uint32(number)
				}

				strNumber = ""
				start = -1
			}
		}
	}

	return result
}

func StarTwo(input string) uint64 {
	gearExpression := regexp.MustCompile(`[\*]`)
	dictionary := make(map[string][]uint64)
	lines := strings.Split(input, "\n")
	result := uint64(0)

	for lineIndex, line := range lines {
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

				symbolIndex := adjacentSymbol(gearExpression, lines, []rune(line), number, lineIndex, previousCharacterIndex, nextCharacterIndex, start, charIndex-1)

				if symbolIndex != "" {
					if dictionary[symbolIndex] == nil {
						dictionary[symbolIndex] = []uint64{uint64(number)}
					} else {
						dictionary[symbolIndex] = append(dictionary[symbolIndex], uint64(number))
					}
				}

				strNumber = ""
				start = -1
			}
		}
	}

	for _, value := range dictionary {
		if len(value) == 2 {
			result += value[0] * value[1]
		}
	}

	return result
}
