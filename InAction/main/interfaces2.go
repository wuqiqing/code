package main

import (
	"fmt"
)

type I4 interface {
	IM1() error
}
type IS1 struct {
	f string
}
func (s *IS1) IM1() error {
	fmt.Println("———————— S"+s.f+"M1 ——————————")
	s.f="2"
	return nil
}
func main() {
	s:=IS1{"1"}
	s.IM1()
	fmt.Printf(s.f)
}


