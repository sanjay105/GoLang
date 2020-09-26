/*
Problem definition
An interesting problem in arithmetic with deep implications to elliptic curve
theory is the problem of nding perfect squares that are sums of consecutive
squares. A classic example is the Pythagorean identity:
32 + 42 = 52 (1)
that reveals that the sum of squares of 3; 4 is itself a square. A more interesting
example is Lucas` Square Pyramid :
12 + 22 + ::: + 242 = 702 (2)
In both of these examples, sums of squares of consecutive integers form the
square of another integer.

Go Implementation of finding Lucas Square Pyramid between 1 and N using go routines :

Run Command: "go run LucasPyramid.go <N> <k>"
*/package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func findLucasSquare(start int64, end int64, k int64, ch chan int64) {

	var i, sumi int64
	sumi = 0
	for i = 0; i < k; i++ {
		sumi = sumi + int64(math.Pow(float64(start+i), 2))
	}
	squareroot := math.Sqrt(float64(sumi))
	var sqrti int64 = int64(squareroot)
	if sumi == (sqrti * sqrti) {
		fmt.Println(start)
	}
	for i = start + k; i < end; i++ {
		sumi = sumi - int64(math.Pow(float64(i-k), 2)) + int64(math.Pow(float64(i), 2))
		squareroot := math.Sqrt(float64(sumi))
		var sqrti int64 = int64(squareroot)
		if sumi == (sqrti * sqrti) {
			fmt.Println(i - k + 1)
		}
	}
	ch <- 0 // one this functions finishes its job it puts 0 in the channel
}

func main() {
	ch := make(chan int64)
	var n, k int64
	ni, _ := strconv.Atoi(os.Args[1])
	ki, _ := strconv.Atoi(os.Args[2])
	n = int64(ni)
	k = int64(ki)
	var part, temp, i int64
	part = n / 8
	for i = 0; i < 8; i++ {
		go findLucasSquare(i*part, (i+1)*part, k, ch) // go routine is called
	}
	temp = 0
	for temp != 8 {
		a := <-ch
		if a == 0 { //if channel receives 0 message from 8 go routines it gets out of the loop and main thread exits
			temp++
		}
	}

}
