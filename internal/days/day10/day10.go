package day10

import (
	"os"
	"strconv"
	"strings"
)

type Day struct {
	scores map[xy]int
}

func (d *Day) GetDay() string {
	return "day10"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

type grid struct {
	width  int
	height int
	data   map[xy]string
}

type xy struct {
	x int
	y int
}

func (d *Day) Part1(filepath string) string {
	lines := strings.Split(d.GetInput(filepath), "\n")
	initialPosition := xy{}
	grid := grid{
		data:   map[xy]string{},
		height: len(lines),
		width:  len(lines[0]),
	}

	for i, line := range lines {
		for j := range line {
			grid.data[xy{x: j, y: i}] = string(line[j])
			if string(line[j]) == "S" {
				initialPosition = xy{x: j, y: i}
			}
		}
	}

	_, maxDist := d.doPart1([]xy{initialPosition}, grid)

	return strconv.Itoa(maxDist)
}

func (d *Day) Part2(filepath string) string {
	lines := strings.Split(d.GetInput(filepath), "\n")
	initialPosition := xy{}
	grid := grid{
		data:   map[xy]string{},
		height: len(lines),
		width:  len(lines[0]),
	}

	for i, line := range lines {
		for j := range line {
			grid.data[xy{x: j, y: i}] = string(line[j])
			if string(line[j]) == "S" {
				initialPosition = xy{x: j, y: i}
			}
		}
	}

	visited, _ := d.doPart1([]xy{initialPosition}, grid)

	total := 0
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			if !d.isVisited(xy{x: x, y: y}, visited) && d.isInside(grid, xy{x: x, y: y}, visited) {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) isInside(data grid, p xy, visited map[xy]int) bool {
	count := 0
	cornerCounts := make(map[string]int)
	for y := p.y; y < data.height; y++ {
		tile := data.data[xy{y: y, x: p.x}]
		if tile == "S" {
			tile = d.findStartTile(xy{x: p.x, y: y}, data)
		}
		if _, ok := visited[xy{x: p.x, y: y}]; ok {
			if tile == "-" {
				count++
			} else if tile != "|" && tile != "." {
				cornerCounts[tile]++
			}
		}
	}

	count += max(cornerCounts["L"], cornerCounts["7"]) - d.abs(cornerCounts["L"]-cornerCounts["7"])
	count += max(cornerCounts["F"], cornerCounts["J"]) - d.abs(cornerCounts["F"]-cornerCounts["J"])
	return count%2 == 1
}

func (d *Day) findStartTile(start xy, data grid) string {

	points := d.findPoints(data, start)
	xymin := xy{
		x: min(points[0].x, points[1].x),
		y: min(points[0].y, points[1].y),
	}

	xymax := xy{
		x: max(points[0].x, points[1].x),
		y: max(points[0].y, points[1].y),
	}

	if points[0].x == points[1].x {
		return "|"
	} else if points[0].y == points[1].y {
		return "-"
	} else if xymin.x < start.x && xymin.y < start.y {
		return "J"
	} else if xymax.x > start.x && xymax.y > start.y {
		return "F"
	} else if xymax.x > start.x && xymin.y < start.y {
		return "L"
	} else if xymin.x < start.x && xymax.y > start.y {
		return "7"
	}
	return "."
}

func (d *Day) doPart1(uncheckedPoints []xy, datagrid grid) (map[xy]int, int) {
	visited := make(map[xy]int)

	var maxDistance int
	for len(uncheckedPoints) != 0 {
		first := uncheckedPoints[0]
		for _, point := range d.findPoints(datagrid, first) {
			if _, found := visited[point]; !found {
				visited[point] = visited[first] + 1
				uncheckedPoints = append(uncheckedPoints, point)
				maxDistance = max(maxDistance, visited[point])
			}
		}
		uncheckedPoints = uncheckedPoints[1:]
	}
	return visited, maxDistance
}

func (d *Day) isVisited(p xy, visited map[xy]int) bool {
	_, ok := visited[p]
	return ok
}

func (d *Day) findPoints(grid grid, p xy) []xy {
	var points []xy
	switch grid.data[p] {
	case "|":
		points = append(points, xy{x: p.x, y: 1 + p.y}, xy{x: p.x, y: p.y - 1})
	case "-":
		points = append(points, xy{x: p.x + 1, y: p.y}, xy{x: p.x - 1, y: p.y})
	case "L":
		points = append(points, xy{x: p.x, y: p.y - 1}, xy{x: p.x + 1, y: p.y})
	case "J":
		points = append(points, xy{x: p.x, y: p.y - 1}, xy{x: p.x - 1, y: p.y})
	case "7":
		points = append(points, xy{x: p.x, y: p.y + 1}, xy{x: p.x - 1, y: p.y})
	case "F":
		points = append(points, xy{x: p.x, y: p.y + 1}, xy{x: p.x + 1, y: p.y})
	default:
		if d.oneOf(grid.data[xy{x: p.x, y: p.y + 1}], "|", "L", "J") {
			points = append(points, xy{x: p.x, y: p.y + 1})
		}
		if d.oneOf(grid.data[xy{x: p.x + 1, y: p.y}], "-", "7", "J") {
			points = append(points, xy{x: p.x + 1, y: p.y})
		}
		if d.oneOf(grid.data[xy{x: p.x + 1, y: p.y}], "|", "7", "F") {
			points = append(points, xy{x: p.x, y: p.y - 1})
		}
		if d.oneOf(grid.data[xy{x: p.x + 1, y: p.y}], "-", "L", "F") {
			points = append(points, xy{x: p.x - 1, y: p.y})
		}
	}
	return points
}

func (d *Day) oneOf(s string, s2 ...string) bool {
	for i := range s2 {
		if s == s2[i] {
			return true
		}
	}
	return false
}

func (d *Day) abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
