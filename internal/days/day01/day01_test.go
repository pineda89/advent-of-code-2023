package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1RealData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input.txt")

	assert.Equal(t, "54338", result)
}

func TestPart1TestData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input_testdata_part1.txt")

	assert.Equal(t, "142", result)
}

func TestPart2RealData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input.txt")

	assert.Equal(t, "53389", result)
}

func TestPart2TestData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input_testdata_part2.txt")

	assert.Equal(t, "281", result)
}
