package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triange struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{sideLength: 2}
	t := triange{
		base:   2,
		height: 2,
	}

	printArea(s)
	printArea(t)
}
