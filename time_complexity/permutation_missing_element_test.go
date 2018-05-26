package timecomplexity

import "testing"

func TestPermMissingElem(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test 1", args{[]int{1, 2, 3, 5, 6}}, 4},
		{"Empty and single", args{[]int{2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermMissingElem(tt.args.A); got != tt.want {
				t.Errorf("PermMissingElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
