package main

import "fmt"

func main() {
label:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break label

		}
		fmt.Println(i)
	}
	d := 0
label1:
	for {
		if d >= 10 {
			break label1
		}
		d++
	}
	fmt.Println("value of d :", d)
}
