package day04

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day04"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

var re = regexp.MustCompile(`Card\s+(\d+):\s+([\d\s]+)\|\s+([\d\s]+)`)

func (d *Day) Part1(filepath string) string {
	var total int

	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		matches := re.FindStringSubmatch(line)
		winningNumbers := make([]int, 0)
		yourNumbers := make([]int, 0)

		for _, winningNumber := range strings.Fields(matches[2]) {
			v, _ := strconv.Atoi(winningNumber)
			winningNumbers = append(winningNumbers, v)
		}

		for _, winningNumber := range strings.Fields(matches[3]) {
			v, _ := strconv.Atoi(winningNumber)
			yourNumbers = append(yourNumbers, v)
		}

		num := 0
		for i := range yourNumbers {
			if contains(winningNumbers, yourNumbers[i]) {
				if num == 0 {
					num = 1
				} else {
					num *= 2
				}
			}
		}
		total += num

	}
	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total int
	var numElements = make(map[int]int)

	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		matches := re.FindStringSubmatch(line)
		winningNumbers := make([]int, 0)
		yourNumbers := make([]int, 0)

		numCard, _ := strconv.Atoi(matches[1])
		numElements[numCard]++

		for _, winningNumber := range strings.Fields(matches[2]) {
			v, _ := strconv.Atoi(winningNumber)
			winningNumbers = append(winningNumbers, v)
		}

		for _, winningNumber := range strings.Fields(matches[3]) {
			v, _ := strconv.Atoi(winningNumber)
			yourNumbers = append(yourNumbers, v)
		}

		num := 0
		for i := range yourNumbers {
			if contains(winningNumbers, yourNumbers[i]) {
				num++
			}
		}

		for i := 0; i < numElements[numCard]; i++ {
			for j := 0; j < num; j++ {
				numElements[numCard+j+1]++
			}
		}

		total += numElements[numCard]
	}
	return strconv.Itoa(total)
}

func contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
