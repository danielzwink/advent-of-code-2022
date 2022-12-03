package types

type Rucksack struct {
	Compartment1 map[rune]int
	Compartment2 map[rune]int
	Items        map[rune]int
}

func NewRucksack(line string) *Rucksack {
	runes := []rune(line)
	size := len(runes)
	c1 := runes[:(size / 2)]
	c2 := runes[(size / 2):]

	return &Rucksack{
		Compartment1: countRunes(c1),
		Compartment2: countRunes(c2),
		Items:        countRunes(runes),
	}
}

func (r *Rucksack) SharedItem() rune {
	for item, _ := range r.Compartment1 {
		_, exists := r.Compartment2[item]
		if exists {
			return item
		}
	}
	panic(1)
}

func countRunes(runes []rune) map[rune]int {
	m := make(map[rune]int)
	for _, r := range runes {
		m[r]++
	}
	return m
}
