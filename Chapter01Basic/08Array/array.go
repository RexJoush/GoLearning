package main

import "fmt"

/*
	go 的数组
*/
func main() {
	// 1.go 的静态数组
	var arr1 [10]int

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 2.go 数组的另外一种定义方式和遍历方式
	arr2 := [10]int{1, 2, 3, 4} // 前四个元素是 1，2，3，4 后面都是 0
	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	// 3.查看数组的数据类型
	fmt.Printf("arr1, type = %T\n", arr1) // type = [10]int
	fmt.Printf("arr2, type = %T\n", arr2) // type = [10]int

	// 4.传递数组
	arr3 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr3, type = %T\n", arr3) // type = [4]int
	sum := sumArr(arr3)
	fmt.Println(sum)
}

/*
	数组传参，此处形参必须写明数组长度，因为不同长度的数组属于不同的数据类型
	同时，当传递固定数组的时候，传递的也是值拷贝，而非引用传递，即在函数中修改后，无法带回原来的地方
*/
func sumArr(arr [4]int) int {
	result := 0
	// 此处，如果 index 不用，就写成 _ 即可，也需要占位
	for _, value := range arr {
		result += value
	}
	return result
}
