package heap

import (
	"fmt"
	"testing"
)

// 测试小顶堆
var minH = NewMinHeap(6)

func TestMinHeap_Push(t *testing.T) {
	minH.Push(7)
	minH.Push(10)
	minH.Push(2)
	minH.Push(5)
	minH.Push(1)
	minH.Push(16)
	fmt.Println("期望为： [1 2 7 10 5 16]")
	fmt.Println("结果是：", minH.Traverse())
}

func TestMinHeap_Peek(t *testing.T) {
	root, err := minH.Peek()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("期望为： 1")
	fmt.Println("结果是：", root)
}

func TestMinHeap_Top(t *testing.T) {
	root, err := minH.Top()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("根的期望为： 1")
	fmt.Println("根的结果为：", root)

	fmt.Println("删除根后的期望为： [16 2 7 10 5]")
	fmt.Println("删除根后的结果为：", minH.Traverse())
}

func TestMinHeap_Delete(t *testing.T) {
	err := minH.Delete(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("删除下标为1的结点后的期望为： [5 16 7 10]")
	fmt.Println("删除下标为1的结点后的结果为：", minH.Traverse())
}

func TestBuildMinHeap(t *testing.T) {
	data := []int{7, 10, 2, 5, 1, 16}
	mh := BuildMinHeap(data)
	fmt.Println("期望为： [1 2 7 10 5 16]")
	fmt.Println("结果是：", mh.Traverse())
}
