package types

import (
	t "advent-of-code-2022/pkg/types"
)

type Cave struct {
	Rocks  [][]bool
	Entry  *t.Coordinate
	Offset int
}

func NormalisedCave(rocks []*t.Coordinate, xMin, xMax, yMax int, endless bool) *Cave {
	const xSandEntry = 500
	const yFloorIncrease = 2

	if !endless {
		// increase min and max values
		yMax += yFloorIncrease
		xMin = xSandEntry - yMax
		xMax = xSandEntry + yMax
		// draw bottom line
		for x := xMin; x <= xMax; x++ {
			rocks = append(rocks, t.NewCoordinate(x, yMax))
		}
	}

	// normalize rock matrix (to 0,0)
	c := &Cave{
		Rocks:  make([][]bool, yMax+1),
		Entry:  t.NewCoordinate(xSandEntry-xMin, 0),
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
		for _, target := range fallOptions(current) {
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

func (c *Cave) valid(target *t.Coordinate) bool {
	return target.X >= 0 && target.X < len(c.Rocks[0]) && target.Y >= 0 && target.Y < len(c.Rocks)
}

func (c *Cave) available(target *t.Coordinate) bool {
	return !c.Rocks[target.Y][target.X]
}

func fallOptions(c *t.Coordinate) []*t.Coordinate {
	down := t.NewCoordinate(c.X, c.Y+1)
	diagonalLeft := t.NewCoordinate(c.X-1, c.Y+1)
	diagonalRight := t.NewCoordinate(c.X+1, c.Y+1)
	return []*t.Coordinate{down, diagonalLeft, diagonalRight}
}
