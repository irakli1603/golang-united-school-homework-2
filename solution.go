package square

import (
	"math"
)

type side int

const (
	triangle side = 0
	square   side = 3
	circle   side = 4
)

func CalcSquare(sideLen float64, sidesNum side) float64 {
	if sidesNum == triangle {
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	}
	if sidesNum == square {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == circle {
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
