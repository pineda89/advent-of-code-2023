package day01

import (
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day01"
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
	var whitelist = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var mapWhitelist = make(map[string]int)
	for i := range whitelist {
		mapWhitelist[whitelist[i]] = i
	}

	return strconv.Itoa(d.calibration(filepath, whitelist, mapWhitelist))
}

func (d *Day) Part2(filepath string) string {
	var whitelist = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var mapWhitelist = make(map[string]int)
	for i := range whitelist {
		if i >= 10 {
			mapWhitelist[whitelist[i]] = i - 10
		} else {
			mapWhitelist[whitelist[i]] = i
		}
	}

	return strconv.Itoa(d.calibration(filepath, whitelist, mapWhitelist))
}

func (d *Day) calibration(filepath string, whitelist []string, whitelistValues map[string]int) (total int) {
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		numbers := make([]int, 0)
		for i := range line {
			for j := range whitelist {
				if strings.HasPrefix(line[i:], whitelist[j]) {
					numbers = append(numbers, whitelistValues[whitelist[j]])
				}
			}
		}
		total += (numbers[0] * 10) + numbers[len(numbers)-1]
	}
	return total
}
