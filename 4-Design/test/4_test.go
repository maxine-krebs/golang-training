package test

import (
	"fmt"
	"testing"
)

type X interface {
	X() int
}

type Y interface {
	Y(int)
}

type Z interface {
	X(int) int
	Z(string) (string, error)
}

type A interface {
	A() A
}

type x struct{}
type y struct{}
type z struct{}
type a struct{}

func IsX(v X) bool {
	return true
}

func IsY(v Y) bool {
	return true
}

func IsZ(v Z) bool {
	return true
}

func IsA(v A) bool {
	return true
}

// Begin implementation, first case done for you
func (v x) X() int {
	return 1
}

// Modify below this line

func (v y) Y(i int) {
	fmt.Println(i)
}

func (v z) X(i int) int {
	return i
}

func (v z) Z(s string) (string, error) {
	return s, nil
}

func (v a) A() A {
	return v
}

// Modify above this line

func TestInterfaces(t *testing.T) {
	// Do not modify this function
	var xx x
	var yy y
	var zz z
	var aa a
	if !IsX(xx) {
		t.Error("xx is not an X")
	}
	if !IsY(yy) {
		t.Error("yy is not an Y")
	}
	if !IsZ(zz) {
		t.Error("zz is not an Z")
	}
	if !IsA(aa) {
		t.Error("aa is not an A")
	}
}
