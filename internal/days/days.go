package days

import (
	"advent-of-code-2023/internal/days/day01"
	"advent-of-code-2023/internal/days/day02"
	"advent-of-code-2023/internal/days/day03"
	"advent-of-code-2023/internal/days/day04"
	"advent-of-code-2023/internal/days/day05"
)

var DaysArray []Day

func init() {
	addDay(&day01.Day{})
	addDay(&day02.Day{})
	addDay(&day03.Day{})
	addDay(&day04.Day{})
	addDay(&day05.Day{})
}

type Day interface {
	GetDay() string
	GetInput(filepath string) string
	GetReadme(filepath string) string
	Part1(filepath string) string
	Part2(filepath string) string
}

func addDay(day Day) {
	DaysArray = append(DaysArray, day)
}
