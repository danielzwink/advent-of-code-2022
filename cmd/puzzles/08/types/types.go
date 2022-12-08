package types

type Tree struct {
	Height      int
	Visible     bool
	ScenicScore int
}

func NewTree(height int) *Tree {
	return &Tree{
		Height:      height,
		Visible:     true,
		ScenicScore: 1,
	}
}

type TreeMap [][]*Tree

func (m TreeMap) CalculateVisibilitiesAndScenicScores() {
	maxY := len(m)
	for y := 1; y < maxY-1; y++ {
		maxX := len(m[y])
		for x := 1; x < maxX-1; x++ {
			currentHeight := m[y][x].Height

			leftVisibility := true
			leftViewingDistance := 0
			for xl := x - 1; xl >= 0; xl-- {
				leftViewingDistance++
				if m[y][xl].Height >= currentHeight {
					leftVisibility = false
					break
				}
			}

			rightVisibility := true
			rightViewingDistance := 0
			for xr := x + 1; xr < maxX; xr++ {
				rightViewingDistance++
				if m[y][xr].Height >= currentHeight {
					rightVisibility = false
					break
				}
			}

			upVisibility := true
			upViewingDistance := 0
			for yt := y - 1; yt >= 0; yt-- {
				upViewingDistance++
				if m[yt][x].Height >= currentHeight {
					upVisibility = false
					break
				}
			}

			downVisibility := true
			downViewingDistance := 0
			for yb := y + 1; yb < maxY; yb++ {
				downViewingDistance++
				if m[yb][x].Height >= currentHeight {
					downVisibility = false
					break
				}
			}

			m[y][x].Visible = leftVisibility || rightVisibility || upVisibility || downVisibility
			m[y][x].ScenicScore = leftViewingDistance * rightViewingDistance * upViewingDistance * downViewingDistance
		}
	}
}

func (m TreeMap) VisibleTreeCount() int {
	visibleTrees := 0
	for _, row := range m {
		for _, tree := range row {
			if tree.Visible {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func (m TreeMap) HighestScenicScore() int {
	highestScore := 1
	for _, row := range m {
		for _, tree := range row {
			if tree.Visible && tree.ScenicScore > highestScore {
				highestScore = tree.ScenicScore
			}
		}
	}
	return highestScore
}
