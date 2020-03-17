package main

import "fmt"

// 常量：定义之后不能修改
const pi = 3.1415926
const (
	statusOk = 200
	notFund  = 404
)

// 批量声明常量时：如果某一行声明后没有赋值，默认跟上面的一样
const (
	n1 = 100
	n2
	n3
)

// iota:类似枚举  在const出现时就会重置为0，每增加一行就会加1
const (
	num1 = iota // 0
	num2        //s1
	num3        //s2
)

func main() {
	fmt.Println(n2)
	fmt.Println(num2, num3)
	fmt.Println(1 << 10)
}
