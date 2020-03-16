// Package markers demonstrates the use of marker traits to simulate tagged unions.
// For information on the marker trait pattern in general, see markers/README.md.
package markers

import "strings"

// Expr is the marker trait that simulated a tagged union.
// It is analagous to a sealed trait / abstract class in scala.
// This example creates a simple expression language that supports
// integers and addition, demonstrating how one might solve the problem
// of evaluating user-supplied input. This problem arises classically
// in compilers but also arises in more common applications when handling
// query parameters.
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

// Query demonstrates a more practical use for marker traits:
// transforming user queries into SQL requests.
type Query interface {
	markQuery()
}

// OrderByQ represents an ORDER BY query.
type OrderByQ struct {
	field string
}

func (OrderByQ) markQuery() {}

// SortByQ represents a SORT BY query.
type SortByQ struct {
	field string
}

func (SortByQ) markQuery() {}

// BuildSQL demonstrates how to use queries to use Query values.
// This example is very minimal and would need to be extended to
// support user-supplied parameters to be practically useful.
// Usage is practically the same as the expression example:
// just type switch on the interface type and work with the
// matched struct types.
func BuildSQL(sb *strings.Builder, queries ...Query) {
	for _, query := range queries {
		switch q := query.(type) {
		case OrderByQ:
			if !validateField(q.field) {
				continue
			}
			sb.WriteString(" ORDER BY ")
			sb.WriteString(q.field)
			sb.WriteString(" ")
		case SortByQ:
			if !validateField(q.field) {
				continue
			}
			sb.WriteString(" SORT BY ")
			sb.WriteString(q.field)
			sb.WriteString(" ")
		}
	}
}

// validateField should be replaced with some mechanism
// to prevent SQL injection if the query field is user-specified.
// The implemenation here is purposefully dumb but is included
// to demonstrate that some kind of validation or sanitation is
// required when building SQL from user inputs.
func validateField(field string) bool {
	if field == "username" {
		return true
	}
	if field == "email" {
		return true
	}
	return false
}

// ExampleSQL shows how a very simple SQL query can be constructed
// very efficiently using a stack-allocated string builder and query
// structs wrapped in marker traits.ß
func ExampleSQL() string {
	sb := strings.Builder{}

	sb.WriteString(" SELECT username, email ")
	sb.WriteString(" FROM users ")

	BuildSQL(&sb, SortByQ{field: "username"}, OrderByQ{field: "email"})

	return sb.String()
}
