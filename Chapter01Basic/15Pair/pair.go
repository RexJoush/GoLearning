package main

import (
	"fmt"
	"io"
	"os"
)

/*
	go 变量
			变量
		/		    \
	 type 类型     value 值 	// 这两个组成了 pair
	/         \
static type   concrete type
int,string..      interface所指向的具体数据类型、系统看的见的类型
	反射 指的是，通过一个变量获取当前变量的 type，即具体类型，也可获得值
*/

func main() {
	var a string

	// pair<static_type:string, value:"ace>
	a = "ace"

	// pair<type:string, value:"ace">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
	//file()
	reflectTest()
}

func file() {
	// pair<type:*os.File, value:"D:\Files\in.txt"> 文件描述符
	tty, err := os.Open("D:\\Files\\in.txt")
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// pair<type: , value: >
	var r io.Reader
	// pair<type:*io.Reader, value:"D:\Files\in.txt"> 文件描述符
	r = tty

	// pair<type: , value: >
	var w io.Writer
	// pair<type: *os.File, value:"D:\Files\in.txt"> 文件描述符
	w = r.(io.Writer)

	w.Write([]byte("Hello, this is a test!"))

}

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// Book 具体类型
type Book struct {
}

func (b Book) ReadBook() {
	fmt.Println("Read a Book!")
}

func (b Book) WriteBook() {
	fmt.Println("Write a Book!")
}

func reflectTest() {
	// pair<type:Book, value:book{}地址>
	b := &Book{}

	// pair<type:, value:>
	var r Reader
	// pair<type:Book, value:book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	// pair<type:Book, value:book{}地址>
	w = r.(Writer)

	w.WriteBook()
}