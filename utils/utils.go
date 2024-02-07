package utils

import "fmt"

func PrintEvenOrOdd(a int) {
	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func PrintRange(low, high int) {
	for i := low; i <= high; i++ {
		fmt.Println(i)
	}
}
