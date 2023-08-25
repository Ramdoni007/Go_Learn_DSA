package main

import (
	"fmt"
	Go_learning_DSA "github.com/ramdoni007/Go_learn_DSA"
)

// Program ini akan melakukan proses GCD dengan otomatis membuat file output dari bilangan ouput .txt
func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := Go_learning_DSA.GCD(a, b)
		fmt.Println(gcd)
	}
}
