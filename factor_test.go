package Go_learning_DSA

import (
	"fmt"
	"sort"
	"testing"
)

func TestFactor(t *testing.T) {
	tenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 31}
	testCases := []struct {
		primes []int
		number int
		want   []int
	}{
		{tenPrimes, 4, []int{2, 2}},
		{tenPrimes, 30, []int{2, 3, 5}},
		{tenPrimes, 28, []int{2, 2, 7}},
		{tenPrimes, 720, []int{2, 2, 2, 2, 3, 3, 5}},
		{tenPrimes, 12, []int{2, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v with primes %v", tc.number, tc.primes), func(t *testing.T) {
			got := Factor(tc.primes, tc.number)
			sort.Ints(got)
			sort.Ints(tc.want)
			if err := inSLiceEquals(got, tc.want); err != nil {
				t.Fatalf("Factor () sorted = %v ; want =  %v err = %v ", got, tc.want, err)
			}
		})
	}

}
