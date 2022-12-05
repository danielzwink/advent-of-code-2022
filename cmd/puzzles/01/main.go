package main

import (
	"advent-of-code-2022/pkg/util"
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	calories := getCaloriesPerElf("01/input")

	sort.Ints(calories)
	length := len(calories)

	return calories[length-1]
}

func part2() int {
	calories := getCaloriesPerElf("01/input")

	sort.Ints(calories)
	length := len(calories)

	return calories[length-1] + calories[length-2] + calories[length-3]
}

func getCaloriesPerElf(day string) []int {
	lines := util.ReadFile(day)
	calories := make([]int, 0)

	sum := 0
	for _, line := range lines {
		number, valid := util.ParseInt(line)

		if valid {
			sum += number
		} else {
			calories = append(calories, sum)
			sum = 0
		}
	}
	return calories
}
