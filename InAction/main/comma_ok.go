package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List [] Element
type Person struct {
	name string
	age  int
}

//定义了String 方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func main() {
	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}
	for index, element := range list {
		if value, ok := element.(int); ok { //if 中初始化
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Println("list[%d] is of a different type", index)
		}
	}
	for index, element := range list{
		switch value := element.(type) {   //	element.(type)语法不能在switch 外的任何逻辑里面使用
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}



}
