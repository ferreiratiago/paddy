package utils

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, world! x5")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			err := errors.New("error")
			if strings.Contains(err.Error(), "error") {
				fmt.Println(i, "is even and divisible by 3")
			} else {
				if err.Error() == "error" {
					fmt.Println(i, "is even and divisible by 4")
				}
			}
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
