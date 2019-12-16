package main

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) Distance() int {
	return p.x + p.y
}

func main() {
	p1 := Point{1, 2}
	p2 := &Point{3, 4}
	fmt.Println(p1.Distance())
	fmt.Println(p2.Distance())
}
