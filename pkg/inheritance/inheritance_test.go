package inheritance

import "testing"

func TestDescribeShape(t *testing.T) {
	c := NewCircle(1)
	r := NewRectangle(2, 2)

	if res := c.Describe(); res != "Shape: Circle, Area: 3" {
		t.Logf("circle.Describe() == %s", res)
		t.Fail()
	}

	if res := r.Describe(); res != "Shape: Rectangle, Area: 4" {
		t.Logf("rectangle.Describe() == %s", res)
		t.Fail()
	}
}

func TestBadConstructor(t *testing.T) {
	sq := BadNewSquare(2)

	defer func() {
		if r := recover(); r != nil {
			t.Log("Panic on method resolution is expected")
		} else {
			t.Fatal("Did not panic on method resolution!")
		}
	}()

	t.Log(sq.Describe())
}

func TestMissingChildMethod(t *testing.T) {
	tr := NewTriangle(1, 2)

	if res := tr.Name(); res != "Unnamed" {
		t.Logf("tr.Name() == %s", res)
		t.Fail()
	}

	// no way to catch stack overflow, uncomment this to see failure
	// tr.Describe()
}
