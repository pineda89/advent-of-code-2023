package day02

import (
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
			splittedCubeData := strings.Split(cubeData, " ")
			parsedValue, _ := strconv.Atoi(splittedCubeData[0])
			if oldValue := maxValues[splittedCubeData[1]]; parsedValue > oldValue {
				maxValues[splittedCubeData[1]] = parsedValue
			}
		}
	}
	return maxValues
}
