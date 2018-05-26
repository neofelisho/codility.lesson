package arrays

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
