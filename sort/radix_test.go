package sort

import (
	"fmt"
	"testing"
)

// 测试基数排序
func TestRadix(t *testing.T) {
	fmt.Println(t.Name())
	data := []int64{-6, 1, 2, -7, 9, -3, 4, 5, -10, 8}
	fmt.Println("输入为：", data)

	fmt.Println("期望为： [-10 -7 -6 -3 1 2 4 5 8 9]")

	Radix(data)

	fmt.Println("结果为：", data)
}

func TestRadixBase(t *testing.T) {
	fmt.Println(t.Name())
	data := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Println("输入为：", data)

	fmt.Println("期望为： [2 3 4 5 15 19 26 27 36 38 44 46 47 48 50]")

	RadixBase(data)

	fmt.Println("结果为：", data)
}
