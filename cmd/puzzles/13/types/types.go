package types

import (
	"advent-of-code-2022/pkg/util"
)

type Pair struct {
	LeftPacket  *Element
	RightPacket *Element
}

func NewPair(left, right *Element) *Pair {
	return &Pair{LeftPacket: left, RightPacket: right}
}

type Element struct {
	Number int
	List   []*Element
}

func NumberElement(value string) *Element {
	return &Element{Number: util.MustParseInt(value)}
}

func ListElement() *Element {
	return &Element{List: make([]*Element, 0)}
}

func SingleElement(element *Element) *Element {
	listElement := ListElement()
	listElement.AddToListElement(element)
	return listElement
}

func (e *Element) AddToListElement(element *Element) {
	e.List = append(e.List, element)
}

func (e *Element) IsNumber() bool {
	return e.List == nil
}

func (e *Element) IsList() bool {
	return e.List != nil
}

func Compare(left, right *Element) int {
	if left.IsNumber() && right.IsNumber() {
		if left.Number < right.Number {
			return -1
		} else if left.Number > right.Number {
			return 1
		}
		return 0

	} else if left.IsList() && right.IsList() {
		leftLength := len(left.List)
		rightLength := len(right.List)
		smallerLength, _ := util.Sort(leftLength, rightLength)

		for i := 0; i < smallerLength; i++ {
			result := Compare(left.List[i], right.List[i])
			if result != 0 {
				return result
			}
		}

		if leftLength < rightLength {
			return -1
		} else if leftLength > rightLength {
			return 1
		}
		return 0

	} else if left.IsList() && right.IsNumber() {
		return Compare(left, SingleElement(right))

	} else if left.IsNumber() && right.IsList() {
		return Compare(SingleElement(left), right)
	}

	// the famous: should not happen
	panic(1)
}

type SortedElements []*Element

func (e SortedElements) Len() int {
	return len(e)
}

func (e SortedElements) Less(i, j int) bool {
	return Compare(e[i], e[j]) == -1
}

func (e SortedElements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
