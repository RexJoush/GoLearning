package main

/*
	go 的函数定义方式
*/
import "fmt"

// 定义函数, 两形参，一个返回值
func foo1(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100
	return b + c
}

// 两参数，两返回值，匿名返回值
func foo2(a string, b int) (int, int) {
	return 666, 777
}

// 两参数，两返回值，有形参名返回值
func foo3(a string, b int) (r1 int, r2 int) {
	// 给有名称的返回，如果未赋值，那么 r1, r2 属于两个形参，初始化默认值是 0，作用域在函数体内
	fmt.Println("r1 =", r1, "r2 =", r2) // r1 = 0 r2 = 0

	r1 = 1000
	r2 = 2000
	return
}

// 两参数，两返回值，有形参名返回值，且类型相同，可以并一起写
func foo4(a string, b int) (r1, r2 int) {
	// 给有名称的返回
	r1 = 1000
	r2 = 2000
	return
}
func main() {
	i := foo1("hello", 10)
	fmt.Println(i)

	a, b := foo2("hello", 20)
	fmt.Println("a =", a, "b =", b)

	r1, r2 := foo3("hello", 30)
	fmt.Println("r1 =", r1, "r2 =", r2)

}
