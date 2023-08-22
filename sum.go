package Go_learning_DSA

func Sum(numbers []int) int {
	res := 0
	for _, val := range numbers {
		res += val
	}
	return res
}
