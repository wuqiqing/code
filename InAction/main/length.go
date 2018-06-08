package main

import "fmt"

var PI float64 = 3.1415926

type square struct {
	//定义正方向边长X
	X float64
}
type circle struct {
	//定义圆形的半径
	r float64
}
type Area_Perimeter interface {
	//定义求面积和边长的接口
	area() float64
	perimeter() float64
}

func (s *square) area() float64 { //给正方向绑定求面积方法
	return s.X * s.X
}
func (c *circle) area() float64 { //给圆形绑定求面积的方法
	return PI * c.r * c.r
}

func (s square) perimeter() float64 {
	square_perimeter := s.X * 4
	return square_perimeter
}
func (c circle) perimeter() float64 {
	circle_perimeter := 2 * PI * c.r
	return circle_perimeter
}

func main() {
	var s, c Area_Perimeter
	s = &square{10} //通过接口给“square”结构体传值。
	c = &circle{20}
	fmt.Printf("正方形的面积是：\033[31;1m%v\033[0m，正方形的周长是：\033[31;1m%v\033[0m\n", s.area(), s.perimeter())
	fmt.Printf("圆形的面积是：\033[31;1m%v\033[0m，圆形的周长是：\033[31;1m%v\033[0m\n", c.area(), c.perimeter())
}
