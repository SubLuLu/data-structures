package tree

import (
	"fmt"
	"testing"
)

// 非递归遍历测试

func TestArrayBinary_PreOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(ab.PreOrderTraverse())
}

func TestArrayBinary_InOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(ab.InOrderTraverse())
}

func TestArrayBinary_PostOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(ab.PostOrderTraverse())
}

func TestListBinary_PreOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(lb.PreOrderTraverse().string())
}

func TestListBinary_InOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(lb.InOrderTraverse().string())
}

func TestListBinary_PostOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(lb.PostOrderTraverse().string())
}

func TestListBinary_LevelOrderTraverse(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println(lb.LevelOrderTraverse().string())
}
