package sort

import (
	"fmt"
	"testing"
)

// 测试堆排序
func TestHeap(t *testing.T) {
	fmt.Println(t.Name())

	data := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}

	fmt.Println("输入为：", data)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9]")

	Heap(data)

	fmt.Println("结果为：", data)
}
