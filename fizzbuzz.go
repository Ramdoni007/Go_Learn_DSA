package Go_learning_DSA

import "fmt"

//Fizzbuzz semcacam pola ouput yang ingin di hasilkan adalah angka dan string
// tetapi setiap 1 sampai 3 , 3 nya pasti berubah menjadi string  "Fizz"
// kemudia 5 nya berubah menjadi string "Buzz" dan nanti akan membentuk kalimat FizzBuzz dengan sendirinya

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("Fizz Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
