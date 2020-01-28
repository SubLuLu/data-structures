package queue

import (
	"fmt"
	"testing"
)

// 链式队列测试
var lq = NewLinkListQueue()

func TestLinkListQueue_Enqueue(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		err := lq.Enqueue(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestLinkListQueue_Traverse(t *testing.T) {
	fmt.Println(t.Name())

	data := lq.Traverse()

	fmt.Println(data)
}

func TestLinkListQueue_Dequeue(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		v, err := lq.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("dequeue value is", v)
		}
	}
}
