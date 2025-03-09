package lib

import "fmt"

func Pointers() {
	// Go has pointers. A pointer holds the memory address of a value.
	// The type *T is a pointer to a T value. Its zero value is nil.
	// declare a pointer to an integer (nil by default)
	var p *int
	// initialize an integer variable with value 42
	i := 42

	// The & operator generates a pointer to its operand.
	// assign the address of i to pointer p
	p = &i

	// print the memory address that p points to
	fmt.Println(p)
	// dereference p to print the value at that address (42)
	fmt.Println(*p)

	*p = 2 // set 2 through the pointer p
	// print the memory address (unchanged since p still points to i)
	fmt.Println(p)
	// print the value at that address (now 2)
	fmt.Println(*p)
	fmt.Println("i", i)

	// create a new int with zero value and assign its address to p
	p = new(int)
	// print the new memory address that p points to
	fmt.Println(p)
	// print the value at that address (0)
	fmt.Println(*p)

	// print the memory address again (same as above)
	fmt.Println(p)
	// print the value again (still 0)
	fmt.Println(*p)
}

// Указатели позволяют создавать новые объекты в памяти с помощью функции `new`.
// Это полезно для работы с динамическими структурами данных, такими как списки, деревья или графы.
