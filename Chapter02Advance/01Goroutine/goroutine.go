package main

/*
	go 的协程 goroutine
*/
import (
	"fmt"
	"time"
)

// 创建一个线程任务
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("New Goroutine: i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main goroutine: i=%d\n", i)
	//	time.Sleep(1 * time.Second)
	//}
	createGoroutine2()
}

func createGoroutine1() {
	// 创建一个 go 线程，去执行 newTask 流程，此处，如果 main 协程退出了，那么 main 创建的所有协程也将退出
	go newTask()
}

func createGoroutine2() {
	// 使用匿名函数创建 go 协程
	go func() {
		defer fmt.Println("A.defer")
		// 此处，如果加上 return，表示退出本层 go 函数，即下面的代码都不会执行
		// return
		func() {
			defer fmt.Println("B.defer")
			// 此处，return 表示退出本层 func，即，不会打印出 B，但，仅仅退出了本层 func，外层不会退出
			// return
			// 如果想要退出外层的循环，那么就需要使用下面的方式
			// runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}

func createGoroutine3() {
	// 定义一个协程带返回值，但此处返回值无法通过变量来获取，
	// 因为 此协程和主协程是一个层级的关系，如果想要获取结果，需要使用 channel
	/*
		sum := go func(a, b int) int {
			fmt.Println("a:", a, "b:", b)
			return a + b
		}(1, 2)
	*/
}
