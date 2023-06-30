package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	count := 0
	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}

		if count < 0 {
			fmt.Println("No")
			return
		}
	}

	if count == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
