package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

type cell struct {
	x int
	y int
}

var row []string
var cost [][]int
var queue []cell

func isNextCell(y int, x int, nextCost int) bool {
	return row[y][x] == '.' && cost[y][x] == -1
}

func pushNext(y int, x int, nextCost int) {
	yDirect := []int{0, 0, 1, -1}
	xDirect := []int{1, -1, 0, 0}

	for i := 0; i < 4; i++ {
		nextY := y + yDirect[i]
		nextX := x + xDirect[i]
		if isNextCell(nextY, nextX, nextCost) {
			cost[nextY][nextX] = nextCost
			queue = append(queue, cell{y: nextY, x: nextX})
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	r := scanInt()
	c := scanInt()
	start := cell{
		y: scanInt(),
		x: scanInt(),
	}
	goal := cell{
		y: scanInt(),
		x: scanInt(),
	}
	row = make([]string, r+1)
	for i := 1; i <= r; i++ {
		row[i] = " " + scanString()
	}

	cost = make([][]int, r+1)
	for i := 0; i <= r; i++ {
		cost[i] = make([]int, c+1)
		for j := 0; j <= c; j++ {
			cost[i][j] = -1
		}
	}

	queue = make([]cell, 0, r*c)
	queue = append(queue, start)
	cost[start.y][start.x] = 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		nextCost := cost[current.y][current.x] + 1
		pushNext(current.y, current.x, nextCost)
	}

	fmt.Println(cost[goal.y][goal.x])
}
