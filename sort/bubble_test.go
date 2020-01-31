package sort

import (
	"fmt"
	"testing"
)

// 测试冒泡排序
func TestBubble(t *testing.T) {
	fmt.Println(t.Name())

	data := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}

	fmt.Println("输入为：", data)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9]")

	Bubble(data)

	fmt.Println("结果为：", data)
}
