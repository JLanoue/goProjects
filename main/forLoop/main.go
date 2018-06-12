package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "This number is divisible by 3 and 5")
		} else if i%5 == 0 {
			fmt.Println(i, "This number is divisible by 5")
		} else if i%3 == 0 {
			fmt.Println(i, "This number is divisible by 3")
		} else {
			fmt.Println(i)
		}
	}
}
