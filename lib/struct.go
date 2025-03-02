// A struct is a collection of fields.
package lib

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) Sum() int {
	return v.X + v.Y
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Struct() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)

	v.Scale(10)

	fmt.Println("sum:", v.Sum())
}
