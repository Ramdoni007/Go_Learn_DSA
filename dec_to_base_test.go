package Go_learning_DSA

import (
	"fmt"
	"testing"
)

func TestDecToBase(t *testing.T) {
	testCases := []struct {
		dec, base int
		want      string
	}{
		{1, 2, "1"},
		{2, 2, "10"},
		{7, 3, "21"},
		{14, 2, "1110"},
		{14, 16, "E"},
		{17, 16, "11"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v to base %v ", tc.dec, tc.base), func(t *testing.T) {
			got := DecToBase(tc.dec, tc.base)
			if got != tc.want {
				t.Fatalf("DecToBase() %v; want %v", got, tc.want)
			}
		})
	}
}
