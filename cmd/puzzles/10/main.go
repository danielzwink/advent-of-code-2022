package main

import (
	"advent-of-code-2022/cmd/puzzles/10/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"strings"
)

func main() {
	cycles := readCycles("10/input")

	fmt.Printf("Part 1: %v\n", part1(cycles))
	fmt.Printf("Part 2: %v\n", part2(cycles))
}

func part1(cycles []*types.Cycle) int {
	signalStrengthSum := 0

	registerX := 1
	for i := 1; i <= 220; i++ {
		cycle := cycles[i-1]

		if i == 20 {
			signalStrengthSum += registerX * 20
		} else if i == 60 {
			signalStrengthSum += registerX * 60
		} else if i == 100 {
			signalStrengthSum += registerX * 100
		} else if i == 140 {
			signalStrengthSum += registerX * 140
		} else if i == 180 {
			signalStrengthSum += registerX * 180
		} else if i == 220 {
			signalStrengthSum += registerX * 220
		}

		registerX += cycle.Summand
	}

	return signalStrengthSum
}

func part2(cycles []*types.Cycle) int {
	var crt [6][40]rune

	registerX := 1
	for row := 0; row <= 5; row++ {
		for col := 0; col <= 39; col++ {
			cycle := row*40 + col

			if col >= registerX-1 && col <= registerX+1 {
				// sprite overlaps
				crt[row][col] = '#'
			} else {
				crt[row][col] = ' '
			}

			registerX += cycles[cycle].Summand
		}
	}

	printCrt(crt)
	return 0
}

func printCrt(crt [6][40]rune) {
	for row := 0; row <= 5; row++ {
		var line string
		for col := 0; col <= 39; col++ {
			line += fmt.Sprintf("%c ", crt[row][col])
		}
		println(line)
	}
}

func readCycles(day string) []*types.Cycle {
	lines := util.ReadFile(day)

	cycles := make([]*types.Cycle, 0)
	for _, line := range lines {
		if line == "noop" {
			cycles = append(cycles, types.EmptyCycle())
		} else {
			split := strings.Split(line, " ")
			value := util.MustParseInt(split[1])

			cycles = append(cycles, types.EmptyCycle())
			cycles = append(cycles, types.NewCycle(value))
		}
	}
	return cycles
}
