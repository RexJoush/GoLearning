package main

import "fmt"

/*
	go 的动态数组 slice 表示切片
*/

func main() {
	slickDefine() // 切片的定义

	sliceAppend() // 切片追加

	sliceSplit() // 切片截取
}
func slickDefine() {
	arr1 := []int{1, 2, 3, 4}
	// fmt.Printf("%T\n", arr1) // []int
	printArray(arr1)

	// slice 的定义方式

	// 1.声明一个切片，并且初始化，默认值是 1,2,3 长度 len 是 3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1) // len = 3, slice = [1 2 3]

	// 2.声明 slice 是一个切片，没有分配空间
	var slice2 []int
	slice2 = make([]int, 3)                                   // 开辟 3 个空间
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2) // len = 3, slice = [0 0 0]

	// 3.声明 slice 是一个切片, 同时分配 3 个空间
	slice3 := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3) // len = 3, slice = [0 0 0]

	// 判断 slice 是否为 0，即是否为空
	if slice3 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Println("非空切片")
	}
}

// 此处表示动态数组，长度不固定，同时属于指针传递，和常规 Java 数组一样了
func printArray(arr []int) {
	for _, value := range arr {
		fmt.Println("value:", value)
	}
	arr[0] = 100
}

func sliceAppend() {
	/*
			长度 3，容量为 5，即 nums 指向头部空间，ptr 指针指向尾部有效的索引位置
			nums	   ptr
			 |           |
			 | 0 | 0 | 0 | | |
		     |--- len ---|
			 |------cap------|
	*/
	nums := make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums) // len = 3, cap = 5, slice = [0 0 0]

	// 将 1 加入切片中
	nums = append(nums, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums) // len = 4, cap = 5, slice = [0 0 0 1]

	// 将 2 加入切片中
	nums = append(nums, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums) // len = 5, cap = 5, slice = [0 0 0 1 2]

	// 将 3 加入切片中，此时容量不够了，那么会再自动开辟一个 cap 的长度的空间，长度翻倍，即 cap = cap * 2
	nums = append(nums, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums) // len = 6, cap = 10, slice = [0 0 0 1 2 3]

	/*
		此处如果不指定 cap，则 cap 默认和 len一样
	*/
	nums2 := make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums2), cap(nums2), nums2)
}

/*
	切片的截取
*/
func sliceSplit() {
	s := []int{1, 2, 3}

	// [0,2], 截取 [0,2)
	s1 := s[0:2]
	fmt.Println(s1) // [1 2]

	// 对于子切片的更改，会造成原切片的更改，类似 js 的浅拷贝
	s1[0] = 100
	fmt.Println(s)  // [100, 2, 3]
	fmt.Println(s1) // [100, 2]

	// copy, 可以将底层数组的 slice 进行拷贝，即类似 js 的深拷贝
	s2 := make([]int, 3)
	// 将 s 中的值拷贝到 s2 中
	copy(s2, s)
	fmt.Println(s2) // [100 2 3]
}
