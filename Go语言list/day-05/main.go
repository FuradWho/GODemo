package main

import (
	"container/list"
	"day05/filelisting"
	"day05/model"
	"day05/util"
	"fmt"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/log"
)

func do01() {
	l := list.New()
	l.PushBack("fist")
	l.PushFront(67)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func do02() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func do03() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	util.PrintList(l)
}

func do04() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
}

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data string) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data string) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

// 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()     // 开启服务
	Log(string) // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(l string) {
	fmt.Println(l)
}

// 游戏服务
type GameService struct {
	Logger // 嵌入日志器
}

// 实现Service的Start()方法
func (g *GameService) Start() {
}

func do05() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Print(value, ",", ok)

}

func do06() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("defer: line ", i)
	}
	fmt.Println("defer: line  3")
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWarpper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Warn("error occurred handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWarpper(filelisting.HandleFileList))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}

}
