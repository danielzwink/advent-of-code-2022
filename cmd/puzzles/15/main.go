package main

import (
	"advent-of-code-2022/cmd/puzzles/15/types"
	t "advent-of-code-2022/pkg/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"regexp"
)

func main() {
	pairs := readSensorAndBeaconPairs("15/input")

	fmt.Printf("Part 1: %v\n", part1(pairs))
	fmt.Printf("Part 2: %v\n", part2(pairs))
}

func part1(pairs []*types.Pair) int {
	yCheck := 2000000

	uniqueX := make(map[int]bool, 0)
	for _, pair := range pairs {
		overlaps, xMin, xMax := pair.OverlapsY(yCheck)
		if overlaps {
			for i := xMin; i <= xMax; i++ {
				uniqueX[i] = true
			}
		}
	}
	for _, pair := range pairs {
		if pair.Beacon.Y == yCheck {
			delete(uniqueX, pair.Beacon.X)
		}
	}
	return len(uniqueX)
}

func part2(pairs []*types.Pair) int {
	max := 4000000

	for _, pair := range pairs {
		for inc := 0; inc <= pair.Distance+1; inc++ {
			dec := pair.Distance + 1 - inc

			q := t.NewCoordinate(pair.Sensor.X+inc, pair.Sensor.Y-dec)
			if validOutlaw(pairs, max, q) {
				return tuningFrequency(q)
			}

			q = t.NewCoordinate(pair.Sensor.X-inc, pair.Sensor.Y-dec)
			if validOutlaw(pairs, max, q) {
				return tuningFrequency(q)
			}

			q = t.NewCoordinate(pair.Sensor.X+inc, pair.Sensor.Y+dec)
			if validOutlaw(pairs, max, q) {
				return tuningFrequency(q)
			}

			q = t.NewCoordinate(pair.Sensor.X-inc, pair.Sensor.Y+dec)
			if validOutlaw(pairs, max, q) {
				return tuningFrequency(q)
			}
		}
	}

	return 0
}

func validOutlaw(pairs []*types.Pair, max int, c *t.Coordinate) bool {
	if c.X < 0 || c.X > max || c.Y < 0 || c.Y > max {
		return false
	}

	for _, pair := range pairs {
		if pair.Contains(c) {
			return false
		}
	}

	return true
}

func tuningFrequency(c *t.Coordinate) int {
	fmt.Printf("%v %v\n", c.X, c.Y)

	return c.X*4000000 + c.Y
}

func readSensorAndBeaconPairs(day string) []*types.Pair {
	lines := util.ReadFile(day)

	linePattern := regexp.MustCompile("Sensor at x=(\\-?[0-9]+), y=(\\-?[0-9]+)\\: closest beacon is at x=(\\-?[0-9]+), y=(\\-?[0-9]+)")
	pairs := make([]*types.Pair, 0)
	for _, line := range lines {
		groups := linePattern.FindStringSubmatch(line)
		pair := types.NewPair(groups[1], groups[2], groups[3], groups[4])
		pairs = append(pairs, pair)
	}
	return pairs
}
