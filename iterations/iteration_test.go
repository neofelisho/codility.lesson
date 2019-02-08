package iteration

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{"Test for -1", -1, 0},
		{"Test for 32", 32, 0},
		{"Test for 15", 15, 0},
		{"Test for 1041", 1041, 5},
		{"Test for 1043", 1043, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryGap(tt.args); got != tt.want {
				t.Errorf("BinaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
