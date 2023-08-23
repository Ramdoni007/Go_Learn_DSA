package Go_learning_DSA

func Reverse(word string) string {

	// with original strategy
	//var res string
	//
	//for i := len(word) - 1; i >= 0; i-- {
	//	res = res + string(word[i])
	//}
	//
	//return res

	// with stringBuilder strategy zero memory value
	//	var sb strings.Builder
	//	//logic of reverse string :)
	//	for i := len(word) - 1; i >= 0; i-- {
	//		sb.WriteByte(word[i])
	//	}
	//
	//	return sb.String()
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
}
