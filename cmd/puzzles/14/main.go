package main

import (
	"advent-of-code-2022/cmd/puzzles/14/types"
	"advent-of-code-2022/pkg/util"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	cave := readCave("14/input", true)

	sand := 0
	for cave.Pour() {
		sand++
	}
	return sand
}

func part2() int {
	cave := readCave("14/input", false)

	sand := 0
	for cave.Pour() {
		sand++
	}
	return sand
}

func readCave(day string, endless bool) *types.Cave {
	lines := util.ReadFile(day)

	xMin, yMin := math.MaxInt, math.MaxInt
	xMax, yMax := math.MinInt, math.MinInt

	// read rock coordinates, calculate min/max
	rocks := make([]*types.Coordinate, 0)
	for _, line := range lines {
		pointValues := strings.Split(line, " -> ")
		for i := 0; i < len(pointValues)-1; i++ {
			from := types.NewCoordinateFromString(pointValues[i])
			to := types.NewCoordinateFromString(pointValues[i+1])
			points := singlePointsFromLine(from, to)
			rocks = append(rocks, points...)

			xMin, xMax = util.EvaluateMinAndMax(xMin, xMax, from.X)
			xMin, xMax = util.EvaluateMinAndMax(xMin, xMax, to.X)
			yMin, yMax = util.EvaluateMinAndMax(yMin, yMax, from.Y)
			yMin, yMax = util.EvaluateMinAndMax(yMin, yMax, to.Y)
		}
	}

	// just for safety
	if yMin <= 0 {
		panic(1)
	}

	// return normalised cave
	return types.NormalisedCave(rocks, xMin, xMax, yMax, endless)
}

func singlePointsFromLine(from, to *types.Coordinate) []*types.Coordinate {
	length := int(math.Max(math.Abs(float64(from.X-to.X)), math.Abs(float64(from.Y-to.Y)))) + 1
	line := make([]*types.Coordinate, 0, length)
	line = append(line, from)

	if from.X == to.X && from.Y == to.Y {
		return line
	} else if from.X == to.X {
		start, end := util.Sort(from.Y, to.Y)
		for i := start; i <= end; i++ {
			line = append(line, types.NewCoordinate(from.X, i))
		}
	} else if from.Y == to.Y {
		start, end := util.Sort(from.X, to.X)
		for i := start; i <= end; i++ {
			line = append(line, types.NewCoordinate(i, from.Y))
		}
	}
	return line
}

func sort(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}
