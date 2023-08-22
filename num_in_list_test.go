package Go_learning_DSA

import (
	"fmt"
	"testing"
)

func TestNumInList(t *testing.T) {
	testCase := []struct {
		list []int
		num  int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, -1}, -1, true},
	}
	for _, tc := range testCase {
		t.Run(fmt.Sprintf("(%v,%v)", tc.list, tc.num), func(t *testing.T) {
			got := NumInList(tc.list, tc.num)
			if got != tc.want {
				t.Fatalf("NumInList () = %v want %v ", got, tc.want)
			}
		})
	}
}
