package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

var (
	wg7 sync.WaitGroup
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	defer wg7.Done()
	for{
		select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			case <- time.After(3*time.Second):
					fmt.Println("		timeout		")
					close(c)
					return
	/*	default:
			fmt.Println("		stop		")*/
		}
	}
}
func main() {
	//runtime.GOMAXPROCS(1)
	wg7.Add(2)
	c := make(chan int, 3)
	quit := make(chan int)
	go fibonacci(c, quit)
	time.Sleep(time.Second)
	go func() {
		defer wg7.Done()
		for i := 0; i < 10; i++ {
			/*_,ok :=<-c
			if !ok {
				return
			}*/
			for j:=0;j<cap(c);j++{
				fmt.Println(<-c)
			}
			fmt.Println("————————"+strconv.Itoa(i)+"——————")
			time.Sleep(4*time.Second)
		}
		quit <- 0
	}()
	wg7.Wait()
	fmt.Println("————————End————————————")
}
