package types

type WorryLevel int

func (l WorryLevel) divisible(divisor int) bool {
	return int(l)%divisor == 0
}

type Forward struct {
	Monkey int
	Item   WorryLevel
}

func NewForward(monkey int, item WorryLevel) *Forward {
	return &Forward{Monkey: monkey, Item: item}
}

type Monkey struct {
	Inspections int

	items          []WorryLevel
	operation      func(WorryLevel) WorryLevel
	divisor        int
	divisibleTrue  int
	divisibleFalse int
}

func NewMonkey(items []WorryLevel, operation func(WorryLevel) WorryLevel, divisor, divisibleTrue, divisibleFalse int) *Monkey {
	return &Monkey{
		Inspections:    0,
		items:          items,
		operation:      operation,
		divisor:        divisor,
		divisibleTrue:  divisibleTrue,
		divisibleFalse: divisibleFalse,
	}
}

func (m *Monkey) Turn(useModulo int) []*Forward {
	forwards := make([]*Forward, 0)

	for _, item := range m.items {
		m.Inspections++

		// worry level operation
		item = m.operation(item)

		// worry level relief
		if useModulo == 0 {
			item = item / 3
		} else {
			item = item % WorryLevel(useModulo)
		}

		// throw target
		forwards = append(forwards, m.target(item))
	}

	m.items = make([]WorryLevel, 0)
	return forwards
}

func (m *Monkey) target(item WorryLevel) *Forward {
	if item.divisible(m.divisor) {
		return NewForward(m.divisibleTrue, item)
	} else {
		return NewForward(m.divisibleFalse, item)
	}
}

func (m *Monkey) AddItem(item WorryLevel) {
	m.items = append(m.items, item)
}
