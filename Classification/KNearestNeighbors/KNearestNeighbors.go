package KNearestNeighbors

import (
	"sort"
)

// xs must be at least 2 items
func getLabel(distances []float64, distanceToLabel map[float64]string) string {

	sort.Float64Slice(distances).Sort()

	var max float64
	var labelGuess string

	{
		labelCount := make(map[string]float64)
		step := (distances[1] - distances[0])
		i := 0
		last := distances[0]

		for last-distances[i] < (step*5) && i < len(distances)/10.0 {
			label := distanceToLabel[distances[i]]
			if _, ok := labelCount[label]; ok {
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
