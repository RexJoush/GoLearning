package main //  包名，每个 main 包即是主函数的所在位置
/*
	go 的 hello world 代码
*/
// 导包
import (
	"fmt"
	"time"
)

func main() { // 此处必须在同一行，不然就编译报错
	// go 中的分号可加可不加
	fmt.Println("hello Go!") // 输入输出
	// 睡一秒
	time.Sleep(1 * time.Second)
}
