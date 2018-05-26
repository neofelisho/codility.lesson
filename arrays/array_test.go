package arrays

import "testing"

func Test_CyclicRotation(t *testing.T) {
	input := []int{3, 8, 9, 7, 6}
	expected := []int{9, 7, 6, 3, 8}
	if actual := CyclicRotation(input, 3); arrayEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Failed to CyclicRotation.")
	}

	input = []int{5, -1000}
	expected = []int{-1000, 5}
	if actual := CyclicRotation(input, 11); arrayEqual(actual, expected) {
		t.Log("Success")
	} else {
		t.Error("Failed to CyclicRotation.")
	}
}

func arrayEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
