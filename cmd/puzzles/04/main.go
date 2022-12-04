package main

import (
	"advent-of-code-2022/cmd/puzzles/04/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	pairs := readPairs("04/input")

	sum := 0
	for _, pair := range pairs {
		if pair.FullyContained() {
			sum += 1
		}
	}
	return sum
}

func part2() int {
	pairs := readPairs("04/input")

	sum := 0
	for _, pair := range pairs {
		if pair.Overlapped() {
			sum += 1
		}
	}
	return sum
}

func readPairs(day string) []*types.Pair {
	lines := util.ReadFile(day)

	rucksacks := make([]*types.Pair, 0)
	for _, line := range lines {
		rucksacks = append(rucksacks, types.NewPair(line))
	}
	return rucksacks
}
