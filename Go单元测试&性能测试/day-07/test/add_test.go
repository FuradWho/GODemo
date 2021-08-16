package test

import "testing"

func add(a,b int) int {
	return  a + b
}

func TestAdd(t *testing.T)  {

	if add(5,5) != 10{
		t.Error("结果错误！")
	}else{
		t.Log("结果正确！")
	}
}

func Benchamrk(b *testing.B)  {
	for i := 0 ; i<b.N ;i++{ // b.N，测试循环次数
		add(5,5)
	}
}