package arrays

import (
	"reflect"
	"testing"
)

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

func TestCyclicRotation(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Test 1", args{A: []int{3, 8, 9, 7, 6}, K: 3}, []int{9, 7, 6, 3, 8}},
		{"Test 2", args{A: []int{5, -1000}, K: 11}, []int{-1000, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyclicRotation(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CyclicRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOddOccurrencesInArray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", []int{9, 3, 9, 3, 9, 7, 9}, 7},
		{"Test 2", []int{4, 1, 1, 1, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddOccurrencesInArray(tt.args); got != tt.want {
				t.Errorf("OddOccurrencesInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
