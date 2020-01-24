package sort

import (
	"math"
)

// 桶排序
// 1. 确定桶的数量
// 2. 确定每个桶需要存储的数据范围(平均分配)
// 3. 对每个桶中的数据采用合适的排序算法进行排序
// 4. 依次将桶中的数据取出
func Bucket(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	// 序列中的最值
	min, max := keyPoint(data)
	// 桶的数量
	num := bucketNum(n)
	// 每个桶中的储存数据范围
	scope := float64(max - min) / float64(num)
	// 创建空桶
	bins := make([]*node, num)

	// 将数据分配到桶中
	for _, v := range data {
		// 计算v需要存储的桶的位置
		// v <= min + n * scope
		// n >= (v - min) / scope
		i := int(math.Ceil(float64(v - min) / scope))
		if i != 0 { // 桶的位置和切片下标差1
			i -= 1
		}

		bin := bins[i]
		if bin == nil { // 使用时初始化桶
			bin = newNode()
			bins[i] = bin
		}
		// 对单个桶中的元素进行排序
		bin.insertVal(v)
	}

	// 遍历桶中的所有元素，依次填充到序列中
	var index int
	for _, b := range bins {
		if b != nil { // 防止出现空桶
			temp := b.Next
			for ; temp != nil; temp = temp.Next {
				data[index] = temp.Val
				index++
			}
		}
	}
}

// 计算桶数量
// 桶的数量以2^num=length确定
func bucketNum(length int) int {
	for num := 1; ;num++ {
		t := length >> uint(num)
		if t == 0 {
			return num
		}
	}
}

// 找出最大值和最小值
func keyPoint(data []int) (int, int) {
	n := len(data)
	// 默认第一个即为最大值，也为最小值
	min, max := data[0], data[0]
	for i := 1; i < n; i++  {
		if data[i] > max {
			max = data[i]
		}
		if data[i] < min {
			min = data[i]
		}
	}
	return min, max
}

// 每个桶使用双向链表进行存储
type node struct {
	Val int
	Next *node
	Prev *node
}

func newNode() *node {
	return new(node)
}

// 插入时进行排序
func (n *node) insertVal(val int) {
	data := newNode()
	data.Val = val
	// 链表为空，直接插入
	if n.Next == nil {
		n.Next = data
		data.Prev = n
		return
	}

	temp := n.Next
	for temp != nil {
		// 在链表中遇到比val大的，则插入到前面
		if temp.Val > val {
			temp.Prev.Next = data
			data.Next = temp
			data.Prev = temp.Prev
			temp.Prev = data
			return
		}
		// 在链表中没有比val大的，插在末尾
		if temp.Next == nil {
			temp.Next = data
			data.Prev = temp
			return
		}
		temp = temp.Next
	}
}
