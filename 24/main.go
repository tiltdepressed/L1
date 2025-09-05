package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (current *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(current.GetX()-other.GetX(), 2) + math.Pow(current.GetY()-other.GetY(), 2))
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(3, 4)
	fmt.Println(a.Distance(b))
}
