package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"os"
	"bufio"
	"strings"
)

type Student struct {
	//定义一个名称为“Student”的结构统；
	ID   int    //定义学员编号
	NAME string //定义姓名
}

type ClassRoom struct {
	//定义一个名为“ClassRoom”的结构体；
	teacher  string              //定义一个名称为“teacher”字符串类型的变量，
	students map[string]*Student //定义一个变量名为“students”其类型为map的变量。
}

var classrooms map[string]*ClassRoom //声明一个名为“classrooms”的变量，指定其类型是map，要注意的是map中的value是指针类型的结构统哟。

var student_num = make(map[int]Student) //定义存取学生成员的函数，注意这里是要初始化的，字典在使用前必须初始化，不然会报错！

func (s *ClassRoom) Delete() {

}

func (c *ClassRoom) Add(args []string) error { //自定义添加学生信息的函数
	if len(args) != 2 {
		fmt.Println("您输入的字符串有问题,案例：add 01 bingan")
		return nil
	}
	id := args[0]
	student_name := args[1]
	student_id, _ := strconv.Atoi(id)

	for _, s := range student_num {
		if s.ID == student_id {
			fmt.Println("您输入的ID已经存在，请重新输入")
			return nil
		}
	}
	student_num[len(student_num)+1] = Student{student_id, student_name}
	fmt.Println("Add successfully!!!")
	return nil
}

func (c *ClassRoom) Drop(args []string) error {
	if len(args) != 1 {
		fmt.Println("你愁啥？改学生ID压根就不存在！")
		return nil
	}
	id := args[0]
	student_id, _ := strconv.Atoi(id)
	for i, j := range student_num {
		if j.ID == student_id {
			delete(student_num, i) //删除map中该id所对应的key值。但是该功能需要完善！
			fmt.Println("delete successfully!")
			return nil
		}
	}
	fmt.Println("你愁啥？学生ID压根就不存在！")
	return nil
}

func (c *ClassRoom) Update(args []string) error { //定义修改的函数
	if len(args) != 2 {
		fmt.Println("您输入的字符串有问题,案例：add 01 bingan")
		return nil
	}
	id := args[0]                     //取出ID
	student_name := args[1]           //取出姓名
	student_id, _ := strconv.Atoi(id) //将字符串ID变成数字类型。
	for i, j := range student_num {
		if j.ID == student_id {
			student_num[i] = Student{student_id, student_name} //这其实就是一个赋值的过程。
			fmt.Println("update successfully!")
			return nil
		}
	}
	fmt.Println("你愁啥？学生ID压根就不存在！")
	return nil
}

func (c *ClassRoom) List(args []string) error { //给“ClassRoom”绑定一个“List”方法。
	if len(student_num) == 0 {
		fmt.Println("数据库为空，请自行添加相关信息！")
		return nil
	}
	for _, value := range student_num {
		fmt.Printf("学员的姓名是：\033[31;1m%s\033[0m，学员编号是：\033[31;1m%d\033[0m\n", value.NAME, value.ID)
	}
	return nil
}

func (c *ClassRoom) Save(args []string) error { //定义存取的函数
	if len(args) == 0 {
		fmt.Println("请输入您想要保存的文件名，例如：save student.txt")
		return nil
	}
	file_name := args[0]
	f, err := json.Marshal(student_num) //把变量持久化，也就是将内存的变量存到硬盘的时进行的序列化的过程

	if err != nil {
		fmt.Println("序列化出错啦！")
	}
	ioutil.WriteFile(file_name, f, 0644) //将数据写入硬盘，并制定文件的权限。
	fmt.Println("写入成功")
	return nil
}

func (c *ClassRoom) Load(args []string) error { //定义加载的函数。
	if len(args) != 1 {
		fmt.Println("输入错误，请重新输入.")
		return nil
	}
	file_name := args[0]
	s, _ := ioutil.ReadFile(file_name)
	json.Unmarshal(s, &student_num)
	fmt.Println("读取成功！")
	return nil
}

func (c *ClassRoom) Exit(args []string) error { //定义对出的脚本
	os.Exit(0) //里面的数字表示用户结束程序的返回值，返回0意味着程序是正常结束的。
	return nil
}

func main() {
	classrooms = make(map[string]*ClassRoom) /*初始化字典，因为上面只定义没有初始化。初始化赋值这里不能加":="，
    因为作用域不同（会将全局作用域的值给覆盖掉），加了得到的结果返回："null".*/
	fmt.Println("学生管理系统迷你版！")
	f := bufio.NewReader(os.Stdin) //用它读取用户输入的内容
	for {
		fmt.Print("请选择您要去的教室\n")
		fmt.Print("请输入:>>>")
		line, _ := f.ReadString('\n')   //将读取的内容按照"\n"换行符来切分，注意里面是单引号哟！
		line = strings.Trim(line, "\n") //表示只脱去换行符："\n",你可以自定义脱去字符，等效于line = strings.TrimSpace(line)
		content := strings.Fields(line) //按照空格将得来的字符串做成一个切片。

		if len(content) == 0 { //脱去空格
			continue
		}
		if len(content) == 1 {
			fmt.Println("您输入的字符串有问题,案例：select yinzhengjie!")
			continue
		}
		ClassRoom_Chose := content[0] //定义执行命令的参数，如add,upadte,list,delete....等等
		Classroom_len := content[1:]
		Classroom := Classroom_len[0] //这个就是讲字符串切片转换成字符串。
		if len(Classroom_len) == 1 {
			classrooms[Classroom] = &ClassRoom{
				students: make(map[string]*Student),
			}
			for k,v :=range classrooms {
				fmt.Printf("%s=%d;", k, v)
			}
			fmt.Print(classrooms)
		} else {
			fmt.Println("您输入的参数有问题！")
		}
		if ClassRoom_Chose == "select" || ClassRoom_Chose == "SELECT" {
			fmt.Printf("               欢迎来到\033[31;1m%s\033[0m教室\n", Classroom)
			for {
				fmt.Print("请输入:>>>")
				line, _ := f.ReadString('\n')   //将读取的内容按照"\n"换行符来切分，注意里面是单引号哟！
				line = strings.Trim(line, "\n") //表示只脱去换行符："\n",你可以自定义脱去字符，等效于line = strings.TrimSpace(line)
				content := strings.Fields(line) //按照空格将得来的字符串做成一个切片。
				if len(content) == 0 { //脱去空格
					continue
				}
				cmd := content[0]   // 定义执行命令的参数，如add,upadte,list,delete....等等
				args := content[1:] // 定义要执行的具体内容
				actiondict := map[string]func([]string) error{//定义用户的输入内容
					"add": classrooms[Classroom].Add, //表示用户输入的是字符串"add"时，其要执行的结构体的方法是classrooms[Classroom].Add，也就是“Add”方法，以下定义同理。
					"list": classrooms[Classroom].List,
					"update": classrooms[Classroom].Update,
					"delete": classrooms[Classroom].Drop,
					"save": classrooms[Classroom].Save,
					"load": classrooms[Classroom].Load,
					"exit": classrooms[Classroom].Exit,
				}
				action_func := actiondict[cmd] //定义用户执行的函数
				if action_func == nil { //如果输入有问题，告知用户用法
					fmt.Println("Usage: {add|list|where|load|upadte|delete|}[int][string]")
					continue
				}
				err := action_func(args)
				if err != nil {
					fmt.Println("您输入的字符串有问题,案例：add 01 yinzhengjie")
					continue
				}
				continue
			}
		}
	}
}
