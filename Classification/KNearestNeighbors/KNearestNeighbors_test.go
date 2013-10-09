package GoMLib

import (
 "testing"
 "math"
)

type Testable struct {
  Fields [3]float64
	Type string
}

func (t Testable) Compare(ts []Testable) (distances []float64, distanceToLabel map[float64]string) {
  distances = make([]float64, len(ts))
  distanceToLabel = make(map[float64]string)

	for i := 0; i < len(ts); i++ {
		distance := float64(0)

		for j := 0; j < len(t.Fields); j++ {
			distance += math.Pow((ts[i].Fields[j] - t.Fields[j]), 2)
		}

		distance = math.Sqrt(distance)
		distances[i] = distance
		distanceToLabel[distance] = ts[i].Type
	}
  return
}

func Test_main(t *testing.T) {
	labels := make([]Testable, 100)

	for i := 0; i < 90; i++ {
		labels[i] = Testable{[3]float64{float64(i), float64(i + 1), float64(i + 2)}, "Wrong"}
	}
	for i := 91; i < 100; i++ {
		labels[i] = Testable{[3]float64{float64(200 + i), float64(200 + i), float64(200 + i)}, "Right"}
	}
  test := Testable{[3]float64{203.0, 203.0, 204.0}, ""}

	if getLabel(test.Compare(labels)) != "Right" {
		t.Errorf("Not working")
	}

  test = Testable{[3]float64{3.0, 3.0, 4.0}, ""}
	if getLabel(test.Compare(labels)) != "Wrong" {
		t.Errorf("Not working")
	}
  
}
