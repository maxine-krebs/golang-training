package test

import (
	"testing"
)

func TestPointers(t *testing.T) {
	type dog struct {
		name string
		age  int
	}

	setAge := func(d *dog) {
		d.age = 5
	}

	d := &dog{}
	setAge(d)
	if d.age != 5 {
		t.Error("Dog is not 5")
	}
}

func TestIteration(t *testing.T) {
	x := []int{4, 5, 6, 7}

	sum := 0
	for _, el := range x {
		sum = sum + el
	}

	if sum != 4+5+6+7 {
		t.Error("Wrong sum")
	}
}
