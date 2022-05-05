package main

/*
	go 中的 map 集合
*/
import "fmt"

func main() {
	// map 定义
	//mapDefine()
	// map 使用
	mapUse()
}

func mapUse() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"

	fmt.Println("--------")
	printMap(cityMap)
	changeMap(cityMap)
	printMap(cityMap)

}

// 传递修改 map，map 是引用传递，因此对 map 的修改会导致原 map 的修改
func changeMap(cityMap map[string]string) {
	cityMap["England"] = "Landon"
}

// 打印 map
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}
}

func mapDefine() {
	// 1.方式一：声明 map 是一种 map 类型，key 是 string; value 是 string
	var map1 map[string]string

	if map1 == nil {
		fmt.Println("map1 是一个空 map")
	}
	// 再使用前，需要分配容量，初始化为 10 个容量的 map
	map1 = make(map[string]string, 10)
	map1["one"] = "java"
	map1["two"] = "C++"
	map1["three"] = "Go"
	// map 是一个非有序集合，默认按照 key 进行排序
	fmt.Println(map1) // map[one:java three:Go two:C++]

	// 2.方式二：不指定空间
	map2 := make(map[int]string)
	map2[1] = "Java"
	map2[2] = "C++"
	map2[3] = "Go"
	fmt.Println(map2) // map[1:Java 2:C++ 3:Go]

	// 3.方式三：声明直接赋值
	map3 := map[string]string{
		"one":   "Java",
		"two":   "C++",
		"three": "Go", // 最后一行也需要有 ,
	}
	fmt.Println(map3) // map[one:Java three:Go two:C++]

}
