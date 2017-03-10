package main

import (
	"fmt"
)

type tree struct {
	value int
	left, right *tree
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func addLeaf(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = addLeaf(t.left, value)
	} else {
		t.right = addLeaf(t.right, value)
	}

	return t
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = addLeaf(root, v)
	}
	appendValues(values[:0], root)
}

func main() {

	var array = []int{5, 7, 4, 12, 11}

	Sort(array)

	fmt.Println(array)


}