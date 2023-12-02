package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day02"
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
	var total int
	for i, line := range strings.Split(d.GetInput(filepath), "\n") {
		if maxValues := d.getMaxs(line); maxValues["red"] <= 12 && maxValues["green"] <= 13 && maxValues["blue"] <= 14 {
			total += i + 1
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total int
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		maxValues := d.getMaxs(line)
		total += maxValues["red"] * maxValues["green"] * maxValues["blue"]
	}

	return strconv.Itoa(total)
}

func (d *Day) getMaxs(line string) map[string]int {
	maxValues := make(map[string]int)
	for _, gamePhase := range strings.Split(strings.Split(line, ": ")[1], "; ") {
		for _, cubeData := range strings.Split(gamePhase, ", ") {
			var parsedValue int
			var key string
			fmt.Sscanf(cubeData, "%d %s", &parsedValue, &key)
			if oldValue := maxValues[key]; parsedValue > oldValue {
				maxValues[key] = parsedValue
			}
		}
	}
	return maxValues
}
