// A struct is a collection of fields.
package lib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Vertex struct {
	X int
	Y int
}

type VertexRequired struct {
	X int `validate:"required"` // Field X is marked as required
	Y int `validate:"required"` // Field Y is marked as required
}

func (v Vertex) Sum() int {
	return v.X + v.Y
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Struct() {
	v := Vertex{X: 1}
	fmt.Println(v)
	fmt.Printf("v: %v\n", v)
	fmt.Printf("Hello, %v\n", v)
	fmt.Printf("Hello, %#v\n", v)
	fmt.Printf("Hello, %+v\n", v)

	v.Scale(10)

	fmt.Println("sum:", v.Sum())

	vr := VertexRequired{X: 1}
	validate := validator.New()
	err := validate.Struct(vr)
	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}
	fmt.Println("Validation succeeded:", vr)
}
