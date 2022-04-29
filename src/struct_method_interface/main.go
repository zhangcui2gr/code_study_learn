package struct_method_interface

import "math"

// 矩形结构体，包含高度和宽度
type Rectangle struct {
	Height float64
	Weight float64
}

// 圆形结构体，包含半径
type Circle struct {
	Radius float64
}

// 计算矩形面积的方法
func (rectangle Rectangle) Area() float64 {
	return rectangle.Height * rectangle.Weight
}

// 计算圆形面积的方法
func (circle Circle) Area() float64 {
	return 3.14 * math.Pow(circle.Radius, 2)
}

//接口实现多态,通过传入不同的结构体，调用响应的方法
type Shape interface {
	Area() float64
}
