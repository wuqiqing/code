package main

import (
	"fmt"
	"strconv"
)

type I1 interface {
	M1() error
	M3() error
}
type I2 interface {
	M2() error
	//M1() error
}

type I3 interface {
	I1
	I2
	//M3()
}

type S1 struct{f int}
type S2 struct{f int}
type S3 struct {
	I1
	I2
}

func (s S1) M1() error {
	fmt.Println("———————— S"+strconv.Itoa(s.f)+"M1 ——————————")
	s.f=2
	return nil
}
func (s S1) M3() error {
	fmt.Println("———————— S"+strconv.Itoa(s.f)+"M1 ——————————")
	return nil
}
func (S2) M2() error {
	fmt.Println("———————— S2M2 ——————————")
	return nil
}
func main() {
	s3:=S3{
		I1:S1{1},
		I2:S2{2},

	}
	exec(&s3)
}

func exec(i3 I3) {
	i3.M1()
}

