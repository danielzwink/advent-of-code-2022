package main

import (
	"advent-of-code-2022/pkg/util"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	line := readFirstLine("06/input")
	marker := findStartOfDistinctPacketMarker([]rune(line), 4)
	return marker
}

func part2() int {
	line := readFirstLine("06/input")
	marker := findStartOfDistinctPacketMarker([]rune(line), 14)
	return marker
}

func findStartOfDistinctPacketMarker(line []rune, size int) int {
	for i := size - 1; i < len(line); i++ {
		m := make(map[rune]bool)

		for k := 0; k < size; k++ {
			m[line[i-k]] = true
		}

		if len(m) == size {
			return i + 1
		}
	}
	panic(1)
}

func readFirstLine(day string) string {
	lines := util.ReadFile(day)
	return lines[0]
}
