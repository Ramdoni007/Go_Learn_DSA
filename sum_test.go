package Go_learning_DSA

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	testCase := []struct {
		numbers []int
		want    int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3, 4, 5}, 120},
		{[]int{}, 0},
	}

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("(%v)", tc.numbers), func(t *testing.T) {
			got := Sum(tc.numbers)
			if got != tc.want {
				t.Fatalf("SUM () = %v want %v ", got, tc.want)
			}
		})
	}
}
