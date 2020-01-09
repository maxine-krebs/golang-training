package test

import (
	"fmt"
	"testing"
)

func TestVariables(*testing.T) {
	var a int
	fmt.Println(a)
	// Change below this line
	a = 45
}

func TestScope(t *testing.T) {
	var a int = 2
	var b int = 2

	// Change above this line
	if a != b {
		t.Error("a and b should be equal.")
	}
}

func TestAnonymous(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	// Change above this line
	ten := add(-10, 20)

	if ten != 10 {
		t.Error("This should equal 10.")
	}
}

func TestControl(t *testing.T) {
	a := 11
	b := 12
	c := 13

	// Change above this line only

	if a+1 != b || b+1 != c {
		t.Error("a + 2 = b + 1 = c")
		return
	}

	var x, y int
	for i := 0; i < 2; i++ {
		switch {
		case a < b && a > 10:
			a = b * 2
			x = 1
		case a < 25 && c < 14 && c > b:
			y = 1
		}
	}

	if x != 1 || y != 1 {
		t.Error("Lost control.")
	}
}
