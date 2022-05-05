package main

/*
	go 中的接口，用于实现多态，同时表示万能类型，类似 java 的 Object
*/
import "fmt"

type AnimalIF interface {
	Sleep()           // 睡觉
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的类型
}

// Cat 具体的实现
type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("Cat.Sleep()...")
}
func (c *Cat) GetColor() string {
	return c.color
}
func (c *Cat) GetType() string {
	return "Cat"
}

// Dog 具体的实现
type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("Dog.Sleep()...")
}
func (d *Dog) GetColor() string {
	return d.color
}
func (d *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("color =", animal.GetColor())
	fmt.Println("type =", animal.GetType())
}
func main() {
	var animal AnimalIF // 定义接口的对象，父类指针
	animal = &Cat{"Green"}
	showAnimal(animal)

	animal = &Dog{"Yellow"}
	showAnimal(animal)

	// 测试万能类型
	cat := Cat{"green"}
	myFunc(cat)
	myFunc(100)
	myFunc("aaa")
	myFunc(3.14)
	// 判断数据类型
	myFunc2(3.14)
}

/*
	万能类型，接受所有类型
*/
func myFunc(arg interface{}) {
	fmt.Println("myFunc()...")
	fmt.Println(arg)

}

// 断言，用来判断某种类型，必须是 interface{} 类型才可以判断
func myFunc2(arg interface{}) {
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string, ", value)
		fmt.Printf("type is %T", value)
	} else {
		fmt.Println("arg is not string")
	}

}
