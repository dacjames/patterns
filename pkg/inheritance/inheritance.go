package inheritance

import (
	"fmt"
	"math"
)

type Shape struct {
	Shapely
}

type Shapely interface {
	Name() string
	Area() int
}

func (*Shape) Name() string {
	return "Unnamed"
}

func (s *Shape) Describe() string {
	return fmt.Sprintf("Shape: %s, Area: %d", s.Name(), s.Area())
}

type Circle struct {
	Shape
	Radius int
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Area() int {
	return int(math.Pi * float64(c.Radius) * float64(c.Radius))
}

func NewCircle(r int) *Circle {
	c := &Circle{Radius: r}
	c.Shape.Shapely = c
	return c
}

type Rectangle struct {
	Shape
	Width  int
	Height int
}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func NewRectangle(w, h int) *Rectangle {
	r := &Rectangle{Width: w, Height: h}
	r.Shape.Shapely = r
	return r
}

type Square struct {
	Shape
	Side int
}

func (*Square) Name() string {
	return "Square"
}

func (s *Square) Area() int {
	return s.Side * s.Side
}

func BadNewSquare(s int) *Square {
	// you get a nil pointer panic you don't construct properly
	return &Square{Side: s}
}

type BadTriangle struct {
	Shape
	Width  int
	Height int
}

func NewTriangle(w, h int) *BadTriangle {
	// you get a nil pointer panic you don't construct properly
	t := &BadTriangle{Width: w, Height: h}
	t.Shape.Shapely = t
	return t
}
