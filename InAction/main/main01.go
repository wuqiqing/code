// 这个示例程序展示如何创建goroutine
 // 以及调度器的行为
package main

import (
	"runtime"
	"sync"
	"fmt"
)

 // main 是所有Go 程序的入口
 func main() {
	 // 分配一个逻辑处理器给调度器使用
	 runtime.GOMAXPROCS(2)
	
	 // wg 用来等待程序完成
	 // 计数加 ，表示要等待两个goroutine
	 var wg sync.WaitGroup
	 wg.Add(2)
	
	 fmt.Println("Start Goroutines")
	 // 声明一个匿名函数，并创建一个goroutine
	 go func() {
		 // 在函数退出时调用Done 来通知main 函数工作已经完成
		defer wg.Done()

		 // 显示字母表 次
		 for count :=0 ; count < 5; count++ {
			 for char := 'A'; char < 'A'+26; char++ {
				 fmt.Printf("%c ", char)
			 }
			 fmt.Print("\n")
		 }

	 }()
	 //wg.Done()
	 fmt.Print("\n")
	 fmt.Print("\n")

	 // 声明一个匿名函数，并创建一个goroutine
	 go func() {
		 // 在函数退出时调用Done 来通知main 函数工作已经完成
		 defer wg.Done()
		
		 // 显示字母表 次
		 for count := 0; count < 5; count++ {
			 for char := 'a'; char < 'a'+26; char++ {
				 fmt.Printf("%c ", char)
				 }
			 fmt.Print("\n")
			 }

		 }()



	 // 等待 goroutine 结束
	 fmt.Println("Waiting To Finish")
	 wg.Wait()
	
	 fmt.Println("\nTerminating Program")
	 }