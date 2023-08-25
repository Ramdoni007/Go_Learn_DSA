package Go_learning_DSA

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{10, 5, 5},
		{25, 5, 5},
		{30, 15, 15},
		{30, 9, 3},
		{10, 5, 5},
		{100, 9, 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			got := GCD(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("GCD () = %v want %v ", got, tc.want)
			}
		})
	}
}
