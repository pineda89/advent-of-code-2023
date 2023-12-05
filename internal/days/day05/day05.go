package day05

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day05"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

type rule struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func (d *Day) Part1(filepath string) string {
	var seeds = make([]int, 0)
	var current string
	var rules = make(map[string][]rule)
	var rulesTitles = make([]string, 0)

	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		switch {
		case strings.HasPrefix(line, "seeds: "):
			for i, field := range strings.Fields(line) {
				if i > 0 {
					v, _ := strconv.Atoi(field)
					seeds = append(seeds, v)
				}
			}
		case strings.HasSuffix(line, " map:"):
			current = strings.Split(line, " map:")[0]
			rulesTitles = append(rulesTitles, current)
			rules[current] = make([]rule, 0)
		case len(line) < 1:
		default:
			var destinationRangeStart, sourceRangeStart, rangeLength int
			fmt.Sscanf(line, "%d %d %d", &destinationRangeStart, &sourceRangeStart, &rangeLength)

			rules[current] = append(rules[current], rule{destinationRangeStart: destinationRangeStart, sourceRangeStart: sourceRangeStart, rangeLength: rangeLength})
		}
	}

	var total = math.MaxInt
	for i := range seeds {
		var value = seeds[i]
		for _, title := range rulesTitles {
			value = d.calculateLocationP1(value, rules[title])
		}
		if value < total {
			total = value
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var seeds = make([]int, 0)
	var intervalSeeds = make([]seed, 0)
	var current string
	var rules = make(map[string][]rule)
	var rulesTitles = make([]string, 0)

	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		switch {
		case strings.HasPrefix(line, "seeds: "):
			for i, field := range strings.Fields(line) {
				if i > 0 {
					v, _ := strconv.Atoi(field)
					seeds = append(seeds, v)
				}
			}
		case strings.HasSuffix(line, " map:"):
			current = strings.Split(line, " map:")[0]
			rulesTitles = append(rulesTitles, current)
			rules[current] = make([]rule, 0)
		case len(line) < 1:
		default:
			var destinationRangeStart, sourceRangeStart, rangeLength int
			fmt.Sscanf(line, "%d %d %d", &destinationRangeStart, &sourceRangeStart, &rangeLength)

			rules[current] = append(rules[current], rule{destinationRangeStart: destinationRangeStart, sourceRangeStart: sourceRangeStart, rangeLength: rangeLength})
		}
	}

	for i := 0; i < len(seeds); i = i + 2 {
		intervalSeeds = append(intervalSeeds, seed{position: seeds[i], size: seeds[i+1]})
	}

	i := 0
	bestCaseFound := false
	for !bestCaseFound {
		i++

		tmp := i
		for j := len(rulesTitles) - 1; j >= 0; j-- {
			tmp = d.calculateLocationP2(tmp, rules[rulesTitles[j]])
		}

		for _, newSeed := range intervalSeeds {
			if newSeed.position <= tmp && tmp < newSeed.position+newSeed.size {
				bestCaseFound = true
			}
		}
	}

	return strconv.Itoa(i)
}

func (d *Day) calculateLocationP1(seed int, data []rule) int {
	for i := range data {
		if seed >= data[i].sourceRangeStart && seed <= data[i].sourceRangeStart+data[i].rangeLength {
			return seed - data[i].sourceRangeStart + data[i].destinationRangeStart
		}
	}
	return seed
}

func (d *Day) calculateLocationP2(seed int, data []rule) int {
	for i := range data {
		if data[i].destinationRangeStart <= seed && seed < (data[i].destinationRangeStart+data[i].rangeLength) {
			return data[i].sourceRangeStart + (seed - data[i].destinationRangeStart)
		}
	}
	return seed
}

type seed struct {
	position int
	size     int
}
