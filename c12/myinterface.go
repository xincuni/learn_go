package main

import "fmt"

// 定义协议
type Programmer interface {
	Coding() string //方法只是声明,
	Debug() string
}

type Designe interface {
	Design() string
}

type Manager interface {
	Programmer
	Designe
	Manage() string
}

type Pythoner struct {
	UiDesigne
}

type UiDesigne struct {
}

func (u UiDesigne) Design() string {
	fmt.Println("我是ui设计师")
	return "我会设计"
}

func (p Pythoner) Coding() string {
	fmt.Println("python 开发者")
	return "Python开发者"
}

// 对于Pythoner来说实现任何方法都可以,但是只要不全部实现,那么你就不是一个程序员
// 1. 对于python来说,本身就是一个类型,那么我何必在意我是不是programmer呢?
// 2. 是为了多态
// 3. 在go语言中,interface是一种类型,是一种抽象类型,那么我可以基于这个类型去声明一个变量

func (p Pythoner) Debug() string {
	fmt.Println("我会debug")
	return "我是debug"
}

func (p Pythoner) Manage() string {
	fmt.Println("我也会管理")
	return "我也会管理"
}

//func (p Pythoner) Design() string {
//	fmt.Println("我是Python开发者我也会设计")
//	return "我也会设计"
//}

type G struct {
}

func (g G) Coding() string {
	fmt.Println("go 开发者")
	return "go开发者"
}
func (g G) Debug() string {
	fmt.Println("go debug")
	return "godebug"
}

//func PyHandle (pythoner Pythoner){
//
//}
//func GHandle (G G){
//
//}

func Handle(programmer Programmer) {

}

func main() {
	var pro Programmer = G{}
	pro.Coding()
	pro.Debug()
	var pros []Programmer
	pros = append(pros, Pythoner{}, G{})
	fmt.Println(pros)

	// 接口虽然是一种类型,但是和其他类型不同,接口是抽象类型,它里面只有方法;struct是具象内容(参数,方法).抽象类型的赋值又必须是具象类型

	var m Manager = Pythoner{}
	m.Design()

	s := "文件不存在"
	err := fmt.Errorf("错误:%s", s)
	fmt.Println(err)
}

// Python是没有多态的,因为它是动态语言,可以基于一个变量赋值不同的类型
