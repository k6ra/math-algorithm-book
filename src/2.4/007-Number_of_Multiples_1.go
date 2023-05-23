package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scanf("%d %d %d", &n, &x, &y)
	count := 0
	for i := 1; i <= n; i++ {
		if i%x == 0 || i%y == 0 {
			count++
		}
	}
	fmt.Println(count)
}
