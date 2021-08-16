package util

import (
	"container/list"
	"fmt"
)

func PrintList(l *list.List) {

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
