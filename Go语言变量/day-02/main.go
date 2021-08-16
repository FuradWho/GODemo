package main

import "fmt"

func do01() {
	var a int      //声明变量
	a = 3          //给变量赋值
	fmt.Println(a) //打印变量
}

func do02() {
	var b = 4      //声明变量
	fmt.Println(b) //打印变量
}

func do03() {
	a := "hello "
	b := "world!"
	c := 123
	fmt.Println(a, b, c)
}

func GetData() (int, int) {
	return 100, 200
}

func do05() {
	a := 5
	var b int
	b = a             //将a赋值给b
	p := &a           //获取a的内存地址
	q := &b           //获取b的内存地址
	fmt.Println(p, q) //对于地址进行打印 分析
}

func do06() {
	a, b := 2, "das"
	fmt.Println(a, b)
}

func do07(a, b int) int {
	fmt.Printf("do07() 函数中 a = %d\n", a)
	fmt.Printf("do07() 函数中 b = %d\n", b)
	num := a + b
	return num
}

func do08() {
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

}

func do09() {
	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func main() {
	do09()
}
