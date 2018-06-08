package main

import "fmt"

type Student3 struct {
	Name string
	Id   int
}

func (s *Student3) Update(id int) {
	s.Id = id
}

func main() {
	var f func(int)
	s := Student3{Name: "yinzhengjie"}
	f = s.Update
	f(200)
	fmt.Println(s) // 方法静态绑定,只能修改s这个学生

	var f1 func(s *Student3, id int)
	f1 = (*Student3).Update
	f1(&s, 300)
	fmt.Println(s) // 方法动态绑定,我们可以修改s这个学生。

	s1 := Student3{Name: "尹正杰"}
	f1(&s1, 400)    //同时也可以修改    s1这个学生。
	fmt.Println(s1) //动态绑定

}

