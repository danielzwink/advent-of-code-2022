package main

import (
	"advent-of-code-2022/cmd/puzzles/05/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"regexp"
)

func main() {
	fmt.Printf("Part 1: %s\n", part1())
	fmt.Printf("Part 2: %s\n", part2())
}

func part1() string {
	configuration := readConfiguration("05/input")
	configuration.Move(false)
	return configuration.Result()
}

func part2() string {
	configuration := readConfiguration("05/input")
	configuration.Move(true)
	return configuration.Result()
}

func readConfiguration(day string) *types.Configuration {
	lines := util.ReadFile(day)
	separatorIndex := findSeparatorIndex(lines)

	return &types.Configuration{
		Stacks:    parseInitialStacks(lines[:separatorIndex-1]),
		Movements: parseMovements(lines[separatorIndex+1:]),
	}
}

func findSeparatorIndex(lines []string) int {
	for i, line := range lines {
		if line == "" {
			return i
		}
	}
	panic(1)
}

func parseInitialStacks(lines []string) map[int][]rune {
	stacks := make(map[int][]rune, 0)

	// init map
	size := (len(lines[0]) + 1) / 4
	for i := 1; i <= size; i++ {
		stacks[i] = make([]rune, 0)
	}

	// parse stacks in reverse order
	for i := len(lines) - 1; i >= 0; i-- {
		line := []rune(lines[i])

		position := 0
		for j := 0; j < len(line); j += 4 {
			char := line[j+1]
			position++

			if char != 32 {
				stacks[position] = append(stacks[position], char)
			}
		}
	}
	return stacks
}

func parseMovements(lines []string) []*types.Movement {
	movementParser := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")

	movements := make([]*types.Movement, 0)
	for _, line := range lines {
		groups := movementParser.FindStringSubmatch(line)

		movement := &types.Movement{
			Count:       util.MustParseInt(groups[1]),
			Source:      util.MustParseInt(groups[2]),
			Destination: util.MustParseInt(groups[3]),
		}
		movements = append(movements, movement)
	}
	return movements
}
