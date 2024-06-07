// déclaration du package pour ce fichier
package main

// importation du package fmt
import (
	"fmt"
)

func main() {
	// invocation de la fonction Println déclarée dans le package fmt
	fmt.Println("printList")
	node1 := Node[int]{
		value: 1,
	}
	node2 := Node[int]{
		value: 2,
		next:  &node1,
	}
	node3 := Node[int]{
		value: 3,
		next:  &node2,
	}
	node := Node[int]{
		value: 4,
		next:  &node3,
	}
	printList(node)
	fmt.Println("reverseList")
	printList(reverseList(node))
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func printList[T any](node Node[T]) {
	fmt.Println(node.value)
	if node.next != nil {
		printList(*node.next)
	}
}

func reverseList[T any](node Node[T]) Node[T] {
	var next *Node[T]
	var previous *Node[T]
	current := &node

	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return *previous
}
