package main

import (
	"fmt"
	"math"
)
type path []pointer
type pointer struct {
	x, y float64
}
func (p path)distance()float64{
	sum:=0.0
	for i:=range p{
		if i>0{
			sum += p[i-1].distance(p[i])
		}
	}
	return sum
}
func distance(m, n pointer) float64 {
	return math.Hypot(n.x-m.x,m.y-n.y)
}
func(m pointer)distance(n pointer)float64{
	return math.Hypot(n.x-m.x,m.y-n.y)
}
func main(){
	m:=pointer{1,2}
	n:=pointer{4,6}
	fmt.Println(distance(m,n))
	fmt.Println(m.distance(n))
	perim:=path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}
	fmt.Println(perim)
	fmt.Println(perim.distance())
}