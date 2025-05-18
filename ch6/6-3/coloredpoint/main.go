package main

import (
	"fmt"
	"image/color"
	"math"
)

// Point 结构体表示二维平面上的一个点
type Point struct {
	X, Y float64
}

// ColoredPoint 结构体嵌入了 Point 和 color.RGBA，表示一个带颜色的点
type ColoredPoint struct {
	*Point    // 匿名嵌入 Point，ColoredPoint 拥有了 Point 的字段和方法
	color.RGBA // 匿名嵌入 color.RGBA，ColoredPoint 拥有了 RGBA 的字段
}

// Distance 方法计算当前点到另一个点的距离
func (p Point) Distance(q Point) float64 {
	// 使用勾股定理计算距离
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

// ScaleBy 方法将当前点的坐标按比例缩放
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}



func main() {
	// 定义红色和蓝色
	red := color.RGBA{255, 0, 0, 255}   // 红色，不透明
	blue := color.RGBA{0, 0, 255, 255}  // 蓝色，不透明

	// 创建两个 ColoredPoint 实例，p 和 q
    //注意变量p和q的赋值方法，我们可以直接把Point{1, 1}赋值给p，因为ColoredPoint内嵌了Point
	var p = ColoredPoint{&Point{1, 1}, red}
	var q = ColoredPoint{&Point{5, 4}, blue}

	// 第一次计算 p 到 q 的距离
	fmt.Println(p.Distance(*q.Point)) // 输出 "5"

	p.Point=q.Point
	// 将 p 和 q 的坐标都放大 2 倍
	p.ScaleBy(2)
	q.ScaleBy(2)

	// 第二次计算 p 到 q 的距离
	fmt.Println(*p.Point,*q.Point)
	fmt.Println(p.Distance(*q.Point)) // 输出 "10"
}
