package main

import "fmt"

// 变量声明: 单独声明
// var name string
// var age int
// var isOk bool

// 批量声明

// 局部变量声明后必须使用，不然编译不过
var (
	name string
	age  int
	isOk bool
	hhh  string
)

func main() {
	name = "伍小龙"
	age = 16
	isOk = false

	fmt.Print("age:", age)      // 在终端输入要打印出的内容
	fmt.Println("isOK:", isOk)  // 打印完指定内容后，在后面加一个换行符
	fmt.Printf("name:%s", name) // %s: 占位符，使用name这个变量去替换占位符

	// 类型推导（go自动根据值获取数据类型）
	var s2 = "s2"
	fmt.Println(s2)

	// 简短变量声明，只能在函数里用
	s3 := "哈哈哈" //常见写法

	// 匿名变量: _ 不占空间 不占内存
}
