package types

import (
	"advent-of-code-2022/pkg/util"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{X: x, Y: y}
}

func NewCoordinateFromString(value string) *Coordinate {
	split := strings.Split(value, ",")
	return NewCoordinate(util.MustParseInt(split[0]), util.MustParseInt(split[1]))
}

func (c *Coordinate) fallOptions() []*Coordinate {
	down := NewCoordinate(c.X, c.Y+1)
	diagonalLeft := NewCoordinate(c.X-1, c.Y+1)
	diagonalRight := NewCoordinate(c.X+1, c.Y+1)
	return []*Coordinate{down, diagonalLeft, diagonalRight}
}

type Cave struct {
	Rocks  [][]bool
	Entry  *Coordinate
	Offset int
}

func NormalisedCave(rocks []*Coordinate, xMin, xMax, yMax int, endless bool) *Cave {
	const xSandEntry = 500
	const yFloorIncrease = 2

	if !endless {
		// increase min and max values
		yMax += yFloorIncrease
		xMin = xSandEntry - yMax
		xMax = xSandEntry + yMax
		// draw bottom line
		for x := xMin; x <= xMax; x++ {
			rocks = append(rocks, NewCoordinate(x, yMax))
		}
	}

	// normalize rock matrix (to 0,0)
	c := &Cave{
		Rocks:  make([][]bool, yMax+1),
		Entry:  NewCoordinate(xSandEntry-xMin, 0),
		Offset: xMin,
	}

	for y := 0; y <= yMax; y++ {
		c.Rocks[y] = make([]bool, xMax-xMin+1)
	}
	for _, r := range rocks {
		c.Rocks[r.Y][r.X-xMin] = true
	}

	return c
}

func (c *Cave) Print() {
	for y, row := range c.Rocks {
		for x, rock := range row {
			if rock {
				print("# ")
			} else if c.Entry.X == x && c.Entry.Y == y {
				print("+ ")
			} else {
				print(". ")
			}
		}
		println("")
	}
	println("")
}

func (c *Cave) Pour() bool {
	current := c.Entry
	open := !c.Rocks[current.Y][current.X]

	for open {
		validAndAvailable := false
		for _, target := range current.fallOptions() {
			if !c.valid(target) {
				return false
			}
			if c.available(target) {
				current = target
				validAndAvailable = true
				break
			}
		}
		if validAndAvailable {
			continue
		}

		c.Rocks[current.Y][current.X] = true
		return true
	}
	return false
}

func (c *Cave) valid(target *Coordinate) bool {
	return target.X >= 0 && target.X < len(c.Rocks[0]) && target.Y >= 0 && target.Y < len(c.Rocks)
}

func (c *Cave) available(target *Coordinate) bool {
	return !c.Rocks[target.Y][target.X]
}
