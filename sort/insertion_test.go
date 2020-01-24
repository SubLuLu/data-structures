package sort

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	fmt.Println(t.Name())

	ints := []int{9, 3, 1, 4, 2, 7, 8, 6, 5}

	fmt.Println("输入为：", ints)

	fmt.Println("期望为： [1 2 3 4 5 6 7 8 9]")

	Insertion(ints)

	fmt.Println("结果为：", ints)
}
