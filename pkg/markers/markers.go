// Package markers demonstrates the use of marker traits to simulate tagged unions.
// For information on the marker trait pattern in general, see markers/README.md.
package markers

// Expr is the marker trait that simulated a tagged union.
// It is analagous to a sealed trait / abstract class in scala.
// This example creates a simple expression language that supports
// integers and addition, demonstrating how one might solve the problem
// of evaluating user-supplied input. This problem arises classically
// in compilers but also arises when handling query parameters in more
// common applications.
type Expr interface {
	// markExpr is a private method because we only want to permit
	// implementators in the markers package. While the compiler cannot
	// check for exhuastiveness (see Eval()), using private
	// still helps maintainers keep track of Expr "subtypes".
	markExpr()
}

// IntE is a newtype for int that allows us to mark integers as expressions.ß
type IntE int

func (IntE) markExpr() {}

// AddE is a struct type that represents adding two arguments, called lhs and rhs
// shorthand for left and right hand side. The arguments are themselves expressions.
type AddE struct {
	lhs Expr
	rhs Expr
}

func (AddE) markExpr() {}

// Eval demonstrates how to use a marker to evaluate the expression language we have
// implemented. Notice that we have to handle the case where Expr is neither an IntE
// or AddE even though we know that cannot occur.
func Eval(expr Expr) int {
	switch e := expr.(type) {
	case IntE:
		return int(e)
	case AddE:
		lhs := Eval(e.lhs)
		rhs := Eval(e.rhs)
		return lhs + rhs
	default:
		// A smarter (but slower) compiler could prove that this case cannot be reached.
		panic("unreachable")
	}
}

// ExampleProgram shows how we might use the our expression langugage to write a trivial
// program. If this kind of an expression tree was used in a compiler, it is often called
// an abstract syntax tree (AST). The program calculates the expression ((3 + (5 + 0)) + (1 + 2))
// but unlike writing that expression in Go, our program is a plain data structure that we could
// create and manipulate in some way other than evaluating it. This concept, called `reification`
// has nothing to do with marker traits, but it's pretty cool, right?
func ExampleProgram() Expr {
	return AddE{
		AddE{IntE(3), AddE{IntE(5), IntE(0)}},
		AddE{IntE(1), IntE(2)},
	}
}

// Add1 is here to show the fun you can have with reification.
func Add1(in Expr) Expr {
	return AddE{in, IntE(1)}
}
