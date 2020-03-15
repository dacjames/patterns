package inheritance

import (
	"fmt"
	"math"
)

// Shape represents the "base" or "parent" class in traditional OO
type Shape struct {
	// THis is where the magic happens. We're embedding an interface member
	// of this struct, so when Go resolves method calls, it will check for them
	// on this embedded interface. When constructing the "child" classes,
	// we set this embedded interface equal to a pointer to the struct itself.
	// This sets up the method resolution chain to match the inheritance relationship
	// that we're interested in.
	Shapely
}

// Shapely is the interface of methods that we want to be inheritable.
type Shapely interface {
	Name() string
	Area() int
}

// Name() implements the default implementation on the parent type.
func (*Shape) Name() string {
	return "Unnamed"
}

// Note that Shape does NOT define the Area method. This is done
// to demonstrate what happens when the child does not implement the expected method.

// Describe is the method that needs to reference the method implementations in the children
// types. If normal embedding were used, these methods would resolve on only the Shape type,
// not on Circle or Rectange that embed Shape. Using this pattern, methods will instead resolve
// via the embedded `Shapely` interface, which is set to reference the child type.
// This method simulates any time you want generic functinality to be implemented in terms
// of type-specific functionality.
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
	// you get a nil pointer panic in .Describe()
	return &Square{Side: s}
}

type BadTriangle struct {
	Shape
	Width  int
	Height int
}

// Note that we do not implement either the Name() or Area() methods.
// The Name() method will resolve to the "default" implementation defined
// on the Shape type. Trying to resolve the  missing Area() method will
// result in an infinite recursion:
//   BadTriange.Area() => BadTriangle.Shape.Area()
//   => Shape.Shapely.Area() => BadTriangle.Area() => ...
// The recursion here results from setting the
func NewTriangle(w, h int) *BadTriangle {
	// you get a nil pointer panic you don't construct properly
	t := &BadTriangle{Width: w, Height: h}
	t.Shape.Shapely = t
	return t
}
