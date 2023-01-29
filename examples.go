package main

import (
	"jongmoss/datastruct/list/singlylinkedlist"
)

func main() {
	lst := singlylinkedlist.List[int]{}

	lst.Add(5)
	lst.Add(3)
	lst.Add(2)
	lst.Add(10)
	lst.Add(4)

}
