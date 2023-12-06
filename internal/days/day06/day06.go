package day06

import (
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day06"
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
	var total = 1
	data := make(map[string][]int)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		fields := strings.Fields(line)
		for j := 1; j < len(fields); j++ {
			v, _ := strconv.Atoi(fields[j])
			data[fields[0][:len(fields[0])-1]] = append(data[fields[0][:len(fields[0])-1]], v)
		}
	}

	for i := 0; i < len(data["Time"]); i++ {
		t := data["Time"][i]
		dist := data["Distance"][i]

		var numWaysToBeatRecord = 0

		for timePress := 1; timePress < t; timePress++ {
			distance := (t - timePress) * timePress
			if distance > dist {
				numWaysToBeatRecord++
			}
		}
		total *= numWaysToBeatRecord
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total int
	data := make(map[string]int)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		fields := strings.Fields(line)
		data[fields[0][:len(fields[0])-1]], _ = strconv.Atoi(strings.Join(fields[1:], ""))
	}

	t := data["Time"]
	dist := data["Distance"]

	for timePress := 1; timePress < t; timePress++ {
		distance := (t - timePress) * timePress
		if distance > dist {
			total++
		}
	}

	return strconv.Itoa(total)
}
