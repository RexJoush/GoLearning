package main

/*
	go 中的类定义
*/
import "fmt"

/*
	定义一个带方法的类，此处，类名大写，那么其他包也可以定义此对象
	属性首字母大写，则表示公有属性，小写则表示私有属性
*/
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("Ad =", this.Ad)
	fmt.Println("Level =", this.Level)
}
func (this Hero) GetName() {
	fmt.Println("Name =", this.Name)
}

// SetName 此处需要注意，必须使用指针传递，默认是 this 对象的拷贝，而非 this 本身
func (this *Hero) SetName(Name string) {
	this.Name = Name
}

func main() {
	// 创建一个对象
	hero := Hero{"Joush", 24, 1}
	hero.GetName()

	hero.Show()
	hero.SetName("Rex")
	hero.Show() // 发现，姓名并没有被改变，即需要改成引用传递
}
