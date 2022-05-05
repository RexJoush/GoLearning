package main

/*
	go 的常量
*/
import "fmt"

// 使用 const 和 iota 来实现枚举类型，注意，iota 只能出现在 const 中
const (
	/*
		在 const 中添加关键字 iota 每行的 iota 都会加 1
	*/
	BEIJING  = iota // 此处默认 北京 = 0
	SHANGHAI        // 默认等于 1
	SHENZHEN        // 默认等于 2
)
const (
	BEIJING1  = 10 * iota // 此处 iota = 0, 北京 = 10 * 0 = 0
	SHANGHAI1             // 但此处默认等于 10
	SHENZHEN1             // 此处默认等于 20
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2  a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2  c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2  e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3  g = 6, h = 9
	i, j                      // iota = 4, i = iota * 2, j = iota * 3  i = 8, j = 12
)

func main() {
	// 常量定义，只读
	const length = 10

	fmt.Println(length)

	// 即常量不可修改 length = 20

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

}
