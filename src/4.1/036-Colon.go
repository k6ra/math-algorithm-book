package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	a := float64(scan())
	b := float64(scan())
	h := float64(scan())
	m := float64(scan())

	angleH := 30*h + 0.5*m
	hRad := angleH * math.Pi / 180.0
	hX := a * math.Cos(hRad)
	hY := a * math.Sin(hRad)

	angleM := 6 * m
	mRad := angleM * math.Pi / 180.0
	mX := b * math.Cos(mRad)
	mY := b * math.Sin(mRad)

	diffX := hX - mX
	diffY := hY - mY
	ans := math.Sqrt(diffX*diffX + diffY*diffY)
	fmt.Printf("%g", ans)
}
