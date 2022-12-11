package main

import (
	"advent-of-code-2022/cmd/puzzles/11/types"
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	monkeys := getInputMonkeys()

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, forward := range monkey.Turn(0) {
				monkeys[forward.Monkey].AddItem(forward.Item)
			}
		}
	}
	return getLevelOfMonkeyBusiness(monkeys)
}

func part2() int {
	monkeys := getInputMonkeys()

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for _, forward := range monkey.Turn(9699690) {
				monkeys[forward.Monkey].AddItem(forward.Item)
			}
		}
	}
	return getLevelOfMonkeyBusiness(monkeys)
}

func getLevelOfMonkeyBusiness(monkeys []*types.Monkey) int {
	inspections := make([]int, 0, len(monkeys))
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}

func getInputMonkeys() []*types.Monkey {
	monkeys := make([]*types.Monkey, 0)

	m0 := types.NewMonkey([]types.WorryLevel{56, 56, 92, 65, 71, 61, 79}, func(value types.WorryLevel) types.WorryLevel {
		return value * 7
	}, 3, 3, 7)
	m1 := types.NewMonkey([]types.WorryLevel{61, 85}, func(value types.WorryLevel) types.WorryLevel {
		return value + 5
	}, 11, 6, 4)
	m2 := types.NewMonkey([]types.WorryLevel{54, 96, 82, 78, 69}, func(value types.WorryLevel) types.WorryLevel {
		return value * value
	}, 7, 0, 7)
	m3 := types.NewMonkey([]types.WorryLevel{57, 59, 65, 95}, func(value types.WorryLevel) types.WorryLevel {
		return value + 4
	}, 2, 5, 1)
	m4 := types.NewMonkey([]types.WorryLevel{62, 67, 80}, func(value types.WorryLevel) types.WorryLevel {
		return value * 17
	}, 19, 2, 6)
	m5 := types.NewMonkey([]types.WorryLevel{91}, func(value types.WorryLevel) types.WorryLevel {
		return value + 7
	}, 5, 1, 4)
	m6 := types.NewMonkey([]types.WorryLevel{79, 83, 64, 52, 77, 56, 63, 92}, func(value types.WorryLevel) types.WorryLevel {
		return value + 6
	}, 17, 2, 0)
	m7 := types.NewMonkey([]types.WorryLevel{50, 97, 76, 96, 80, 56}, func(value types.WorryLevel) types.WorryLevel {
		return value + 3
	}, 13, 3, 5)

	monkeys = append(monkeys, m0, m1, m2, m3, m4, m5, m6, m7)
	return monkeys
}
