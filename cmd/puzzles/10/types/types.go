package types

type Cycle struct {
	Summand int
}

func NewCycle(value int) *Cycle {
	return &Cycle{Summand: value}
}

func EmptyCycle() *Cycle {
	return NewCycle(0)
}
