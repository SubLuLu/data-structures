package stack

import (
	"fmt"
	"testing"
)

// 链式栈测试
var ls = NewLinkListStack()

func TestLinkListStack_Push(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		err := ls.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestLinkListStack_Traverse(t *testing.T) {
	fmt.Println(t.Name())

	data := ls.Traverse()

	fmt.Println(data)
}

func TestLinkListStack_Pop(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		v, err := ls.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("pop value is", v)
		}
	}
}
