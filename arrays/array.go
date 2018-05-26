package arrays

import (
	"sort"
)

// CyclicRotation means that each element in array is shifted right by one index,
// and the last element of the array is moved to the first place. For example,
// the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7]
// (elements are shifted right by one index and 6 is moved to the first place).
func CyclicRotation(A []int, K int) []int {
	lengthA := len(A)
	if lengthA > 100 || lengthA == 0 {
		return A
	}
	if K > 100 || K <= 0 {
		return A
	}
	shift := K % lengthA
	if shift == 0 {
		return A
	}
	result := make([]int, lengthA)
	for i := 0; i < lengthA; i++ {
		result[(i+shift)%lengthA] = A[i]
	}
	return result
}

// OddOccurrencesInArray A non-empty array A consisting of N integers is given.
// The array contains an odd number of elements,
// and each element of the array can be paired with another element that has the same value,
// except for one element that is left unpaired.
func OddOccurrencesInArray(A []int) int {
	sort.Ints(A)
	pre := 0
	for i := 0; i < len(A); i++ {
		if i%2 == 1 {
			pre -= A[i]
			if pre != 0 {
				return A[i-1]
			}
		} else {
			pre = A[i]
		}
	}
	return pre
}
