package main

import (
	"io"
	"os"
	"fmt"
)

type Byte_Counter struct {
	Sum int
}

func (b *Byte_Counter) Write(p []byte) (int, error) { //读取文件字节大小。
	b.Sum += len(p)
	return len(p), nil
}

type Line_Counter struct {
	Sum int
}

func (L *Line_Counter) Write(p []byte) (int, error) { //用于读取文件行号

	for _, j := range p { //循环p的内容
		if j == '\n' { //循环读取每一行，遇到换行符就自加“1”。
			L.Sum += 1 //由于对象“L”传的的是指针类型（*Line_Counter），换句话说，“L”是指针接受者，最终“L	”的参数会变动。
		}
	}
	return len(p), nil
}

func main() {
	lines := new(Line_Counter)
	bytes := new(Byte_Counter)
	w := io.MultiWriter(lines, bytes) //io模块的“MultiWriter”方法可以接受2个指针类型。将两个Writer（lines,bytes），合并成单个的Writer。类似于管道，但是他们是有区别的。

	io.Copy(w, os.Stdin) //将用户输入的数据传给w,最终交给lines和bytes指针去处理。
	fmt.Printf("该文本的行号是：\033[31;1m%d\033[0m 行\n", lines.Sum)
	fmt.Printf("该文本的字节大小是：\033[31;1m%d\033[0m 字节\n", bytes.Sum)
}
