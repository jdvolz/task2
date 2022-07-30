package main

import "testing"

func TestSubtraction(t *testing.T) {
	doubleSetOne := DoubleSet{1: 2, 2: 1, 4: 1}
	doubleSetTwo := DoubleSet{1: 1, 2: 2, -3: 1}

	result := doubleSetOne.Subtract(doubleSetTwo)

	target := DoubleSet{1: 1, 4: 1}

	for tk, tv := range target {
		if v, ok := result[tk]; !ok {
			t.Fatalf("%v is a missing key", tk)
		} else {
			if v != tv {
				t.Fatalf("1 has value %v when it should be %v", v, tv)
			}
		}
	}
	if len(result) != len(target) {
		t.Fatalf("result has a different length than the target")
	}
}

func TestAdd(t *testing.T) {
	doubleSetOne := DoubleSet{1: 2, 2: 1}
	doubleSetTwo := DoubleSet{1: 1, 2: 1, -3: 1}

	result := doubleSetOne.Add(doubleSetTwo)

	target := DoubleSet{1: 2, 2: 2, -3: 1}

	for tk, tv := range target {
		if v, ok := result[tk]; !ok {
			t.Fatalf("%v is a missing key", tk)
		} else {
			if v != tv {
				t.Fatalf("1 has value %v when it should be %v", v, tv)
			}
		}
	}
	if len(result) != len(target) {
		t.Fatalf("result has a different length than the target")
	}
}
