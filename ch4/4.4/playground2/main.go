package main

import "fmt"

type Point struct {
	X int
	Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	a := Wheel{Circle{Point{2, 3}, 5}, 3}
	fmt.Printf("%#v\n",a)
	b:=Wheel{
		Circle: Circle{
			Point: Point{2,3},
			Radius: 5,
		},
		Spokes: 3,
	}
	fmt.Printf("%#v\n",b)
}
