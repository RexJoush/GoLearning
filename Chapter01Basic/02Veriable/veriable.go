package main // 声明变量
/*
	go 的变量
*/
import "fmt"

// 声明全局变量，只能使用方法一二三进行定义
var gA int = 100
var gB = 200

func main() {
	// 方法一：声明一个变量，默认值是 0
	var a int
	fmt.Println("a = ", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)

	// 方法三：声明一个变量，省略数据类型初始化一个值
	var c = 100
	fmt.Println("c = ", c)
	// 打印一个变量的数据类型，用 %T 打印
	fmt.Printf("type of c : %T", c)

	// 方法四，此方法必须在函数体内使用，省略 var 关键字，自动匹配
	d := 100
	fmt.Printf("d: \n%d %T", d, d)
	e := 3.14
	fmt.Printf("e: \n%f %T", e, e)

	fmt.Println("gA", gA)
	fmt.Println("gB", gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, "yy =", yy)
	// 多个变量也支持不同类型同时定义
	var kk, ll = 100, "Hello"
	fmt.Println("kk =", kk, "ll =", ll)

	// 可以拆成多行写
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, "jj =", jj)

	mm, nn := 30.10, 'c'
	fmt.Println("mm =", mm, "nn =", nn)
}
