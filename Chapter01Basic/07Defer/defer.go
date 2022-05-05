package main

import "fmt"

/*
	go 的 defer, 相当于 C++ 的析构函数和 Java 的 finalized 方法
*/

func main() {
	defer fmt.Println("main end1") // defer 需要写在 return 之前.
	defer fmt.Println("main end2") // 此处，压栈的方式，所以先执行 end2，后执行 end1

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	returnAndDefer()
}

// defer 和 return 谁先谁后，结果是 return 先执行，defer 后执行
func returnAndDefer() int {
	defer deferFunc()   //
	return returnFunc() //
}

func returnFunc() int {
	fmt.Println("return function called...")
	return 1
}
func deferFunc() int {
	fmt.Println("defer function called...")
	return 0
}
