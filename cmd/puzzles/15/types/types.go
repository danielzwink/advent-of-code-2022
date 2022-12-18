package types

import (
	t "advent-of-code-2022/pkg/types"
	"advent-of-code-2022/pkg/util"
)

type Pair struct {
	Sensor   *t.Coordinate
	Beacon   *t.Coordinate
	Distance int
}

func NewPair(sensorX, sensorY, beaconX, beaconY string) *Pair {
	sensor := &t.Coordinate{X: util.MustParseInt(sensorX), Y: util.MustParseInt(sensorY)}
	beacon := &t.Coordinate{X: util.MustParseInt(beaconX), Y: util.MustParseInt(beaconY)}
	return &Pair{
		Sensor:   sensor,
		Beacon:   beacon,
		Distance: sensor.ManhattanDistance(beacon),
	}
}

func (p *Pair) OverlapsY(y int) (bool, int, int) {
	distance := p.Sensor.ManhattanDistance(p.Beacon)
	yMin := p.Sensor.Y - distance
	yMax := p.Sensor.Y + distance
	overlap := yMin <= y && y <= yMax

	if overlap {
		yDeviation := util.Abs(p.Sensor.Y - y)
		xDeviation := distance - yDeviation
		xMin := p.Sensor.X - xDeviation
		xMax := p.Sensor.X + xDeviation
		return true, xMin, xMax
	} else {
		return false, 0, 0
	}
}

func (p *Pair) Contains(c *t.Coordinate) bool {
	target := p.Sensor.ManhattanDistance(c)
	return target <= p.Distance
}
