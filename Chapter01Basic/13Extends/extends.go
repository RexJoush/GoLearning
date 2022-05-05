package main

/*
	go 的继承，有点类似组合
*/
import "fmt"

/*
	定义一个 Human 类
*/
type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (h *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// SuperMan 继承 Human 的写法
type SuperMan struct {
	Human // 表示继承，同时 H 大写表示公有继承
	level int
}

// Eat 重写父类的 eat 方法
func (s *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// Fly 子类的新方法
func (s *SuperMan) Fly() {
	fmt.Println("SuperMan.fly()")
}

func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	//s := SuperMan{Human{"li4", "female"}, 88}
	// 另外一种方法
	var s SuperMan
	s.name = "wang5"
	s.sex = "male"
	s.level = 20
	s.Walk() // 调用 父类的方法
	s.Eat()  // 已经重写了，调用子类的方法
	s.Fly()  // 子类的新方法
}
