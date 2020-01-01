package stack

import (
	"fmt"
	"testing"
)

// 数组实现栈测试

var as = NewArrayStack(5)

func TestArrayStack_Push(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		err := as.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestArrayStack_Traverse(t *testing.T) {
	fmt.Println(t.Name())

	data := as.Traverse()

	fmt.Println(data)
}

func TestArrayStack_Pop(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		v, err := as.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("pop value is", v)
		}
	}
}
