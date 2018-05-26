package timecomplexity

// PermMissingElem is used to find the missing element
// in a array with N elements which contains integers in the range [1..(N+1)]
func PermMissingElem(A []int) int {
	length := len(A)
	expected := (2 + length) * (1 + length) / 2
	sum := 0
	for i := 0; i < length; i++ {
		sum += A[i]
	}
	return expected - sum
}
