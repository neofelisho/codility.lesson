package iteration

// BinaryGap is a function to calculate the maximal sequence of consecutive zeros
// that is surrounded by ones at both ends in the binary representation of positive integer N.
func BinaryGap(N int) int {
	if N < 0 {
		return 0
	}

	result := 0
	temp := 0
	ok := false
	for N > 1 {
		current := N % 2
		N /= 2
		if current == 0 {
			if ok {
				temp++
			}
		} else {
			if temp > result {
				result = temp
			}
			temp = 0
			ok = true
		}
	}

	if N == 0 {
		temp++
	}
	if temp > result {
		result = temp
	}
	return result
}
