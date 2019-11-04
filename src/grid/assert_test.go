package grid

import (
	"testing"
)

func TestThatAssertWillPanicWhenOutOfBounds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	g := *NewGrid(5, 10)

	assertRowBoundary(g, 5)
	assertColBoundary(g, 11)
}

func TestThatAssertWillNotPaninWhenBounds(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code DID panic")
		}
	}()

	g := *NewGrid(5, 10)

	assertRowBoundary(g, 0)
	assertRowBoundary(g, 1)
	assertRowBoundary(g, 2)
	assertRowBoundary(g, 3)
	assertRowBoundary(g, 4)

	assertColBoundary(g, 0)
	assertColBoundary(g, 1)
	assertColBoundary(g, 2)
	assertColBoundary(g, 3)
	assertColBoundary(g, 4)
	assertColBoundary(g, 5)
	assertColBoundary(g, 6)
	assertColBoundary(g, 7)
	assertColBoundary(g, 8)
	assertColBoundary(g, 9)
}
