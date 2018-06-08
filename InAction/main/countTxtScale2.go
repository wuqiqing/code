package main

import (
	"io"
	"os"
	"fmt"
)

type Byte_Counter3 int

func (b *Byte_Counter3) Write(p []byte) (int, error) {
	*b += Byte_Counter3(len(p))
	return len(p), nil
}


type Byte_Counter4 struct {  //定义一个统计字节变量的类。
	Sum int
}

func (b *Byte_Counter4) Write(p []byte)(int, error) { //读取文件字节大小。
	b.Sum += len(p)
	return len(p),nil
}


func main() {
	T := new(Byte_Counter4) //此时的b其实是指针
	io.Copy(T, os.Stdin)
	fmt.Printf("文件的字节大小为：\033[31;1m%v\033[0m\n", *T)
}
