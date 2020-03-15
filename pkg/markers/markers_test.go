package markers

import "testing"

// TestMarkersEval tests evaluating several instances of our expression language.
func TestMarkersEval(t *testing.T) {
	// Aside: the if true {} block is used to make := vs = usage consistent
	// Aside: this could be abstracted in several ways, but I prefer to make
	// tests as unabstracted as possible.
	if true {
		prog := IntE(1)
		expected := 1
		if res := Eval(prog); res != expected {
			t.Logf("Eval(%#v) => %d != %d", prog, res, expected)
			t.Fail()
		}
	}
	if true {
		prog := AddE{IntE(2), IntE(-1)}
		expected := 1
		if res := Eval(prog); res != expected {
			t.Logf("Eval(%#v) => %d != %d", prog, res, expected)
			t.Fail()
		}
	}
	if true {
		prog := ExampleProgram()
		expected := 11
		if res := Eval(prog); res != expected {
			t.Logf("Eval(%#v) => %d != %d", prog, res, expected)
			t.Fail()
		}
	}
	if true {
		prog := Add1(ExampleProgram())
		expected := 12
		if res := Eval(prog); res != expected {
			t.Logf("Eval(%#v) => %d != %d", prog, res, expected)
			t.Fail()
		}
	}
}
