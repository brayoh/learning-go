package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs method has a value receiver, so it cannot modify the value that it points to.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale method has a pointer receiver, so it can modify the value that it points to.
// This is different from the Abs method, which has a value receiver and cannot modify the value that it points to.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func receiver_functions() float64 {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	return v.Abs()
}