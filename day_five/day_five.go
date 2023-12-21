package day_five

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type SeedRange struct {
	start uint32
	end   uint32
}

func (seedRange *SeedRange) SetStart(start uint32) {
	seedRange.start = start
}

func (seedRange *SeedRange) SetEnd(end uint32) {
	seedRange.end = end
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

func seedToLocation(lines []string, seeds mapset.Set[SeedRange]) uint64 {
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
			fmt.Println("Seeds", seeds, "Removals", removals, "Conversions", conversions)

			temp := seeds.Difference(removals)
			temp2 := conversions.Difference(removals)
			seeds = temp.Union(temp2)
		} else {
			// process map
			fmt.Println("processing map")
			conversionMap := extractConversionMap(line)
			baseReplacement := conversionMap[0]
			higherLimit := conversionMap[1] + conversionMap[2] - 1
			lowerLimit := conversionMap[1]
			delta := baseReplacement - lowerLimit

			for seed := range seeds.Clone().Iter() {
				// fmt.Println("seeds", seeds)
				if seed.start > higherLimit {
					// above limits
					fmt.Println("above limits", seed)
					conversions.Add(seed)
				} else if seed.end < lowerLimit {
					// below limits
					fmt.Println("below limits", seed)
					conversions.Add(seed)
				} else {
					// in limits
					// fmt.Println("in limits")
					if seed.start < lowerLimit && seed.end <= higherLimit {
						// partially below limit and within range
						fmt.Println("partially below limit and within range", seed)
						// conversions.Add(SeedRange{start: seed.start, end: lowerLimit - seed.start})
						seeds.Remove(seed)
						seeds.Add(SeedRange{start: seed.start, end: lowerLimit - 1})
						conversions.Add(SeedRange{start: lowerLimit + delta, end: seed.end + delta})
						// removals.Add(seed)
					}
					if seed.start >= lowerLimit && seed.end > higherLimit {
						// partially above limit and within range
						fmt.Println("partially above limit and within range", seed)
						// conversions.Add(SeedRange{start: seed.start + baseReplacement - lowerLimit, end: higherLimit - seed.start})
						seeds.Remove(seed)
						seeds.Add(SeedRange{start: higherLimit, end: seed.end})
						conversions.Add(SeedRange{start: seed.start + delta, end: higherLimit + delta})
						// removals.Add(seed)
					}
					if seed.start >= lowerLimit && seed.end <= higherLimit {
						// within range
						fmt.Println("within range", seed)
						conversions.Add(SeedRange{start: seed.start + delta, end: seed.end + delta})
						removals.Add(seed)
					}
					if seed.start < lowerLimit && seed.end > higherLimit {
						// contains range within
						fmt.Println("contains range within", seed)
						seeds.Remove(seed)
						seeds.Add(SeedRange{start: seed.start, end: lowerLimit - 1})
						seeds.Add(SeedRange{start: higherLimit + 1, end: seed.end})
						// isto esta malll!!!!!! >
						conversions.Add(SeedRange{start: baseReplacement, end: baseReplacement + conversionMap[2] - 1})
					}
				}
			}
		}
	}

	var result uint64 = 99999999999999
	for seed := range seeds.Iter() {
		if uint64(seed.start) < result && uint64(seed.start) > 0 {
			result = uint64(seed.start)
		}
	}
	return uint64(result)
}

func StarOne(input string) uint64 {
	lines := strings.Split(input, "\n")
	seedRange := mapset.NewSet[SeedRange]()
	for info := range extractSeedsInfo(lines[0]).Iter() {
		seedRange.Add(SeedRange{start: info, end: info})
	}
	return seedToLocation(lines, seedRange)
}

func StarTwo(input string) uint64 {
	lines := strings.Split(input, "\n")
	seedsInfo := append([]uint32{}, extractSeedsInfoSlice(lines[0])...)
	seeds := mapset.NewSet[SeedRange]()

	for i := uint32(0); i < uint32(len(seedsInfo))-1; i += 2 {
		start := seedsInfo[i]
		rangeSize := seedsInfo[i+1]
		seeds.Add(SeedRange{start: start, end: start + rangeSize - 1})
	}

	return seedToLocation(lines, seeds)
}
