package main

import (
	"fmt"
	"reflect"
)

/*
	反射是指，在程序运行过程中，动态的获取变量的 pair，即类型和值
	reflect 包，两个方法
		func ValueOf(i interface{}) Value {} 返回输入数据的 value 值，如果为空返回 0
		func TypeOf(i interface{}) Type {}   返回输入数据的 type 值，如果为空返回 nil
*/
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is Called")
}

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("type: ", reflect.ValueOf(arg))
}

func main() {
	// var num = 1.2345
	//reflectNum(num)
	var user = User{1, "Rex", 24}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	// 获取输入数据的类型
	inputType := reflect.TypeOf(input)
	fmt.Println("type is", inputType.Name())

	// 获取输入数据的 value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value is", inputValue)

	/*
		通过 type 获取结构体的字段
		1.获取 interface 的 reflect.Type，通过 Type 得到 NumFiled，进行遍历
		2.得到每个 field，即数据类型
		3.通过 field Interface() 方法得到对应的 value
	*/
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	/*
		通过 type 获取结构体的方法
	*/
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}

}
