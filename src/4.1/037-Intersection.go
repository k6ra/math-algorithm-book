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

func cross(x1, y1, x2, y2 int) int {
	return x1*y2 - y1*x2
}

func minMax(x1, y1, x2, y2 int) (int, int, int, int) {
	if x1 < x2 || (x1 == x2 && y1 < y2) {
		return x1, y1, x2, y2
	}
	return x2, y2, x1, y1
}

func lessEqual(x1, y1, x2, y2 int) bool {
	return x1 < x2 || (x1 == x2 && y1 < y2) || (x1 == x2 && y1 == y2)
}

func main() {
	sc.Split(bufio.ScanWords)

	x1 := scan()
	y1 := scan()

	x2 := scan()
	y2 := scan()

	x3 := scan()
	y3 := scan()

	x4 := scan()
	y4 := scan()

	cross1 := cross(x2-x1, y2-y1, x3-x1, y3-y1)
	cross2 := cross(x2-x1, y2-y1, x4-x1, y4-y1)
	cross3 := cross(x4-x3, y4-y3, x1-x3, y1-y3)
	cross4 := cross(x4-x3, y4-y3, x2-x3, y2-y3)
	if cross1 == 0 && cross2 == 0 && cross3 == 0 && cross4 == 0 {
		min12X, min12Y, max12X, max12Y := minMax(x1, y1, x2, y2)
		min34X, min34Y, max34X, max34Y := minMax(x3, y3, x4, y4)
		_, _, maxX, maxY := minMax(min12X, min12Y, min34X, min34Y)
		minX, minY, _, _ := minMax(max12X, max12Y, max34X, max34Y)

		if lessEqual(maxX, maxY, minX, minY) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}

	cross12 := cross1 >= 0 && cross2 <= 0
	cross12 = cross12 || (cross1 <= 0 && cross2 >= 0)

	cross34 := cross3 >= 0 && cross4 <= 0
	cross34 = cross34 || (cross3 <= 0 && cross4 >= 0)

	if cross12 && cross34 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
