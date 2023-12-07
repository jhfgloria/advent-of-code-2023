package day_one

import (
	"regexp"
	"strconv"
	"strings"
)

var expression = regexp.MustCompile(`[a-zA-Z]*`)

func StarOne(input string) uint16 {
	lines := strings.Split(input, "\n")
	return sumLines(lines)
}

func StarTwo(input string) uint16 {
	numbers := numericReplacement(input)
	lines := strings.Split(numbers, "\n")
	return sumLines(lines)
}

func numericReplacement(word string) string {
	wordCopy := strings.Clone(word)
	wordCopy = strings.ReplaceAll(wordCopy, "one", "o1e")
	wordCopy = strings.ReplaceAll(wordCopy, "two", "t2o")
	wordCopy = strings.ReplaceAll(wordCopy, "three", "t3e")
	wordCopy = strings.ReplaceAll(wordCopy, "four", "f4r")
	wordCopy = strings.ReplaceAll(wordCopy, "five", "f5e")
	wordCopy = strings.ReplaceAll(wordCopy, "six", "s6x")
	wordCopy = strings.ReplaceAll(wordCopy, "seven", "s7n")
	wordCopy = strings.ReplaceAll(wordCopy, "eight", "e8t")
	wordCopy = strings.ReplaceAll(wordCopy, "nine", "n9e")

	return wordCopy
}

func sumLines(lines []string) uint16 {
	result := 0
	for _, line := range lines {
		_line := strings.Trim(expression.ReplaceAllString(line, ""), "\t")
		if len(_line) == 0 {
			// do nothing
		} else {
			first, _ := strconv.Atoi(_line[0:1])
			second, _ := strconv.Atoi(_line[len(_line)-1:])
			result += (first * 10) + second
		}
	}
	return uint16(result)
}
