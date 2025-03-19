package generic

import "fmt"

type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) append(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
}

func (l List[T]) print() {
	current := &l
	
	for current != nil {
		fmt.Printf("%v -> ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *List[T]) length() int {
	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func GenericTypes() {
	intSnake := &List[int]{val: 1}

	intSnake.append(2)
	intSnake.append(3)
	intSnake.append(4)

	fmt.Print("Linked List: ")
	intSnake.print()

	fmt.Printf("Length of Linked List: %d\n", intSnake.length())

	stringSnake := &List[string]{val: "Hello"}
	stringSnake.append("World")
	stringSnake.append("Go")
	stringSnake.append("Generics")

	fmt.Print("\nString Linked List: ")
	stringSnake.print()
	fmt.Printf("Length of String Linked List: %d\n", stringSnake.length())
}
