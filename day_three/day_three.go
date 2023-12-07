package day_three

import (
	"fmt"
	"regexp"
	"strings"
)

type Line struct {
	Index       uint16
	Symbols     []Symbol
	PartNumbers []PartNumber
}

type Symbol struct {
	Position uint8
}

type PartNumber struct {
	Number uint16
	Start  uint8
	End    uint8
}

var symbolExpression = regexp.MustCompile(`[^\w\.]`)

func StarOne(input string) uint16 {
	lines := strings.Split(input, "\n")
	var allLines []Line

	for _, l := range lines {
		chars := []rune(l)
		line := Line{}
		for i := 0; i < len(chars); i++ {
			if symbolExpression.MatchString(string(chars[i])) {
				line.Index = uint16(i)
				fmt.Println("SYMBOL")
			}
		}
	}

	return uint16(0)
}
