package day09

import (
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day09"
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
		data := make([][]int, 1)
		for _, field := range strings.Fields(line) {
			v, _ := strconv.Atoi(field)
			data[0] = append(data[0], v)
		}

		total += d.predict(d.fill(data))
	}
	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total = 0
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		data := make([][]int, 1)
		for _, field := range strings.Fields(line) {
			v, _ := strconv.Atoi(field)
			data[0] = append(data[0], v)
		}

		data[0] = d.reverse(data[0])

		total += d.predict(d.fill(data))
	}
	return strconv.Itoa(total)
}

func (d *Day) fullZeros(ints []int) bool {
	for i := range ints {
		if ints[i] != 0 {
			return false
		}
	}
	return true
}

func (d *Day) reverse(data []int) []int {
	newData := make([]int, len(data))
	for i := range data {
		newData[len(data)-1-i] = data[i]
	}
	return newData
}

func (d *Day) fill(data [][]int) [][]int {
	level := 0
	for !d.fullZeros(data[level]) {
		data = append(data, make([]int, len(data[level])-1))
		for i := 1; i < len(data[level]); i++ {
			data[len(data)-1][i-1] = data[level][i] - data[level][i-1]
		}
		level++
	}
	return data
}

func (d *Day) predict(data [][]int) int {
	for i := len(data) - 2; i >= 0; i-- {
		data[i] = append(data[i], data[i][len(data[i])-1]+data[i+1][len(data[i+1])-1])
	}
	return data[0][len(data[0])-1]
}
