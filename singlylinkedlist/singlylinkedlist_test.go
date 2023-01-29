package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given a linked list 3 -> 1 -> 2 -> 1
// When we remove the first occurrence of node with value 1
// Then final list is 3 -> 2 -> 1
func TestListRemoveCheckFirstOccurrence(t *testing.T) {
	lst := List[int]{}
	lst.Add(1)
	lst.Add(2)
	lst.Add(1)
	lst.Add(3)
	assert.Equal(t, 4, lst.Size)
	assert.True(t, lst.Remove(1))
	assert.Equal(t, 3, lst.Size)
	assert.Equal(t, []int{3, 2, 1}, lst.ToArray())
}

// Given an initialized list
// Then the head of the list should be nil with size 0
func TestList(t *testing.T) {
	lst := List[int]{}
	assert.Nil(t, lst.Head)
	assert.Equal(t, 0, lst.Size)
}

func TestListAdd(t *testing.T) {
	lst := List[int]{}
	elemOne := lst.Add(1)
	elemTwo := lst.Add(2)
	assert.Equal(t, 2, lst.Size)
	assert.Equal(t, 1, elemOne.Val)
	assert.Equal(t, elemTwo, lst.Head)
}

func TestListRemove(t *testing.T) {
	lst := List[int]{}
	lst.Add(1)
	lst.Add(2)
	assert.Equal(t, 2, lst.Size)
	assert.True(t, lst.Remove(1))
	assert.Equal(t, 1, lst.Size)
}

func TestToArray(t *testing.T) {
	lst := List[int]{}
	lst.Add(1)
	lst.Add(2)
	lst.Add(3)
	assert.Equal(t, []int{3, 2, 1}, lst.ToArray())
}
