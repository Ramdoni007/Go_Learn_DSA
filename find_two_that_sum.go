package Go_learning_DSA

func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, val := range numbers {
		for j, val2 := range numbers {
			if i == j {
				continue
			}
			if val+val2 == sum {
				return i, j
			}
		}
	}

	return -1, -1
}
