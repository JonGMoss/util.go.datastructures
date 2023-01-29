package singlylinkedlist

import "fmt"

// Base struct with a single value head which is an element with generic type T.
type List[T comparable] struct {
	Head *element[T]
	Size int
}

// Define an element type, containing the value of the element with type T and the next element
type element[T comparable] struct {
	Val  T
	Next *element[T]
}

// Adds an element of value 'val' to the end of the list
// returns the element that was created as a pointer
func (lst *List[T]) Add(val T) *element[T] {
	elem := &element[T]{Val: val}
	if lst.Head == nil {
		lst.Head = elem
		lst.Size++
		return elem
	}
	elem.Next = lst.Head
	lst.Head = elem
	lst.Size++
	return elem

}

// Removes the first occurrence of element 'val' from the list.
// If it doesn't exist return false
func (lst *List[T]) Remove(val T) bool {
	if lst.Head == nil {
		return false
	}
	currentNode := lst.Head

	if currentNode.Val == val {
		lst.Head = currentNode.Next
		lst.Size--
		return true
	}

	var prev *element[T] = nil

	// while loop to check if we have got to the last node
	for currentNode != nil && currentNode.Val != val {
		prev = currentNode
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return false
	}
	prev.Next = currentNode.Next
	lst.Size--
	return true
}

// Clears the entire list of elements
func (lst *List[T]) Clear() bool {
	// If the Head of the list is nil we do not need to do anything
	if lst.Head == nil {
		lst.Size = 0
		return true
	}
	// Otherwise, we should set the Head to nil and the size to 0
	// Create an empty list
	var tmp List[T]
	lst.Head = tmp.Head
	lst.Size = 0

	return true
}

// Convert the linked list to an array
// Simple iterate over each element and append to a slice
func (lst *List[T]) ToArray() []T {

	r := make([]T, lst.Size)
	currentNode := lst.Head
	i := 0
	for currentNode != nil {
		// r = append(r, currentNode.Val)
		r[i] = currentNode.Val
		currentNode = currentNode.Next
		i++
	}
	return r
}

// Pretty printing the linked list
func (lst *List[T]) Pretty() {
	array := lst.ToArray()
	fmt.Printf("%+v\n", array)
}
