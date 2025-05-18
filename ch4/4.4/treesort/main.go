package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// 构建二叉树
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if t.value > value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func appendvalues(values []int, t *tree) []int {
	if t != nil {
		values = appendvalues(values, t.left)
		values = append(values, t.value)
		values = appendvalues(values, t.right)
	}
	return values
}
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendvalues(values[:0], root)
}
func main() {
	b := []int{1, 3, 4, 9, 8, 6, 7, 4, 98, 45, 67, 84, 56, 78}
	Sort(b)
	fmt.Println(b)
}