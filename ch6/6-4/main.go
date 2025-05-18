package main

import "fmt"

type Pointer struct {
	X, Y float32
}

func (p Pointer) add(q Pointer) Pointer {
	return Pointer{p.X + q.X, p.Y + q.Y}
}
func (p Pointer) sub(q Pointer) Pointer {
	return Pointer{p.X - q.X, p.Y - q.Y}
}

type Path []Pointer //只是声明而已
func (p Path) tranby(offset Pointer, is bool) {
	var op func(p, q Pointer) Pointer
	if is {
		op = Pointer.add
	} else {
		op = Pointer.sub
	}
	for i := range p {
		p[i] = op(p[i], offset)
	}
}
func main() {
	p := Pointer{1, 3}
	path := Path{{1, 2}, {2, 3}}
	path.tranby(p, true)
	fmt.Println(path)
	path.tranby(p,false)
	fmt.Println(path)

}