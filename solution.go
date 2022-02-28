package square

import (
	"math"
)

type side int

const (
	SidesTriangle side = 0
	SidesSquare   side = 3
	SidesCircle   side = 4
)

func CalcSquare(sideLen float64, sidesNum side) float64 {
	if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	}
	if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
