package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
	go 反射解析结构体标签，tag，个人感觉有点像 java 的注解
*/
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str)

	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagString)
		fmt.Println("doc:", tagDoc)
	}
}

func main() {
	findTag(resume{"Rex Joush", "male"})

	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	// 1.结构体转化成 json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		此处因为打印的是字节流，所以使用 ln 的话，出现的是字符的 arscii 码
		fmt.Println("json: ", jsonStr)
		可以使用强制转换
		fmt.Println("json: ", string(jsonStr))
	*/
	fmt.Println("json:", string(jsonStr))
	fmt.Printf("json: %s\n", jsonStr)

	// 2.将 json 转化为结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", myMovie)
}
