package main

import "fmt"

type defaultStu struct {
   x int
   y int
}
func main() {

	fmt.Println(&defaultStu{})
}
