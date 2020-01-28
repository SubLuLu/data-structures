package queue

import (
	"fmt"
	"testing"
)

// 数组实现循环队列测试
var aq = NewArrayQueue(5)

func TestArrayQueue_Enqueue(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		err := aq.Enqueue(i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestArrayQueue_Traverse(t *testing.T) {
	fmt.Println(t.Name())

	data := aq.Traverse()

	fmt.Println(data)
}

func TestArrayQueue_Dequeue(t *testing.T) {
	fmt.Println(t.Name())

	for i := 0; i < 6; i++ {
		v, err := aq.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("dequeue value is", v)
		}
	}
}
