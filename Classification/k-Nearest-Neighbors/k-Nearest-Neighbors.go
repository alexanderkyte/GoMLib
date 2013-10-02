package KNearestNeighbors

import (
	"math"
)

type KComparable interface {
	YieldComparing() []float64
	YieldLabel() String
}

// xs must be at least 2 items
func getLabel(x KComparable, xs []KComparable) KComparable {
	distanceToLabel := make(map[float64]string)
	distances := new([len(xs)]float64)
	xFields := x.YieldComparing()

	for i := 0; i < len(xs); i++ {
		distance := float64(0)
		xthFields := xs[i].YieldComparing()

		for j := 0; j < len(xFields); j++ {
			distance += math.Pow((xFields[j] - xthFields[j]), 2)
		}

		distance = math.Sqrt(distance)
		distances[i] = distance
		distanceToLabel[distance] = xs[i].YieldLabel()
	}

	sort.Sort(distances)

	var max float64
	var labelGuess string

	{
		labelCount = make(map[string]float)
		step := (distances[1] - distances[0])
		i := 0
		last := distances[0]

		for last-distances[i] < (step*5) && i < math.Floor(len(distances)/10) {
			label := distanceToLabel[distances[i]]
			if labelCount[label] {
				labelCount[label] += 1
			} else {
				labelCount[label] = 0
			}
			if labelCount[label] > max {
				max = labelCount[label]
				labelGuess = label
			}
			last = distances[i]
			i++
		}
	}
	return labelGuess

}
