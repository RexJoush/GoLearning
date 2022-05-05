package main

/*
	go 中的结构体
*/
import "fmt"

// 定义一个数据类型 myInt，此时，myInt 和 int 等价
type myInt int

// Book 定义一个结构体
type Book struct {
	title string // 书名
	auth  string // 作者
}

func main() {
	/*
		var a myInt = 10
		println(a) // 10
	*/

	// 定义一个 book 结构体
	var book Book
	book.title = "Go"
	book.auth = "zhang3"
	fmt.Printf("%v\n", book) // {Go zhang3}

	changeBook(book)
	fmt.Printf("%v\n", book) // {Go zhang3}

	changeBook2(&book)
	fmt.Printf("%v\n", book) // {Go 666}
}

// 结构体传参是值传递，所以此处没有修改成功
func changeBook(book Book) {
	book.auth = "666"
}

// 想要修改则需要引用传递
func changeBook2(book *Book) {
	book.auth = "666"
}
