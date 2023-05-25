package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() float64 {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return float64(n)
}

func main() {
	sc.Split(bufio.ScanWords)
	ax := scan()
	ay := scan()
	bx := scan()
	by := scan()
	cx := scan()
	cy := scan()

	bax := bx - ax
	bay := by - ay
	bcx := bx - cx
	bcy := by - cy
	bip := bax*bcx + bay*bcy

	cax := cx - ax
	cay := cy - ay
	cbx := cx - bx
	cby := cy - by
	cip := cax*cbx + cay*cby

	if bip < 0 {
		fmt.Printf("%f", math.Sqrt(bax*bax+bay*bay))
		return
	}
	if cip < 0 {
		fmt.Printf("%f", math.Sqrt(cax*cax+cay*cay))
		return
	}

	oip := bax*bcy - bay*bcx
	if oip < 0 {
		oip *= -1
	}

	fmt.Printf("%f", oip/math.Sqrt(bcx*bcx+bcy*bcy))
}
