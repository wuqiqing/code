package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}
// 子结构
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//实现结构
func main() {
	var m Abser
	f:=MyFloat(-math.Sqrt2)
	m=f
	fmt.Print(m.Abs())
}
