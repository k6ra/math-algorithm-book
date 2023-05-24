package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	r := 0
	y := 0
	b := 0
	for _, v := range a {
		switch v {
		case 1:
			r++
		case 2:
			y++
		case 3:
			b++
		}
	}

	rC := 0
	if r > 1 {
		rC = r * (r - 1) / 2
	}
	yC := 0
	if y > 1 {
		yC = y * (y - 1) / 2
	}
	bC := 0
	if b > 1 {
		bC = b * (b - 1) / 2
	}
	fmt.Println(rC + yC + bC)
}
