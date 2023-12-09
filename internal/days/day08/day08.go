package day08

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day struct {
	nodes map[string]*node
}

func (d *Day) GetDay() string {
	return "day08"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

type node struct {
	tag   string
	left  *node
	right *node
}

var re = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func (d *Day) get(tag string) *node {
	if v, ok := d.nodes[tag]; ok {
		return v
	} else {
		v = &node{tag: tag}
		d.nodes[tag] = v
		return v
	}
}

func (d *Day) Part1(filepath string) string {
	d.nodes = make(map[string]*node)
	lines := strings.Split(d.GetInput(filepath), "\n")

	data := lines[0]
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		matches := re.FindStringSubmatch(line)
		n := d.get(matches[1])
		n.left = d.get(matches[2])
		n.right = d.get(matches[3])
	}

	return strconv.Itoa(d.calculatePath(d.get("AAA"), data, "ZZZ"))
}

func (d *Day) Part2(filepath string) string {
	d.nodes = make(map[string]*node)
	lines := strings.Split(d.GetInput(filepath), "\n")

	sources := make([]*node, 0)
	data := lines[0]
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		matches := re.FindStringSubmatch(line)
		n := d.get(matches[1])
		n.left = d.get(matches[2])
		n.right = d.get(matches[3])
		if strings.HasSuffix(n.tag, "A") {
			sources = append(sources, n)
		}
	}

	results := make([]int, len(sources))
	for i := range sources {
		results[i] = d.calculatePath(sources[i], data, "Z")
	}

	return strconv.Itoa(lcm(results...))
}

func (d *Day) calculatePath(current *node, data string, end string) int {
	var total = 0
	for !strings.HasSuffix(current.tag, end) {
		instruction := string(data[total%len(data)])
		switch instruction {
		case "L":
			current = current.left
		case "R":
			current = current.right
		}
		total++
	}
	return total
}

func lcm(values ...int) int {
	result := values[0] * values[1] / gcd(values[0], values[1])

	for i := 2; i < len(values); i++ {
		result = lcm(result, values[i])
	}

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
