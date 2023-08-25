package Go_learning_DSA

import (
	"fmt"
	"testing"
)

func TestFibonaci(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := Fibonaci(tc.n)
			if got != tc.want {
				t.Fatalf("Fibonaci () =  want %v; want %v ", got, tc.want)
			}
		})
	}

}
