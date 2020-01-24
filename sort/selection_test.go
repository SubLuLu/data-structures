package sort

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	fmt.Println(t.Name())

	ints := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}

	fmt.Println("输入为：", ints)

	fmt.Println("期望为： [2 3 4 5 15 19 26 27 36 38 44 46 47 48 50]")

	data := IntSlice(ints)

	Select(data)

	fmt.Println("结果为：", data)
}
