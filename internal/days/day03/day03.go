package day03

import (
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day03"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

type number struct {
	fromX int
	toX   int
	y     int
	value int
}

func (d *Day) Part1(filepath string) string {
	var total int
	input := d.GetInput(filepath)
	splittedInput := strings.Split(input, "\n")

	numbers := d.getNumbers(splittedInput)
	for i := range numbers {
		if d.hasSymbolsClose(splittedInput, numbers[i].y, numbers[i].fromX, numbers[i].toX-numbers[i].fromX+1) {
			total += numbers[i].value
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total int
	input := d.GetInput(filepath)
	splittedInput := strings.Split(input, "\n")

	numbers := d.getNumbers(splittedInput)

	for i, line := range splittedInput {
		for j := 0; j < len(line); j++ {
			if line[j] == 42 {
				// is an *
				var val, count = 1, 0
				for _, num := range numbers {
					if num.y >= i-1 && num.y <= i+1 &&
						(j-1 >= num.fromX && j-1 <= num.toX || j >= num.fromX && j <= num.toX || j+1 >= num.fromX && j+1 <= num.toX) {
						val *= num.value
						count++
					}
				}

				if count > 1 {
					total += val
				}
			}
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) hasSymbolsClose(splittedInput []string, i int, j int, size int) bool {
	for myI := i - 1; myI <= i+1; myI++ {
		for myJ := j - 1; myJ <= j+size; myJ++ {
			if myI >= 0 && myI < len(splittedInput) && myJ >= 0 && myJ < len(splittedInput[j]) {
				if !(splittedInput[myI][myJ] >= 48 && splittedInput[myI][myJ] <= 57) && splittedInput[myI][myJ] != 46 {
					return true
				}
			}
		}
	}
	return false
}

func (d *Day) getNumbers(input []string) []number {
	numbers := make([]number, 0)
	for i, line := range input {
		for j := 0; j < len(line); j++ {
			if line[j] >= 48 && line[j] <= 57 {
				// is number, check for neighbors
				size := 1
				for j+size < len(line) && line[j+size] >= 48 && line[j+size] <= 57 {
					size++
				}
				val, _ := strconv.Atoi(line[j : j+size])
				numbers = append(numbers, number{value: val, fromX: j, toX: j + size - 1, y: i})
				j += size
			}
		}
	}
	return numbers
}
