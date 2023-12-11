package day_four

import (
	"math"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

var cardExpression = regexp.MustCompile(`Card[ ]+\d+: (?P<winningNumbers>.*)\|(?P<ownedNumbers>.*)`)

func matchingNumbers(winningStr string, ownedStr string) uint16 {
	winning := mapset.NewSet[string]()
	owned := mapset.NewSet[string]()

	for _, winningNumber := range strings.Split(winningStr, " ") {
		if winningNumber != "" {
			winning.Add(winningNumber)
		}
	}
	for _, ownedNumber := range strings.Split(ownedStr, " ") {
		if ownedNumber != "" {
			owned.Add(ownedNumber)
		}
	}

	return uint16(winning.Intersect(owned).Cardinality())
}

func StarOne(input string) uint16 {
	lines := strings.Split(input, "\n")
	result := uint16(0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		cardMatch := cardExpression.FindStringSubmatch(line)
		points := matchingNumbers(cardMatch[1], cardMatch[2])

		if points == 0 {
			continue
		}
		result += uint16(math.Pow(2, float64(points-1)))
	}

	return result
}

func StarTwo(input string) uint32 {
	lines := strings.Split(input, "\n")

	var filteredLines []string
	result := uint32(0)
	for _, line := range lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	repitions := make(map[uint16]uint32)

	for i := 0; i < len(filteredLines); i++ {
		repitions[uint16(i)] = 1
	}

	for lineIndex, line := range filteredLines {
		if line == "" {
			continue
		}

		cardMatch := cardExpression.FindStringSubmatch(line)
		extraCards := matchingNumbers(cardMatch[1], cardMatch[2])

		for i := uint16(1); i <= extraCards; i++ {
			if _, ok := repitions[uint16(lineIndex)+i]; ok {
				repitions[uint16(lineIndex)+i] += (repitions[uint16(lineIndex)])
			}
		}
	}

	for _, repition := range repitions {
		result += repition
	}

	return result
}
