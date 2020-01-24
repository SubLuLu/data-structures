package sort

import (
	"fmt"
	"testing"
)

func TestShell(t *testing.T) {
	fmt.Println(t.Name())

	data := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}

	fmt.Println("输入为：", data)

	fmt.Println("期望为： [0 1 2 3 4 5 6 7 8 9]")

	Shell(data)

	fmt.Println("结果为：", data)
}
