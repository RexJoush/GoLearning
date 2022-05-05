package main

import (
	"fmt"
	"time"
)

/*
	channel 用来在 goroutine 之间进行通信
*/

func main() {
	// NoCacheChannel() // 无缓冲管道
	// CacheChannel() // 有缓冲管道
	// CloseChannel() // 关闭管道
	// RangeChannel() // 通过 range 关键字读取 channel
	SelectChannel() //
}

func NoCacheChannel() {
	// 定义一个 channel，存储 int 值，同时没有缓存，最多只能存储一个元素
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine end")
		fmt.Println("goroutine running...")

		c <- 666 // 将 666 放入 channel 中
	}()

	/*
		从 c 中读取数据，并赋值给 num
		同时，main 一定能够读取到另外一个协程中的数据，即 channel 自动做了同步
		如果 main 执行到这一步时，go 协程没有写入数据，那么 main 协程就会阻塞，等待写入后再执行
		同时，如果 go 先写到 channel 中，如果 main 还没有读取，因为定义的是无缓冲的 channel
		所以，go 协程也会阻塞，等待 main 协程读取完后，继续执行
	*/
	num := <-c
	fmt.Println("num:", num)
}

func CacheChannel() {
	// 创建一个 3 个容量的缓冲 channel
	c := make(chan int, 3)
	fmt.Println("len(c) =", len(c), "cap(c) =", cap(c)) // len(c) = 0 cap(c) = 3

	go func() {
		defer fmt.Println("go end...")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("go 正在运行，发送的元素=", i, "len(c) =", len(c), "cap(c) =", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num =", num)
	}

	fmt.Println("main end...")
}

// CloseChannel 关闭 channel
func CloseChannel() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- 1
		}
		// 使用 关键字 close 关闭 channel，如果此处不关闭的话，那么主线程就会被阻塞，造成死锁
		close(c)
	}()

	for {
		// if 的增强写法，先判断管道是否关闭，如果没关闭，就进入 if
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("finished...")
}

func RangeChannel() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- 1
		}
		close(c)
	}()

	// 可以使用 range 不断迭代 channel，获取数据
	for data := range c {
		fmt.Println(data)
	}
}

/*
	默认情况下，go 只能监控一个 channel的情况，通过 select 可以监控一组 channel 的情况
*/
func SelectChannel() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		// 退出
		quit <- 0
	}()

	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		// c 可读的话，就进行计算
		case c <- x:
			y = x + y
			x = y - x
		// quit 可读，即结束写入了，就退出
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
