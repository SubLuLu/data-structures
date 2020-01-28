package tree

import (
	"fmt"
	"testing"
)

// 二叉树递归遍历测试

var (
	ab = newArrayBinary()
	lb = newListBinary()
)

func TestArrayBinary_PreOrder(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println("前序遍历期望为：[1 2 4 8 11 5 3 6 9 7 10]")

	fmt.Println("前序遍历结果为：", ab.PreOrder())
}

func TestArrayBinary_InOrder(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println("中序遍历期望为：[8 11 4 2 5 1 9 6 3 7 10]")

	fmt.Println("中序遍历结果为：", ab.InOrder())
}

func TestArrayBinary_PostOrder(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println("后序遍历期望为：[11 8 4 5 2 9 6 10 7 3 1]")

	fmt.Println("后序遍历结果为：", ab.PostOrder())
}

func TestListBinary_PreOrder(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println("前序遍历结果为：", lb.PreOrder().string())
}

func TestListBinary_InOrder(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println("中序遍历结果为：", lb.InOrder().string())
}

func TestListBinary_PostOrder(t *testing.T) {
	fmt.Println(t.Name())

	fmt.Println("后序遍历结果为：", lb.PostOrder().string())
}

// 构建一个顺序存储的二叉树
// 1-2-3-4-5-6-7-8-^-^-^-9-^-^-10-^-11
func newArrayBinary() arrayBinary {
	ab := make(arrayBinary, 17)
	ab[0] = new(int)
	*ab[0] = 1
	ab[1] = new(int)
	*ab[1] = 2
	ab[2] = new(int)
	*ab[2] = 3
	ab[3] = new(int)
	*ab[3] = 4
	ab[4] = new(int)
	*ab[4] = 5
	ab[5] = new(int)
	*ab[5] = 6
	ab[6] = new(int)
	*ab[6] = 7
	ab[7] = new(int)
	*ab[7] = 8
	ab[11] = new(int)
	*ab[11] = 9
	ab[14] = new(int)
	*ab[14] = 10
	ab[16] = new(int)
	*ab[16] = 11

	fmt.Println("二叉树分层遍历输入为：")

	fmt.Print("[ ")
	for _, v := range ab {
		if v != nil {
			fmt.Print(*v)
		} else {
			fmt.Print("nil")
		}
		fmt.Print(" ")
	}
	fmt.Println("]")
	return ab
}

// 构建一个链式存储的二叉树
func newListBinary() *listBinary {
	lb := &listBinary{
		data: 1,
	}

	node11 := &listBinary{
		data:  11,
		left:  nil,
		right: nil,
	}

	node10 := &listBinary{
		data:  10,
		left:  nil,
		right: nil,
	}

	node9 := &listBinary{
		data:  9,
		left:  nil,
		right: nil,
	}

	node8 := &listBinary{
		data:  8,
		left:  nil,
		right: node11,
	}

	node7 := &listBinary{
		data:  7,
		left:  nil,
		right: node10,
	}

	node6 := &listBinary{
		data:  6,
		left:  node9,
		right: nil,
	}

	node5 := &listBinary{
		data:  5,
		left:  nil,
		right: nil,
	}

	node4 := &listBinary{
		data:  4,
		left:  node8,
		right: nil,
	}

	node3 := &listBinary{
		data:  3,
		left:  node6,
		right: node7,
	}

	node2 := &listBinary{
		data:  2,
		left:  node4,
		right: node5,
	}

	lb.left = node2
	lb.right = node3

	return lb
}
