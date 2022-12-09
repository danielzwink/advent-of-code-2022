package main

import (
	"advent-of-code-2022/cmd/puzzles/09/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
)

func main() {
	steps := readSingleSteps("09/input")

	fmt.Printf("Part 1: %v\n", part1(steps))
	fmt.Printf("Part 2: %v\n", part2(steps))
}

func part1(steps []*types.Step) int {
	head := types.NewPosition(0, 0)
	tail := types.NewPosition(0, 0)
	tailVisits := make(map[string]bool, 0)

	for _, step := range steps {
		head.Move(step)
		head.Tighten(tail)
		tailVisits[tail.Key()] = true
	}
	return len(tailVisits)
}

func part2(steps []*types.Step) int {
	head := types.NewPosition(0, 0)
	k1 := types.NewPosition(0, 0)
	k2 := types.NewPosition(0, 0)
	k3 := types.NewPosition(0, 0)
	k4 := types.NewPosition(0, 0)
	k5 := types.NewPosition(0, 0)
	k6 := types.NewPosition(0, 0)
	k7 := types.NewPosition(0, 0)
	k8 := types.NewPosition(0, 0)
	tail := types.NewPosition(0, 0)
	tailVisits := make(map[string]bool, 0)

	for _, step := range steps {
		head.Move(step)
		head.Tighten(k1)
		k1.Tighten(k2)
		k2.Tighten(k3)
		k3.Tighten(k4)
		k4.Tighten(k5)
		k5.Tighten(k6)
		k6.Tighten(k7)
		k7.Tighten(k8)
		k8.Tighten(tail)
		tailVisits[tail.Key()] = true
	}
	return len(tailVisits)
}

func readSingleSteps(day string) []*types.Step {
	lines := util.ReadFile(day)

	steps := make([]*types.Step, 0)
	for _, line := range lines {
		direction := rune(line[0])
		count := util.MustParseInt(line[2:])

		for i := 0; i < count; i++ {
			steps = append(steps, types.NewStep(direction))
		}
	}
	return steps
}
