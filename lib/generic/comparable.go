package generic

import "fmt"

func Comparable() {
	si := []int{10, 20, 15, -10}

	ss := []string{"Andrei", "Dashka"}

	// The problem is that fmt.Printf expects a format specifier when printing a slice. You should use %v to print the slice
	// fmt.Printf(si)
	fmt.Printf("%d\n", si)
	fmt.Printf("%v\n", si)
	fmt.Printf("%v\n", si)

	fmt.Println(Index(si, -10))
	fmt.Println(Index(ss, "Andrei"))

}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}
