package sort

// 堆排序
// 1. 创建小顶堆，取出对顶值
// 2. 数组剩余部分再创建小顶堆
// 3. 重复1-2步骤，直到数组只剩下一个元素
func Heap(data []int) {
	n := len(data)
	// golang中切片底层是共享的一个数组
	for i := 0; i < n - 1; i++ {
		buildMinHeap(data[i:])
	}
}

// 小顶堆
type minHeap struct {
	length int   // 堆中有效数据
	size   int   // 堆申请的数组大小
	data   []int // 数据
}

// shiftUp 如果一个节点比它的父节点小
// 则将它同父节点交换位置，并返回true
// 如果没有进行过交换，则返回false
func (mh *minHeap) shiftUp(index int) bool {
	parent := (index - 1) / 2
	if parent < 0 {
		return false
	}
	if mh.data[parent] > mh.data[index] {
		mh.data[parent], mh.data[index] = mh.data[index], mh.data[parent]
		return true
	}
	return false
}

// buildMinHeap 根据给定的数组创建小顶堆
// 从数组中的第二个元素开始进行Push操作
func buildMinHeap(data []int) {
	length := len(data)
	mh := &minHeap{
		length: length,
		size:   length,
		data:   data,
	}

	for i := 1; i < length; i++ {
		for j := i; j >= 0; j = (j - 1) / 2 {
			if !mh.shiftUp(j) {
				break
			}
		}
	}
}
