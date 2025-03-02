package lib

import "fmt"

func Defer() {
	x := 10
	defer fmt.Println(x) // x = 10 at this point
	x = 20
	fmt.Println(x) // Prints 20
}
