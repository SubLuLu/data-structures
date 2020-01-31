package sort

import (
	"fmt"
	"testing"
)

// 测试快速排序
func TestQuickBaseLeft(t *testing.T) {
	fmt.Println(t.Name())
	data := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println("输入为：", data)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9 10]")

	QuickBaseLeft(data, 0, len(data) - 1)

	fmt.Println("结果为：", data)
}

func TestQuickBaseMiddle(t *testing.T) {
	fmt.Println(t.Name())
	data := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println("输入为：", data)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9 10]")

	QuickBaseMiddle(data, 0, len(data) - 1)

	fmt.Println("结果为：", data)
}

func TestQuickBaseRight(t *testing.T) {
	fmt.Println(t.Name())
	data := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println("输入为：", data)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9 10]")

	QuickBaseRight(data, 0, len(data) - 1)

	fmt.Println("结果为：", data)
}
