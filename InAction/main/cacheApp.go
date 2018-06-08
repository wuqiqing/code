package main

import (
	"io"
	"fmt"
	"bytes"
)

type Byte_Counter2 struct {
	Sum int
}

func (b *Byte_Counter2) Write(p []byte) (int, error) { //读取文件字节大小。
	b.Sum += len(p)
	return len(p), nil
}

type Line_Center struct {
	Sum int
}

func (b *Line_Center) Write(p []byte) (int, error) { //读取文件行号.
	b.Sum = 1
	for _, j := range p {
		if j == '\n' {
			b.Sum ++
		}
	}
	return len(p), nil
}

func main() {
	l := new(Line_Center)
	b := new(Byte_Counter2)
	buf := new(bytes.Buffer)       // bytes.buffer是一个缓冲byte类型的“Buffer”缓冲器。里面存放着都是byte类型的数据。
	buf.WriteString(`yinzhengjie`) // 往缓冲器中写入字节类型，注意写入是用的符号哟！
	w := io.MultiWriter(l, b)      // 可以理解将l和b方法传给w，也就是说w具有这两种方法去处理数据。
	io.Copy(w, buf)                // 将缓存的数据传给w,这样w就可以调用它的方法去执行相应的代码啦。
	fmt.Printf("该文本的行号是：\033[31;1m%d\033[0m行；\n", l.Sum)
	fmt.Printf("该文本的字节大小是：\033[31;1m%d\033[0m字节.\n", b.Sum)
}
