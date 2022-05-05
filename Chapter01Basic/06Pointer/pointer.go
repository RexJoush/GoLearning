package main

/*
	go 的指针
*/
import "fmt"

func main() {
	a := 1
	changeValue(a)   // 值传递
	fmt.Println(a)   // 1, 此处 a 并未被改变，即值传递
	changeValue2(&a) // 指针传递
	fmt.Println(a)   // 10, 此处 a 被改变了，即指针传递，和 c 语言一样
	m, n := 2, 3
	swap(&m, &n) // 函数交换
	fmt.Println("m =", m, "n =", n)

	// 二级指针
	p := &a
	fmt.Println(&a) // 0xc000018098, a 的地址
	fmt.Println(p)  // 0xc000018098, p 中存的是 a 的地址

	var pp **int
	pp = &p         // pp 存的是 p 的地址
	fmt.Println(&p) // 0xc000006030, p 的地址
	fmt.Println(pp) // 0xc000006030, pp 的值
}

func changeValue(p int) {
	p = 10
}

// 此处 c 语言的理解方式， * 表示指针传递
func changeValue2(p *int) {
	*p = 10
}

func swap(a, b *int) {
	p := *a
	*a = *b
	*b = p
}
