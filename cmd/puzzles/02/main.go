package main

import (
	"advent-of-code-2022/cmd/puzzles/02/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	rounds := readRounds("02/input")

	totalScore := 0
	for _, round := range rounds {
		round.EvaluateOutcome()
		totalScore += round.Score()
	}
	return totalScore
}

func part2() int {
	rounds := readRounds("02/input")

	totalScore := 0
	for _, round := range rounds {
		round.EvaluateSelf()
		totalScore += round.Score()
	}
	return totalScore
}

func readRounds(day string) []*types.Round {
	lines := util.ReadFile(day)

	rounds := make([]*types.Round, 0)
	for _, line := range lines {
		pair := strings.Split(line, " ")

		m := &types.Round{
			Opponent: valueToMove(pair[0]),
			Self:     valueToMove(pair[1]),
			Outcome:  valueToOutcome(pair[1]),
		}
		rounds = append(rounds, m)
	}
	return rounds
}

func valueToMove(value string) types.Move {
	switch value {
	case "A", "X":
		return types.Rock
	case "B", "Y":
		return types.Paper
	case "C", "Z":
		return types.Scissors
	}
	panic(1)
}

func valueToOutcome(value string) types.Result {
	switch value {
	case "X":
		return types.Lost
	case "Y":
		return types.Draw
	case "Z":
		return types.Win
	}
	panic(1)
}
