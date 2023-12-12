package day12

import (
	"os"
	"strconv"
	"strings"
)

const (
	state_operational = "."
	state_damaged     = "#"
	state_unknown     = "?"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day12"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) Part1(filepath string) string {
	var total = 0
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		fields := strings.Fields(line)
		md := data{
			cache: make(map[[2]int]int),
		}
		for _, v := range fields[0] {
			md.springs = append(md.springs, string(v))
		}
		for _, v := range strings.Split(fields[1], ",") {
			t, _ := strconv.Atoi(v)
			md.groups = append(md.groups, t)
		}
		total += md.solve(0, 0)
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total = 0
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		fields := strings.Fields(line)
		md := data{
			cache: make(map[[2]int]int),
		}
		for _, v := range fields[0] {
			md.springs = append(md.springs, string(v))
		}
		for _, v := range strings.Split(fields[1], ",") {
			t, _ := strconv.Atoi(v)
			md.groups = append(md.groups, t)
		}

		originalStr := md.springs
		for i := 0; i < len(originalStr)*(5-1); i++ {
			if i%len(originalStr) == 0 {
				md.springs = append(md.springs, state_unknown)
			}
			md.springs = append(md.springs, originalStr[i%len(originalStr)])
		}

		original := md.groups
		for i := 0; i < len(original)*(5-1); i++ {
			md.groups = append(md.groups, original[i%len(original)])
		}

		total += md.solve(0, 0)
	}

	return strconv.Itoa(total)
}

type data struct {
	springs []string
	groups  []int
	cache   map[[2]int]int
}

func (d data) solve(springsIndex int, groupIndex int) int {
	if springsIndex >= len(d.springs) {
		if groupIndex < len(d.groups) {
			return 0
		}
		return 1
	}

	if v, ok := d.cache[[2]int{springsIndex, groupIndex}]; ok {
		return v
	}

	validArrangements := 0
	switch d.springs[springsIndex] {
	case state_operational:
		validArrangements = d.solve(springsIndex+1, groupIndex)
	case state_unknown:
		validArrangements += d.solve(springsIndex+1, groupIndex)
	}

	if groupIndex < len(d.groups) {
		count := 0
		for k := springsIndex; k < len(d.springs); k++ {
			if d.springs[k] == state_operational || (count == d.groups[groupIndex] && d.springs[k] == state_unknown) {
				break
			}
			count++
		}

		if count == d.groups[groupIndex] {
			validArrangements += d.solve(springsIndex+count+1, groupIndex+1)
		}
	}

	d.cache[[2]int{springsIndex, groupIndex}] = validArrangements
	return validArrangements
}
