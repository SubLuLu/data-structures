package sort

import (
	"fmt"
	"testing"
)

// 测试直接插入排序
func TestInsert(t *testing.T) {
	fmt.Println(t.Name())

	ints := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}

	fmt.Println("输入为：", ints)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9]")

	Insertion(ints)

	fmt.Println("结果为：", ints)
}
