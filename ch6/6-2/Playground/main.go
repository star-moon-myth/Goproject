package main

import "fmt"

type intlist struct {
	value int
	tail  *intlist
}

func (list *intlist) Sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.tail.Sum()
}
func main() {
	l := &intlist{
		value: 3,
		tail: &intlist{
			value: 4,
			tail: &intlist{
				value: 5,
				tail:  nil,
			},
		},
	}
	fmt.Println(l.Sum())
}