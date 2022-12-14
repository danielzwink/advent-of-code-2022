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

func NewCoordinateCommaSeparated(value string) *Coordinate {
	split := strings.Split(value, ",")
	return NewCoordinate(util.MustParseInt(split[0]), util.MustParseInt(split[1]))
}

func (c *Coordinate) ManhattanDistance(o *Coordinate) int {
	return util.Abs(c.X-o.X) + util.Abs(c.Y-o.Y)
}

func (c *Coordinate) Add(o *Coordinate) *Coordinate {
	return NewCoordinate(c.X+o.X, c.Y+o.Y)
}
