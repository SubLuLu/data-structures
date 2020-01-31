package sort

import (
	"fmt"
	"testing"
)

// 测试桶排序
func TestBucket(t *testing.T) {
	fmt.Println(t.Name())

	data := []int{18, 11, 28, 45, 23, 50, 67, 98, 76}

	fmt.Println("输入为：", data)

	fmt.Println("期望为： [11 18 23 28 45 50 67 76 98]")

	Bucket(data)

	fmt.Println("结果为：", data)
}
