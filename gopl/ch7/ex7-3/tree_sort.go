package ch7

import (
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// String 从小到大排序，返回string
func (t *tree) String() string {
	defaultN := 10
	var bf strings.Builder
	values := make([]int, 0, defaultN)
	values = appendValues(values, t)
	for _, v := range values {
		bf.WriteString(strconv.Itoa(v))
	}
	return bf.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root.String())
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
