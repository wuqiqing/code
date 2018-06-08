package main

import (
	"fmt"
	"strconv"
)

type Point2 struct {
	px int
	py int
}
func f2(m []int){
	m[3]=1
}

func  f1(p Point2) {
/*	c :=[]byte(*str)
	c[0] ='2'
	s := string(c)
	str = &s
	fmt.Println(*str)*/
	p.px=2
	p.py=3
}
func init(){
	fmt.Print("init")
}
func main() {
	/*	s := "hello"
		c := []byte(s) // 将字符串s 转换为[]byte 类型
		c[0] = 'c'
		fmt.Printf("%x\n",&s)
		s = string(c) // 再转换回string 类型
		s2 :=s+"_"*/
	//doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]string{5, 6, 7, 8}}

	p2 :=Point2{}
	//p2 := make(map[string]string)
	//p1 :=[6]int{}
 	//s :="111"
	//p2 :=p1[1:]
	fmt.Println(p2)
	f1(p2)
	fmt.Println(p2)
	intval :=make([]int,3)
	fmt.Print(strconv.Itoa(intval[0]))
}
