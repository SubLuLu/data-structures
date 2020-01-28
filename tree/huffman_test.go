package tree

import (
	"fmt"
	"testing"
)

// Huffman树的构造测试

func TestHuffman_PreOrder(t *testing.T) {
	weights := []int{5, 29, 7, 8, 14, 23, 3, 11}
	tree := NewHuffman(weights)

	data := tree.PreOrder()

	fmt.Println("前序遍历预期为： [100 42 19 8 11 23 58 29 14 15 7 8 3 5 29]")

	fmt.Println("前序遍历结果为：", data)
}
