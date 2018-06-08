package main

import (
	"fmt"
	"reflect"
)

type boy struct {
	Name string
	age  int
}

type human interface {
	SayName()
	SayAge()
}

func (this *boy) SayName() {
	fmt.Println(this.Name)
}

func (this *boy) SayAge() {
	fmt.Println(this.age)
}

func main() {
	// 定义接口变量
	var i human
	// 初始化对象，jown持有对象指针。
	jown := &boy{
		Name: "jown",
		age:  15,
	}
	// 因为boy实现了human中的方法，所以它实现了human接口。
	// 这时，i就指向jown对象。
	i = jown

	// 通过反射获取接口i 的类型和所持有的值。
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	// ... 后续操作
	fmt.Println(t,"\n",v)
	method, _ := t.MethodByName("SayAge")
	fmt.Println(method)
	//fmt.Println(t.Elem().Field(1))
/*	jown2 := boy{
		Name: "jown2",
		age:  10,
	}*/
	jown3 :=jown
	jown3.age=9
	v.Elem().Set(reflect.ValueOf(*jown3))
	fmt.Println(v.Elem().FieldByName("age"))



















}
