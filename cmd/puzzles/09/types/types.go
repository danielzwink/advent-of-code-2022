package types

import "fmt"

type Step struct {
	X int
	Y int
}

func NewStep(direction rune) *Step {
	switch direction {
	case 'L':
		return &Step{X: -1, Y: 0}
	case 'R':
		return &Step{X: 1, Y: 0}
	case 'U':
		return &Step{X: 0, Y: 1}
	case 'D':
		return &Step{X: 0, Y: -1}
	default:
		panic(1)
	}
}

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) *Position {
	return &Position{X: x, Y: y}
}

func (c *Position) Move(step *Step) {
	c.X += step.X
	c.Y += step.Y
}

func (c *Position) Tighten(other *Position) {
	// overlapping
	if c.X == other.X && c.Y == other.Y {
		return
	}

	otherXp1 := other.X + 1
	otherXm1 := other.X - 1
	otherYp1 := other.Y + 1
	otherYm1 := other.Y - 1

	// still in touch
	if c.X >= otherXm1 && c.X <= otherXp1 && c.Y >= otherYm1 && c.Y <= otherYp1 {
		return
	}

	otherXp2 := other.X + 2
	otherXm2 := other.X - 2
	otherYp2 := other.Y + 2
	otherYm2 := other.Y - 2

	// two steps apart, one shared dimension
	if c.X == otherXp2 && c.Y == other.Y {
		other.X++
		return
	}
	if c.X == otherXm2 && c.Y == other.Y {
		other.X--
		return
	}
	if c.X == other.X && c.Y == otherYp2 {
		other.Y++
		return
	}
	if c.X == other.X && c.Y == otherYm2 {
		other.Y--
		return
	}

	// jump required
	if c.X == otherXp2 && (c.Y == otherYm1 || c.Y == otherYp1) {
		other.X++
		other.Y = c.Y
		return
	}
	if c.X == otherXm2 && (c.Y == otherYm1 || c.Y == otherYp1) {
		other.X--
		other.Y = c.Y
		return
	}
	if c.Y == otherYp2 && (c.X == otherXm1 || c.X == otherXp1) {
		other.X = c.X
		other.Y++
		return
	}
	if c.Y == otherYm2 && (c.X == otherXm1 || c.X == otherXp1) {
		other.X = c.X
		other.Y--
		return
	}

	// two steps apart, diagonal
	if c.X == otherXp2 && c.Y == otherYp2 {
		other.X++
		other.Y++
		return
	}
	if c.X == otherXp2 && c.Y == otherYm2 {
		other.X++
		other.Y--
		return
	}
	if c.X == otherXm2 && c.Y == otherYp2 {
		other.X--
		other.Y++
		return
	}
	if c.X == otherXm2 && c.Y == otherYm2 {
		other.X--
		other.Y--
		return
	}
}

func (c *Position) Key() string {
	return fmt.Sprintf("%v-%v", c.X, c.Y)
}
