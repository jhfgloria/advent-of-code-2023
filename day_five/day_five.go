package day_five

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type SeedRange struct {
	start     uint32
	rangeSize uint32
}

var seedsExpression = regexp.MustCompile(`seeds: (.*)`)
var mapExpression = regexp.MustCompile(`(.*) map:`)

func extractSeedsInfoSlice(seedsLine string) []uint32 {
	seedsStr := strings.Split(seedsExpression.FindStringSubmatch(seedsLine)[1], " ")
	var seeds []uint32
	for _, seedStr := range seedsStr {
		seedInt, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, uint32(seedInt))
	}
	return seeds
}

func extractSeedsInfo(seedsLine string) mapset.Set[uint32] {
	seeds := mapset.NewSet[uint32]()
	for _, seed := range extractSeedsInfoSlice(seedsLine) {
		seeds.Add(seed)
	}
	return seeds
}

func extractConversionMap(mapLine string) []uint32 {
	mapStr := strings.Split(mapLine, " ")
	conversionMap := make([]uint32, len(mapStr))
	for index, mapStr := range mapStr {
		mapInt, _ := strconv.Atoi(mapStr)
		conversionMap[index] = uint32(mapInt)
	}
	return conversionMap
}

func seedToLocation(lines []string, seeds mapset.Set[uint32]) uint32 {
	conversions := mapset.NewSet[uint32]()
	removals := mapset.NewSet[uint32]()

	for _, line := range lines[2:] {
		if mapExpression.MatchString(line) {
			// start processing map
			conversions.Clear()
			removals.Clear()
		} else if line == "" {
			// finish processing map
			seeds = seeds.Difference(removals)
			conversions = conversions.Union(seeds)
			seeds = conversions.Clone()
		} else {
			// process map
			conversionMap := extractConversionMap(line)
			for seed := range seeds.Iter() {
				if seed >= conversionMap[1] && seed <= conversionMap[1]+conversionMap[2]-1 {
					conversions.Add(seed + (conversionMap[0] - conversionMap[1]))
					removals.Add(seed)
				}
			}
		}
	}

	return slices.Min(seeds.ToSlice())
}

func StarOne(input string) uint32 {
	lines := strings.Split(input, "\n")
	seeds := extractSeedsInfo(lines[0])
	return seedToLocation(lines, seeds)
}

// func StarTwo(input string) uint32 {
// 	lines := strings.Split(input, "\n")
// 	seedsInfo := append([]uint32{}, extractSeedsInfoSlice(lines[0])...)
// 	seeds := mapset.NewSet[uint32]()

// 	for i := uint32(0); i < uint32(len(seedsInfo))-1; i += 2 {
// 		start := seedsInfo[i]
// 		rangeSize := seedsInfo[i+1]
// 		for j := start; j < start+rangeSize; j++ {
// 			seeds.Add(j)
// 		}
// 	}

// 	return seedToLocation(lines, seeds)
// }
