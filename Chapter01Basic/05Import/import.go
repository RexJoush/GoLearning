package main

/*
	go 的导包和 init 函数
*/
import (
	_ "GoLearning/05Import/lib1" // 导包但不使用，意义在于执行包 lib1 的 init 函数，但并不调用 lib1 的方法
	. "GoLearning/Chapter01Basic/05Import/lib2"
)

func main() {
	// 回默认优先调用包的 init 函数，函数名首字母大写是因为，大写默认开放函数，可以当 public 理解
	// lib1.Lib1Test()
	// myLib2.Lib2Test()
	Lib2Test() // 直接使用
}
