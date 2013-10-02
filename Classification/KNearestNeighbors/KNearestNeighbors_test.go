package KNearestNeighbors 

import (
 "testing"
)

type Testable struct {
	X    int
	Y    int
	Z    int
	Type string
}

func (t Testable) YieldComparing() []float64 {
	return []float64{float64(t.X), float64(t.Y), float64(t.Z)}
}

func (t Testable) YieldLabel() string {
	return t.Type
}

func Test_main(t *testing.T) {
	labels := new([100]KComparable)

	for i := 0; i < 90; i++ {
		labels[i] = Testable{i, i + 1, i + 2, "Wrong"}
	}
	for i := 91; i < 100; i++ {
		labels[i] = Testable{200 + i, 200 + i, 200 + i, "Right"}
	}

	if getLabel(Testable{203, 203, 204, ""}, labels[:]) != "Right" {
		t.Errorf("Not working")
	}
}
