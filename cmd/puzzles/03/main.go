package main

import (
	"advent-of-code-2022/cmd/puzzles/03/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	rucksacks := readRucksacks("03/input")

	sum := 0
	for _, r := range rucksacks {
		sharedItem := r.SharedItem()
		sum += priority(sharedItem)
	}
	return sum
}

func part2() int {
	rucksacks := readRucksacks("03/input")

	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		badge := findBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		sum += priority(badge)
	}
	return sum
}

func priority(item rune) int {
	asciiValue := util.AsciiValue(item)

	if unicode.IsLower(item) {
		return asciiValue - 96
	} else if unicode.IsUpper(item) {
		return asciiValue - 38
	}
	panic(1)
}

func findBadge(r1, r2, r3 *types.Rucksack) rune {
	for item, _ := range r1.Items {
		_, exists2 := r2.Items[item]
		if exists2 {
			_, exists3 := r3.Items[item]
			if exists3 {
				return item
			}
		}
	}
	panic(1)
}

func readRucksacks(day string) []*types.Rucksack {
	lines := util.ReadFile(day)

	rucksacks := make([]*types.Rucksack, 0)
	for _, line := range lines {
		r := types.NewRucksack(line)
		rucksacks = append(rucksacks, r)
	}
	return rucksacks
}
