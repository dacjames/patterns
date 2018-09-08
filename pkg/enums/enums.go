// Package unions illustrates how to implement tagged unions in Go
// Let's start with simple enums that you're probably familiar with.
package enums

import (
	"fmt"
)

// The simplest form of enums using consts and `iota`
// Here, we encode the cardinal directions as ints
const (
	// North as in the North Star
	North int = iota
	// South as in South L.A.
	South
	// East as in East Coast
	East
	// West as in Best
	West
)

// Cardinal demonstrates how to use simple enums with the `switch` statement.
// This is simple, but comes with the risk of passing a cardinal that the
// function does not understand.
func Cardinal(card int) string {
	switch card {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

// Example1 shows how to use the Cardinal function to print names
func Example1() {
	// >>> "North South East West"
	fmt.Println(Cardinal(North), Cardinal(South), Cardinal(East), Cardinal(West))

	// But nothing stops you from trying to print non-existant constants
	fmt.Println(Cardinal(10), Cardinal(-1))
}

// Direction is introduced a newtype to mitigate this problem
type Direction int

const (
	// Left like all your stuff
	Left Direction = iota
	// Right as in not Wrong
	Right
	// Up as in Pixar
	Up
	// Down as in Down Under
	Down
)

// String shows the side benefit of using a newtype: we can implement
// the standard String() method that fmt.Println already understands
func (d Direction) String() string {
	switch d {
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Down:
		return "Down"
	}
	// Notice that we still have to handle the case of a non-existant enum
	// The compiler can't prove that we didn't call Direction(5) somewhere,
	// but at least we have to explicitely try to break things
	return "Unknown"
}

// icon is a private type that enables us to hide the enum constructor,
// ensuring that it is impossible for another module to accidentally construct
// an invalid enum.
type icon int

// Icon is public to allow external modules to store variables of the
// icon type. The method is private, preventing other modules from implementing
// the interface
type Icon interface {
	markIcon()
}

func (i icon) markIcon() {}

const (
	// Circle is a round shape
	Circle icon = iota
	// Square or Box
	Square
	// Arrow is pointy
	Arrow
)

// String is defined, again.
func (i icon) String() string {
	switch i {
	case Circle:
		return "()"
	case Square:
		return "[]"
	case Arrow:
		return "=>"
	default:
		// The compiler still can't prove away this default.
		// This results from having to mimic the pattern
		// instead of having it supported in the compiler
		return ""
	}
}
