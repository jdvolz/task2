package main

import "testing"

func TestSubtraction(t *testing.T) {
	doubleSetOne := DoubleSet{1: 2, 2: 1, 4: 1}
	doubleSetTwo := DoubleSet{1: 1, 2: 2, -3: 1}

	result := doubleSetOne.Subtract(doubleSetTwo)

	target := DoubleSet{1: 1, 4: 1}

	isEqual(target, result, t)
}

func isEqual(ds1 DoubleSet, ds2 DoubleSet, t *testing.T) {

	for tk, tv := range ds1 {
		if v, ok := ds2[tk]; !ok {
			t.Fatalf("%v is a missing key", tk)
		} else {
			if v != tv {
				t.Fatalf("1 has value %v when it should be %v", v, tv)
			}
		}
	}
	if len(ds2) != len(ds1) {
		t.Fatalf("result has a different length than the target")
	}
}

func TestAdd(t *testing.T) {
	doubleSetOne := DoubleSet{1: 2, 2: 1}
	doubleSetTwo := DoubleSet{1: 1, 2: 1, -3: 1}

	result := doubleSetOne.Add(doubleSetTwo)

	target := DoubleSet{1: 2, 2: 2, -3: 1}

	isEqual(target, result, t)
}
