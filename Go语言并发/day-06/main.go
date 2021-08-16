package main

import (
	"fmt"
	"time"
)

func running()  {

	var times int

	// 构建一个无限循环
	for{
		times++
		fmt.Println("时间序列",times)

		// 延时1秒
		time.Sleep(time.Second)
	}
}

func tryChan01() {
	ch := make(chan int)

	ch <- 0
}

func tryChan02()  {
	// 构建一个通道
	ch := make(chan string)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("创建一个新线程")

		//通过通道通知main的goroutine
		ch <- "helo main"

		fmt.Println("线程结束")
	}()

	fmt.Println("等待开始")

	// 等待匿名goroutine
	<- ch

	fmt.Println("结束")

}

func tryChan03(c chan int)  {
	// 开始无限循环等待数据
		for{
			// 从channel中获取一个数据 将0视为数据结束
			if data := <-c ;data == 0 {
				break
			}else{
				fmt.Println(data)
			}

		}
	// 通知main已经结束循环(我搞定了!)
		c <- 0
}


func main()  {
	// 创建一个channel
	c := make(chan int)
	// tryChan03, 传入channel
	go tryChan03(c)

	for i := 1;i<=10;i++{
		// 将数据通过channel投送给线程
		c <- i
	}
	// 通知并发的线程结束循环(没数据啦!)
	c<-0
	// 等待线程结束(搞定喊我!)
	<-c
}
