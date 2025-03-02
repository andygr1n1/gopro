package lib

import "fmt"

func FmtExamples() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "John", Age: 30}

	fmt.Println("Hello, World!")

	fmt.Printf("Hello, %v!\n", person)
	fmt.Printf("Hello, %#v!\n", person)
	fmt.Printf("Hello, %+v!\n", person)

}
