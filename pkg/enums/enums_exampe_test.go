package enums_test

import (
	"fmt"

	"github.com/dacjames/patterns/pkg/enums"
)

// BasicEnums use integer types
func BasicEnums(i int) {
	fmt.Println(enums.Cardinal(i))
}

func ExampleBasicEnums() {
	BasicEnums(0)
	BasicEnums(1)
	BasicEnums(2)
	BasicEnums(3)
	BasicEnums(4)

	// Output: North
	// South
	// East
	// West
	// Unknown
}

func JustInts() {
	var u int = enums.North

	fmt.Println(u)
}

func ExampleJustInts() {
	JustInts()
	// Output: 0
}

type move struct {
	n         int
	direction enums.Direction
}

func EnumsTypeSafty(d enums.Direction) {
	m := move{
		n:         2,
		direction: d,
	}

	// var x int = 2
	// mbad := move{
	// 	n:         2,
	// 	direction: x,
	// }
	// Error: cannot use x (type int) as type enums.Direction in field value

	fmt.Println(m)

}

func ExampleEnumsTypeSafty() {
	// Direction does not have a String() method
	EnumsTypeSafty(enums.Left)
	EnumsTypeSafty(enums.Right)
	EnumsTypeSafty(enums.Up)
	EnumsTypeSafty(enums.Down)
	// Output: {2 0}
	// {2 1}
	// {2 2}
	// {2 3}
}

func EnumsLiteralsGotcha() {
	//
	mbad := move{
		n:         2,
		direction: 7,
	}
	// Gotcha: literals are untyped so 2 is inferred to have type Direction
	fmt.Println(mbad)
}

func ExampleEnumsLiteralsGotcha() {
	EnumsLiteralsGotcha()
	// Output: {2 7}
}

func EnumsStringer(s enums.Icon) {
	fmt.Println(s)
}

// ExampleEnumsStringer
func ExampleEnumsStringer() {
	EnumsStringer(enums.Square)
	EnumsStringer(enums.Circle)
	EnumsStringer(enums.Arrow)
	// Output: []
	// ()
	// =>
}

func InvalidEnum(e enums.Direction) {
	fmt.Println(e)
}

func ExampleInvalidEnum() {
	InvalidEnum(enums.Direction(7))
	// Output: Unknown
}

type Hack int

func (h Hack) markIcon() {
	panic("not implemented")
}

func PrivateEnums() {
	var i enums.Icon = enums.Circle

	fmt.Println(i)
	// fmt.Println(enums.icon(7))
	// cannot refer to unexported name

	// var h enums.Icon = Hack(7)

	// >>> cannot use Hack(7) (type Hack) as type enums.Icon in assignment:
	// Hack does not implement enums.Icon (missing enums.markIcon method)
	//     have markIcon()
	//     want enums.markIcon()

}
func ExamplePrivateEnums() {
	PrivateEnums()
	// Output: ()
}
