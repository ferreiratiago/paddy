package main

import "fmt"

func main() {
	fmt.Println("Hello, world! x5")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			if i%3 == 0 {
				fmt.Println(i, "is even and divisible by 3")
			} else {
				if i%4 == 0 {
					fmt.Println(i, "is even and divisible by 4")
				}
			}
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
