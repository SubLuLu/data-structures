package sort

import (
	"fmt"
	"testing"
)

// 测试计数排序
func TestCounting(t *testing.T) {
	fmt.Println(t.Name())

	data := []int{2, 3, 8, 7, 1, 2, 2, 2, 7, 3, 9, 8, 2, 1, 4, 2, 4, 6, 9, 2}

	fmt.Println("输入为：", data)

	fmt.Println("期望为： [1 1 2 2 2 2 2 2 2 3 3 4 4 6 7 7 8 8 9 9]")

	Counting(data)

	fmt.Println("结果为：", data)
}
