package day_six

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regularExpression = regexp.MustCompile(`\d+`)

func winningOpportunities(times []uint64, distances []uint64) uint64 {
	var result uint64 = 1
	for index, time := range times {
		winningOpportunities := uint64(0)
		velocity := 1
		for i := 1; i < int(time); i++ {
			timeLeft := time - uint64(i)
			if velocity*int(timeLeft) > int(distances[index]) {
				winningOpportunities++
			}
			velocity++
		}
		if winningOpportunities > 0 {
			result *= winningOpportunities
		}
	}
	return result
}

func StarOne(input string) uint64 {
	lines := strings.Split(input, "\n")

	var times []uint64
	for _, timeStr := range regularExpression.FindAllStringSubmatch(lines[0], -1) {
		time, _ := strconv.Atoi(timeStr[0])
		times = append(times, uint64(time))
	}
	var distances []uint64
	for _, distanceStr := range regularExpression.FindAllStringSubmatch(lines[1], -1) {
		distance, _ := strconv.Atoi(distanceStr[0])
		distances = append(distances, uint64(distance))
	}

	return winningOpportunities(times, distances)
}

func StarTwo(input string) uint64 {
	lines := strings.Split(input, "\n")

	var wholleTimeStr = ""
	for _, timeStr := range regularExpression.FindAllStringSubmatch(lines[0], -1) {
		wholleTimeStr += timeStr[0]
	}
	var wholleDistanceStr = ""
	for _, distanceStr := range regularExpression.FindAllStringSubmatch(lines[1], -1) {
		wholleDistanceStr += distanceStr[0]
	}

	time, _ := strconv.Atoi(wholleTimeStr)
	distance, _ := strconv.Atoi(wholleDistanceStr)

	fmt.Println([]uint64{uint64(time)}, []uint64{uint64(distance)})
	return winningOpportunities([]uint64{uint64(time)}, []uint64{uint64(distance)})
}
