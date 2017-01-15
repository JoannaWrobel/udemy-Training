package main

import "fmt"

func main()  {
	for i := 6000; i <= 6005; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

}
