package main

import (
	"advent-of-code-2022/cmd/puzzles/13/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	pairs := readPairs("13/input")

	fmt.Printf("Part 1: %v\n", part1(pairs))
	fmt.Printf("Part 2: %v\n", part2(pairs))
}

func part1(pairs []*types.Pair) int {
	sum := 0
	for i, pair := range pairs {
		if types.Compare(pair.LeftPacket, pair.RightPacket) == -1 {
			sum += i + 1
		}
	}
	return sum
}

func part2(pairs []*types.Pair) int {
	packets := make([]*types.Element, 0, len(pairs)*2)
	for _, pair := range pairs {
		packets = append(packets, pair.LeftPacket)
		packets = append(packets, pair.RightPacket)
	}
	sep1, _ := parseList("[[2]]")
	sep2, _ := parseList("[[6]]")
	packets = append(packets, sep1)
	packets = append(packets, sep2)

	sort.Stable(types.SortedElements(packets))

	i1, i2 := 0, 0
	for i := 0; i < len(packets); i++ {
		if sep1 == packets[i] {
			i1 = i + 1
		}
		if sep2 == packets[i] {
			i2 = i + 1
		}
	}

	return i1 * i2
}

func readPairs(day string) []*types.Pair {
	lines := util.ReadFile(day)

	pairs := make([]*types.Pair, 0)
	for i := 0; i < len(lines); i += 3 {
		left, _ := parseList(lines[i])
		right, _ := parseList(lines[i+1])
		pairs = append(pairs, types.NewPair(left, right))
	}
	return pairs
}

func parseList(line string) (*types.Element, int) {
	result := types.ListElement()

	for offset := 0; offset < len(line); {
		switch line[offset] {
		case '[':
			offset++
			list, i := parseList(line[offset:])
			result.AddToListElement(list)
			offset += i
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number, i := parseNumber(line[offset:])
			result.AddToListElement(number)
			offset += i
		case ',':
			offset++
		case ']':
			offset++
			return result, offset
		}
	}

	return result, len(line)
}

func parseNumber(line string) (*types.Element, int) {
	number := strings.Builder{}
	offset := 0
	for _, r := range line {
		if unicode.IsNumber(r) {
			number.WriteRune(r)
			offset++
		} else {
			break
		}
	}
	return types.NumberElement(number.String()), offset
}
