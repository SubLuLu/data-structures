package tree

import "sort"

type (
	// Huffman树
	huffman struct {
		size int          // 结点个数
		root *huffmanNode // 根结点
	}

	// Huffman树结点
	huffmanNode struct {
		weight int          // 结点权值
		left   *huffmanNode // 左孩子结点
		right  *huffmanNode // 右孩子结点
	}

	// 实现sort.Interface接口，方便排序
	huffmanNodes []*huffmanNode
)

// PostOrder 前序遍历
func (h *huffman) PreOrder() []int {
	// 根据有效结点个数申请相应大小的切片
	result := make([]int, h.size)
	var pos = index() // 索引值用闭包
	// 后续遍历从左子树开始
	orderPre(h.root, result, pos)
	return result
}

func orderPre(node *huffmanNode, result []int, pos func() int) {
	if node == nil {
		return
	}
	result[pos()] = node.weight
	orderPre(node.left, result, pos)
	orderPre(node.right, result, pos)
}

// Len is the number of elements in the collection.
func (hns huffmanNodes) Len() int {
	return len(hns)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (hns huffmanNodes) Less(i, j int) bool {
	return hns[i].weight < hns[j].weight
}

// Swap swaps the elements with indexes i and j.
func (hns huffmanNodes) Swap(i, j int) {
	hns[i], hns[j] = hns[j], hns[i]
}

// NewHuffmanNode 构造一个带权值结点
func NewHuffmanNode(weight int, left, right *huffmanNode) *huffmanNode {
	return &huffmanNode{
		weight: weight,
		left:   left,
		right:  right,
	}
}

// LeafHuffmanNode 构造一个带权值的叶子结点
func LeafHuffmanNode(weight int) *huffmanNode {
	return &huffmanNode{
		weight: weight,
		left:   nil,
		right:  nil,
	}
}

// NewHuffman 根据给定的权值数组构造哈夫曼树
func NewHuffman(weights []int) *huffman {
	l := len(weights)
	if l == 0 { // 权值数组为空，返回空树
		return nil
	}
	if l == 1 { // 权值数组长度为1，返回只有根的树
		return &huffman{
			root: LeafHuffmanNode(weights[0]),
		}
	}
	// 声明一个与权值数组等长的huffmanNodes
	hns := make(huffmanNodes, l)
	// 根据权值数组中的值构造叶子结点，并填充huffmanNodes
	for i, weight := range weights {
		hns[i] = LeafHuffmanNode(weight)
	}
	// 对huffmanNodes进行排序
	sort.Sort(hns)
	var root *huffmanNode
	for len(hns) > 1 { // 当切片中的结点只剩下一个时，即为哈夫曼树的根结点
		// 取前两个结点构造二叉树
		root = NewHuffmanNode(hns[0].weight+hns[1].weight, hns[0], hns[1])
		// 将新得到的根结点放大第一个位置
		hns[1] = root
		// 利用切片性质，删除第一个结点
		hns = hns[1:]
		// 对剩下的结点进行排序
		sort.Sort(hns)
	}

	return &huffman{
		root: hns[0],
		size: 2*l - 1,
	}
}
