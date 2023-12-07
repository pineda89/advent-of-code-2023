package day07

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day07"
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

type hand struct {
	line        string
	handType    string
	handTypeP2  string
	cards       [5]string
	scores      [5]int
	bid         int
	numOfJokers int
}

var handTypes = []string{"Five of a kind", "Four of a kind", "Full house", "Three of a kind", "Two pair", "One pair", "High card"}
var scoreConversion = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

func (h *hand) parse(line string) {
	seen := make(map[string]int)
	fields := strings.Fields(line)
	for i := range fields[0] {
		h.cards[i] = string(line[i])
		seen[string(line[i])]++
		h.scores[i] = scoreConversion[h.cards[i]]
	}
	h.bid, _ = strconv.Atoi(fields[1])

	repeatedElements := make([]int, 0)
	for _, v := range seen {
		repeatedElements = append(repeatedElements, v)
	}
	sort.Ints(repeatedElements)

	h.evaluateHandType(seen, repeatedElements)
}

func (h *hand) parsep2(line string) {
	seen := make(map[string]int)
	fields := strings.Fields(line)
	for i := range fields[0] {
		h.cards[i] = string(line[i])
		h.scores[i] = scoreConversion[h.cards[i]]
		if h.cards[i] != "J" {
			seen[string(line[i])]++
		} else {
			h.scores[i] = 0
		}
	}

	h.bid, _ = strconv.Atoi(fields[1])
	h.numOfJokers = strings.Count(fields[0], "J")

	repeatedElements := make([]int, 0)
	for _, v := range seen {
		repeatedElements = append(repeatedElements, v)
	}
	sort.Ints(repeatedElements)

	if len(repeatedElements) == 0 {
		// special case JJJJJ
		repeatedElements = append(repeatedElements, 0)
		seen["J"] = 5
	}

	if h.numOfJokers > 0 {
		repeatedElements[len(repeatedElements)-1] += h.numOfJokers
		for i := range repeatedElements {
			if repeatedElements[i] == h.numOfJokers {
				repeatedElements = append(repeatedElements[:i], repeatedElements[i:]...)
				break
			}
		}
	}

	h.evaluateHandType(seen, repeatedElements)
}

func (h *hand) evaluateHandType(seen map[string]int, repeatedElements []int) {
	switch len(seen) {
	case 1:
		h.handType = handTypes[0]
	case 2:
		if repeatedElements[1] == 4 {
			h.handType = handTypes[1]
		} else {
			h.handType = handTypes[2]
		}
	case 3:
		h.handType = handTypes[4]
		if repeatedElements[2] == 3 {
			h.handType = handTypes[3]
		}
	case 4:
		h.handType = handTypes[5]
	case 5:
		h.handType = handTypes[6]
	}
}

func (d *Day) Part1(filepath string) string {
	hands := make(map[string][]*hand)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		h := &hand{line: line}
		h.parse(line)
		hands[h.handType] = append(hands[h.handType], h)
	}

	return d.solve(filepath, hands)
}

func (d *Day) Part2(filepath string) string {
	hands := make(map[string][]*hand)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		h := &hand{line: line}
		h.parsep2(line)
		hands[h.handType] = append(hands[h.handType], h)
	}

	return d.solve(filepath, hands)
}

func (d *Day) solve(filepath string, hands map[string][]*hand) string {
	total := 0
	currentRank := len(strings.Split(d.GetInput(filepath), "\n"))

	for i := range handTypes {
		h := hands[handTypes[i]]
		sort.Slice(h, func(i, j int) bool {
			for tmp := range h {
				if h[i].scores[tmp] > h[j].scores[tmp] {
					return false
				} else if h[i].scores[tmp] < h[j].scores[tmp] {
					return true
				}
			}
			return false
		})

		for j := len(h) - 1; j >= 0; j-- {
			total += h[j].bid * currentRank
			currentRank--
		}
	}

	return strconv.Itoa(total)
}
