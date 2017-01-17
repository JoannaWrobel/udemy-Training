package main

import "fmt"

func main()  {
	var number1 int
	var number2 int
	fmt.Println("Enter small numer: ")
	fmt.Scan(&number1)
	fmt.Println("Enter larger numer: ")
	fmt.Scan(&number2)
	x := number2 % number1
	fmt.Println("The reminder of the division is", x)

}
