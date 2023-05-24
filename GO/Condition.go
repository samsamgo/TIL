package main

import "fmt"

func main() {
	number := 10

	if number > 10 {
		fmt.Println("the nember is greater than 10.")
	} else if number == 10 {
		fmt.Println("the nember is equal to 10")
	} else {
		fmt.Println("the nember is less than 10")
	}
}
