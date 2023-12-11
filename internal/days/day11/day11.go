package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day11"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

type xypair struct {
	first  xy
	second xy
}

type xy struct {
	x int
	y int
}

func (x xypair) sort() xypair {
	first := x.first
	second := x.second
	if first.x > second.x {
		first, second = second, first
	} else if first.x == second.x {
		if first.y > second.y {
			first, second = second, first
		}
	}

	return xypair{first: first, second: second}
}

func (d *Day) Part1(filepath string) string {
	data := make([][]string, 0)
	galaxies := make([]xy, 0)
	for i, line := range strings.Split(d.GetInput(filepath), "\n") {
		data = append(data, make([]string, 0))
		for j := range line {
			data[i] = append(data[i], string(line[j]))
			if string(line[j]) == "#" {
				galaxies = append(galaxies, xy{x: j, y: i})
			}
		}
	}

	expandedData := d.expand(data, galaxies, 1)

	galaxies = make([]xy, 0)
	for y := 0; y < len(expandedData); y++ {
		for x := 0; x < len(expandedData[0]); x++ {
			if v := expandedData[y][x]; v == "#" {
				galaxies = append(galaxies, xy{x: x, y: y})
			}
		}
	}

	var total int
	for k := range d.findCombinations(galaxies) {
		total += abs(k.first.x-k.second.x) + abs(k.first.y-k.second.y)
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	data := make([][]string, 0)
	galaxies := make([]xy, 0)
	for i, line := range strings.Split(d.GetInput(filepath), "\n") {
		data = append(data, make([]string, 0))
		for j := range line {
			data[i] = append(data[i], string(line[j]))
			if string(line[j]) == "#" {
				galaxies = append(galaxies, xy{x: j, y: i})
			}
		}
	}

	values := make(map[int]int)
	for i := 1; i <= 2; i++ {
		expandedData := d.expand(data, galaxies, i)

		repairedGalaxies := make([]xy, 0)
		for y := 0; y < len(expandedData); y++ {
			for x := 0; x < len(expandedData[0]); x++ {
				if v := expandedData[y][x]; v == "#" {
					repairedGalaxies = append(repairedGalaxies, xy{x: x, y: y})
				}
			}
		}

		for k := range d.findCombinations(repairedGalaxies) {
			values[i] += abs(k.first.x-k.second.x) + abs(k.first.y-k.second.y)
		}
	}

	return strconv.Itoa(values[1] + ((values[2] - values[1]) * (1000000 - 2)))
}

func (d *Day) expand(data [][]string, galaxies []xy, timesToExpand int) [][]string {
	rowsToExpand := make(map[int]bool)
	colsToExpand := make(map[int]bool)

	rowsToExpandArray := make([]int, 0)
	colsToExpandArray := make([]int, 0)

	for i := range galaxies {
		rowsToExpand[galaxies[i].y] = true
		colsToExpand[galaxies[i].x] = true
	}

	for k := range rowsToExpand {
		rowsToExpandArray = append(rowsToExpandArray, k)
	}

	for k := range colsToExpand {
		colsToExpandArray = append(colsToExpandArray, k)
	}

	newData := make([][]string, len(data)+((len(data)-len(rowsToExpand))*timesToExpand))
	for i := range newData {
		newData[i] = make([]string, len(data[0])+((len(data[0])-len(colsToExpand))*timesToExpand))
	}

	for i := range newData {
		for j := range newData[i] {
			newData[i][j] = "."
		}
	}

	for i := range galaxies {
		galaxy := galaxies[i]
		yMod := 0
		for y := 0; y < galaxy.y; y++ {
			if !rowsToExpand[y] {
				yMod += timesToExpand
			}
		}
		xMod := 0
		for x := 0; x < galaxy.x; x++ {
			if !colsToExpand[x] {
				xMod += timesToExpand
			}
		}
		newData[galaxy.y+yMod][galaxy.x+xMod] = "#"
	}

	return newData
}

func (d *Day) print(data [][]string) {
	for y := range data {
		for x := range data[y] {
			fmt.Print(data[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Day) findCombinations(galaxies []xy) map[xypair]bool {
	pairs := make(map[xypair]bool)
	for i := range galaxies {
		for j := range galaxies {
			if galaxies[i] != galaxies[j] {
				p := xypair{first: galaxies[i], second: galaxies[j]}
				pairs[p.sort()] = true
			}
		}
	}
	return pairs
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
