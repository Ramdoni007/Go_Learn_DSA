package Go_learning_DSA

func Factor(primes []int, number int) []int {
	var res []int
	for _, primes := range primes {
		for number%primes == 0 {
			res = append(res, number)
			number = number / primes
		}
	}

	if number > 1 {
		res = append(res, number)
	}
	return res
}
