package main

import "fmt"

func main() {
	number := 1

	switch number {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("기타")
	}
}
