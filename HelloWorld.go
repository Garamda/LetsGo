package main

import "fmt"

func main() {
	var i, j, k int = 1, 2, 3
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)
	fmt.Println(i, j, k)
	fmt.Println("Hello, Garamda!")
}
