package heap

import (
	"fmt"
	"testing"
)

// 测试大顶堆
var maxH = NewMaxHeap(6)

func TestMaxHeap_Push(t *testing.T) {
	maxH.Push(7)
	maxH.Push(10)
	maxH.Push(2)
	maxH.Push(5)
	maxH.Push(1)
	maxH.Push(16)
	fmt.Println("期望为： [16 7 10 5 1 2]")
	fmt.Println("结果是：", maxH.Traverse())
}

func TestMaxHeap_Peek(t *testing.T) {
	root, err := maxH.Peek()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("期望为： 16")
	fmt.Println("结果是：", root)
}

func TestMaxHeap_Top(t *testing.T) {
	root, err := maxH.Top()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("根的期望为： 16")
	fmt.Println("根的结果为：", root)

	fmt.Println("删除根后的期望为： [10 7 2 5 1]")
	fmt.Println("删除根后的结果为：", maxH.Traverse())
}

func TestMaxHeap_Delete(t *testing.T) {
	err := maxH.Delete(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("删除下标为1的结点后的期望为： [10 5 2 1]")
	fmt.Println("删除下标为1的结点后的结果为：", maxH.Traverse())
}

func TestBuildMaxHeap(t *testing.T) {
	data := []int{7, 10, 2, 5, 1, 16}
	mh := BuildMaxHeap(data)
	fmt.Println("期望为： [16 7 10 5 1 2]")
	fmt.Println("结果是：", mh.Traverse())
}
