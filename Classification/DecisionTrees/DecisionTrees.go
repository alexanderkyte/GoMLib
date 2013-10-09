package GoMLib

import (
  "math"
)

func shannonEntropy(input []string) (shannonEnt float64) {
	inputLen := len(input)
	labels := make(map[string]int)

	for _, val := range input {
		if _, ok := labels[val]; ok {
			labels[val] += 1
		} else {
			labels[val] = 0
		}
	}

	shannonEnt = 0
	for _, val := range labels {
		prob := float64(val) / float64(inputLen)
		shannonEnt = shannonEnt - (prob * math.Log2(prob))
	}

	return
}

