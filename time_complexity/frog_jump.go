package timecomplexity

// FrogJump is used to count the minimal number of jumps
// that the small frog must perform to reach its target.
func FrogJump(X int, Y int, D int) int {
	jumps := (Y - X) / D
	remains := (Y - X) % D
	if remains == 0 {
		return jumps
	}
	return jumps + 1
}
