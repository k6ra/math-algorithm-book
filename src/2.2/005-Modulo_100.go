package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		sum += a
	}

	fmt.Println(sum % 100)
}
