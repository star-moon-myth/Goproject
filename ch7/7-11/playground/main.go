package main

import "fmt"

// 定义一个接口 Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 定义正方形
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// 处理形状的函数
func processShape(s Shape) { // 参数类型是 Shape 接口
	switch actualShape := s.(type) { // 类型分支！探测 s 实际类型
	case Circle: // 如果是圆形
		fmt.Println("这是一个圆形，半径是:", actualShape.Radius)
		fmt.Println("面积:", actualShape.Area())
	case Square: // 如果是正方形
		fmt.Println("这是一个正方形，边长是:", actualShape.Side)
		fmt.Println("周长:", actualShape.Perimeter())
	default: // 如果是其他类型 (Shape 接口的其他实现)
		fmt.Println("我不知道这是什么形状，但是它是一个 Shape 接口!")
	}
}

func main() {
	c := Circle{Radius: 5}
	sq := Square{Side: 10}

	processShape(c)  // 传入圆形
	processShape(sq) // 传入正方形
	// processShape("hello") // ❌ 报错！ "hello" 不是 Shape 接口的实现！
}
