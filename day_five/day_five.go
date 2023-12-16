package day_five

import (
	"fmt"
	"regexp"
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

func seedToLocation(lines []string, seeds mapset.Set[SeedRange]) uint32 {
	conversions := mapset.NewSet[SeedRange]()
	removals := mapset.NewSet[SeedRange]()

	fmt.Println(seeds)

	for _, line := range lines[2:] {
		if mapExpression.MatchString(line) {
			// start processing map
			fmt.Println("start processing map")
			conversions.Clear()
			removals.Clear()
		} else if line == "" {
			// finish processing map
			fmt.Println("finish processing map")
			fmt.Println(seeds)
			seeds = seeds.Difference(removals)
			conversions = conversions.Union(seeds)
			seeds = conversions.Clone()
		} else {
			// process map
			fmt.Println("processing map")
			conversionMap := extractConversionMap(line)
			baseReplacement := conversionMap[0]
			higherLimit := conversionMap[1] + conversionMap[2] - 1
			lowerLimit := conversionMap[1]

			for seed := range seeds.Iter() {
				if seed.start > higherLimit {
					// above limits
					fmt.Println("above limits")
					conversions.Add(seed)
				} else if seed.start+seed.rangeSize-1 < lowerLimit {
					// below limits
					fmt.Println("below limits")
					conversions.Add(seed)
				} else {
					// in limits
					// fmt.Println("in limits")
					if seed.start < lowerLimit && seed.start+seed.rangeSize-1 <= higherLimit {
						// partially below limit and within range
						fmt.Println("partially below limit and within range")
						conversions.Add(SeedRange{start: seed.start, rangeSize: lowerLimit - seed.start})
						conversions.Add(SeedRange{start: baseReplacement, rangeSize: seed.rangeSize - lowerLimit})
						// WHEN I ADD A REPLACEMENT SEED, I NEED TO ADD THE ORIGINAL SEED TO THE REMOVALS
					}
					if seed.start >= lowerLimit && seed.start+seed.rangeSize-1 > higherLimit {
						// partially above limit and within range
						fmt.Println("partially above limit and within range")
						conversions.Add(SeedRange{start: seed.start + baseReplacement - lowerLimit, rangeSize: higherLimit - seed.start})
						conversions.Add(SeedRange{start: lowerLimit + 1, rangeSize: seed.start + seed.rangeSize + 1 - higherLimit})
					}
					if seed.start >= lowerLimit && seed.start+seed.rangeSize-1 <= higherLimit {
						// within range
						fmt.Println("within range", seed.start, seed.start+baseReplacement-lowerLimit)
						conversions.Add(SeedRange{start: seed.start + baseReplacement - lowerLimit, rangeSize: seed.rangeSize})
					}
				}
			}
		}
	}

	result := 999999999999
	for seed := range seeds.Iter() {
		if int(seed.start) < result {
			result = int(seed.start)
		}
	}
	return uint32(result)
}

func StarOne(input string) uint32 {
	lines := strings.Split(input, "\n")
	seedRange := mapset.NewSet[SeedRange]()
	for info := range extractSeedsInfo(lines[0]).Iter() {
		seedRange.Add(SeedRange{start: info, rangeSize: 1})
	}
	return seedToLocation(lines, seedRange)
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
