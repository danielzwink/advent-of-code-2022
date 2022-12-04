package types

import (
	"strconv"
	"strings"
)

type Section struct {
	LowerBound int
	UpperBound int
}

func (s *Section) containsPoint(point int) bool {
	return s.LowerBound <= point && s.UpperBound >= point
}

func (s *Section) overlapsSection(other *Section) bool {
	return s.containsPoint(other.LowerBound) || s.containsPoint(other.UpperBound)
}

func (s *Section) containsSection(other *Section) bool {
	return s.containsPoint(other.LowerBound) && s.containsPoint(other.UpperBound)
}

type Pair struct {
	Assignment1 *Section
	Assignment2 *Section
}

func NewPair(line string) *Pair {
	assignments := strings.Split(line, ",")
	bounds1 := strings.Split(assignments[0], "-")
	bounds2 := strings.Split(assignments[1], "-")

	a1Lb, _ := strconv.Atoi(bounds1[0])
	a1Ub, _ := strconv.Atoi(bounds1[1])
	a2Lb, _ := strconv.Atoi(bounds2[0])
	a2Ub, _ := strconv.Atoi(bounds2[1])

	return &Pair{
		Assignment1: &Section{
			LowerBound: a1Lb,
			UpperBound: a1Ub,
		},
		Assignment2: &Section{
			LowerBound: a2Lb,
			UpperBound: a2Ub,
		},
	}
}

func (p *Pair) FullyContained() bool {
	return p.Assignment1.containsSection(p.Assignment2) || p.Assignment2.containsSection(p.Assignment1)
}

func (p *Pair) Overlapped() bool {
	return p.Assignment1.overlapsSection(p.Assignment2) || p.Assignment2.overlapsSection(p.Assignment1)
}
