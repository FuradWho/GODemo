package main

import (
	"fmt"
)

func do01() {

	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i + 10
	}

	for i, x := range n {
		fmt.Printf("第%d个元素，值为%d\n", i, x)
	}

	fmt.Printf("数组元素个数：%d\n", len(n))

}

func do02() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
}

func do03() {
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /*  第三行索引为 2 */
	}

	fmt.Println(a[2][3])

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func do04(arr [][]int) {

	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}
}

func do05() {
	numbers := make([]int, 3, 5)
	printSlice(numbers)
}

func do06() {
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func do07() {
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func do08() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

func do09() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func do10() {
	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}

func do11() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
}

func do12() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示其值
	for _, value := range slice {
		fmt.Printf("Value: %d\n", value)
	}
}

func do13() {

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func do14() {

	var mapList map[string]int

	mapList = map[string]int{"one": 1, "two": 2}

	for k, v := range mapList {
		fmt.Printf("%s -> %d\n", k, v)
	}

	mapList1 := make(map[string]float32)

	mapList1["key1"] = 4.5
	mapList1["key2"] = 3.1425

	for k, v := range mapList1 {
		fmt.Printf("%s -> %f\n", k, v)
	}
}

func do15() {

	var mapList map[string]int

	mapList = map[string]int{"one": 1, "two": 2}

	delete(mapList, "one")

	for k, v := range mapList {
		fmt.Printf("%s -> %d\n", k, v)
	}

}
func main() {
	do15()
}
