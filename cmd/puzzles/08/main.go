package main

import (
	"advent-of-code-2022/cmd/puzzles/08/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
)

func main() {
	trees := readTreeMap("08/input")
	trees.CalculateVisibilitiesAndScenicScores()

	fmt.Printf("Part 1: %v\n", part1(trees))
	fmt.Printf("Part 2: %v\n", part2(trees))
}

func part1(trees types.TreeMap) int {
	return trees.VisibleTreeCount()
}

func part2(trees types.TreeMap) int {
	return trees.HighestScenicScore()
}

func readTreeMap(day string) types.TreeMap {
	lines := util.ReadFile(day)

	trees := make(types.TreeMap, 0)
	for _, line := range lines {
		row := make([]*types.Tree, 0)
		for _, c := range line {
			height := util.MustParseInt(string(c))
			row = append(row, types.NewTree(height))
		}
		trees = append(trees, row)
	}
	return trees
}
