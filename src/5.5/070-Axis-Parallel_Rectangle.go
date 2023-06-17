package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func countPoints(n int, x, y []int, xl, xr, yh, yl int) int {
	count := 0
	for i := 0; i < n; i++ {
		if xl <= x[i] && x[i] <= xr && yl <= y[i] && y[i] <= yh {
			count++
		}
	}
	return count
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	k := scan()
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = scan()
		y[i] = scan()
	}

	ans := 4000000000000000000
	for left := 0; left < n; left++ {
		for right := 0; right < n; right++ {
			for high := 0; high < n; high++ {
				for low := 0; low < n; low++ {
					xl := x[left]
					xr := x[right]
					yh := y[high]
					yl := y[low]
					if countPoints(n, x, y, xl, xr, yh, yl) >= k {
						area := (xr - xl) * (yh - yl)
						if area < ans {
							ans = area
						}
					}
				}
			}
		}
	}

	fmt.Println(ans)
}
