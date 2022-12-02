package types

type Round struct {
	Opponent Move
	Self     Move
	Outcome  Result
}

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissors
)

type Result int

const (
	Lost Result = 0
	Draw Result = 3
	Win  Result = 6
)

func (r *Round) EvaluateOutcome() {
	if r.Opponent == r.Self {
		r.Outcome = Draw
	} else if (r.Opponent == Rock && r.Self == Paper) || (r.Opponent == Paper && r.Self == Scissors) || (r.Opponent == Scissors && r.Self == Rock) {
		r.Outcome = Win
	} else if (r.Opponent == Rock && r.Self == Scissors) || (r.Opponent == Paper && r.Self == Rock) || (r.Opponent == Scissors && r.Self == Paper) {
		r.Outcome = Lost
	} else {
		panic(1)
	}
}

func (r *Round) EvaluateSelf() {
	if r.Outcome == Draw {
		r.Self = r.Opponent
	} else if (r.Opponent == Rock && r.Outcome == Lost) || (r.Opponent == Paper && r.Outcome == Win) {
		r.Self = Scissors
	} else if (r.Opponent == Rock && r.Outcome == Win) || (r.Opponent == Scissors && r.Outcome == Lost) {
		r.Self = Paper
	} else if (r.Opponent == Paper && r.Outcome == Lost) || (r.Opponent == Scissors && r.Outcome == Win) {
		r.Self = Rock
	} else {
		panic(1)
	}
}

func (r *Round) Score() int {
	return int(r.Self) + int(r.Outcome)
}
