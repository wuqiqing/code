package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	/*
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
	*/

	// buffered
	c2 := make(chan int,2)//修改2为1就报错，修改2为3可以正常运行
	c2 <- 1
	c2 <- 2
	c2 <- 3
	x, y := <-c2, <-c2 // receive from c
	z :=<- c2
	fmt.Println(x, y, x+y,z)
	for i := range c2 {
		fmt.Println(i)
	}

}
