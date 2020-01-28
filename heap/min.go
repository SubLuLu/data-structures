package heap

import (
	"errors"
	"sync"
)

// 小顶堆
type minHeap struct {
	mu     sync.RWMutex // 读写锁
	length int          // 堆中有效数据
	size   int          // 堆申请的数组大小
	data   []int        // 数据
}

// Peek 返回根结点的值(最小值)
func (mh *minHeap) Peek() (int, error) {
	if mh.IsEmpty() {
		err := errors.New("min heap is empty")
		return 0, err
	}

	mh.mu.RLock()
	defer mh.mu.RUnlock()
	return mh.data[0], nil
}

// Top 返回根结点的值并删除
func (mh *minHeap) Top() (int, error) {
	if mh.IsEmpty() {
		err := errors.New("min heap is empty")
		return 0, err
	}

	mh.mu.RLock()
	root := mh.data[0]
	mh.mu.RUnlock()

	err := mh.Delete(0)
	if err != nil {
		return 0, err
	}
	return root, nil
}

// Push 插入元素
// 根据完全二叉树的性质可知
// 当下标从1开始计数时
// 如果第i个结点有左右孩子结点
// 则左孩子结点是第2*i个
// 右孩子结点是第2*i+1个
// 当下标从0开始计数时
// 则左孩子是2*i+1
// 右孩子是2*i+2
// 父结点为(i - 1) / 2
// 1. 将元素插入到最后一个
// 2. 与父结点进行比较，若比父结点大则插入成功
//    若比父结点小，则与父结点进行交换
// 3. 直到与根结点进行比较
func (mh *minHeap) Push(value int) error {
	if mh.IsFull() {
		return errors.New("heap is full")
	}

	mh.mu.Lock()
	defer mh.mu.Unlock()

	i := mh.length
	// 先将待插入的元素放在最后的叶子结点位置
	mh.data[i] = value
	// 依次和父结点进行比较
	for ; i >= 0; i = (i - 1) / 2 {
		if !mh.shiftUp(i) {
			break
		}
	}
	mh.length++
	return nil
}

// Delete 删除元素
// 分为删除根结点和非根结点两种情况
// 1. 删除的是根结点
//   1.1 用最后一个元素和根结点进行交换
//   1.2 该元素和根结点的左右孩子结点进行比较
//   1.3 和孩子结点中较小的结点进行交换
// 2. 删除的是其他结点
//   2.1 用最后一个元素和该结点进行交换
//   2.2 如果删除的是叶子结点，则和父结点进行比较
//   2.3 如果删除的是非叶子结点，交换后则和孩子结点进行比较
func (mh *minHeap) Delete(index int) error {
	if mh.IsEmpty() {
		return errors.New("min heap is empty")
	}

	mh.mu.Lock()
	defer mh.mu.Unlock()

	// 把最后一个结点赋给待删除的结点
	mh.data[index] = mh.data[mh.length-1]
	// 删除最后一个结点
	mh.length--
	// 删除根结点
	if index == 0 {
		// 依次和孩子结点进行比较
		for ; index < mh.length; {
			if i, ok := mh.shiftDown(index); !ok {
				break
			} else {
				index = i
			}
		}
		return nil
	}
	// 删除叶子结点
	if 2*index+1 >= mh.length {
		// 依次和父结点进行比较
		for ; index >= 0; index = (index - 1) / 2 {
			if !mh.shiftUp(index) {
				break
			}
		}

		return nil
	}
	// 删除非根也非叶子结点
	var flag bool // false表示比父结点小，true表示比父结点大
	// 先与父结点进行比较
	// 如果比父结点小，则不需要与孩子结点进行比较
	// 如果比父结点大，则与孩子结点进行比较

	// 依次和父结点进行比较
	for i := index; i >= 0; i = (i - 1) / 2 {
		if !mh.shiftUp(i) {
			break
		}
		flag = true
	}

	if !flag { // 比父结点小
		// 依次和孩子结点进行比较
		for ; index < mh.length; {
			if i, ok := mh.shiftDown(index); !ok {
				break
			} else {
				index = i
			}
		}
	}

	return nil
}

// IsEmpty 判断是否为空
func (mh *minHeap) IsEmpty() bool {
	mh.mu.RLock()
	defer mh.mu.RUnlock()

	return mh.length == 0
}

// IsFull 判断是否满了
func (mh *minHeap) IsFull() bool {
	mh.mu.RLock()
	defer mh.mu.RUnlock()

	return mh.length == mh.size
}

// Traverse 遍历
func (mh *minHeap) Traverse() []int {
	mh.mu.RLock()
	defer mh.mu.RUnlock()

	return mh.data[:mh.length]
}

// shiftUp 如果一个节点比它的父节点小
// 则将它同父节点交换位置，并返回true
// 如果没有进行过交换，则返回false
// 经常被循环调用，所以方法内部不加锁
// 确保调用该方法的代码加锁
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

// shiftDown 如果一个节点比它的孩子结点小
// 那么需要将它和较大的孩子结点进行交换
// 返回结果 true表示经过交换 false表示未经过交换
// 经常被循环调用，所以方法内部不加锁
// 确保调用该方法的代码加锁
func (mh *minHeap) shiftDown(index int) (int, bool) {
	left := 2*index + 1    // 左孩子结点下标
	right := left + 1      // 右孩子结点下标
	if left >= mh.length { // 已经没有孩子结点了
		return index, false
	}
	if right >= mh.length { // 已经没有右孩子结点了
		// 比左孩子结点大，进行交换
		if mh.data[index] > mh.data[left] {
			mh.data[index], mh.data[left] = mh.data[left], mh.data[index]
			return left, true
		}
		return index, false
	}
	// 两个孩子结点都存在，从中选一个较小者
	var smaller int
	// false 左孩子结点小， true 右孩子结点小
	var flag bool
	if mh.data[left] < mh.data[right] {
		smaller = mh.data[left]
	} else {
		smaller = mh.data[right]
		flag = true
	}
	// 与较小的孩子结点进行比较，如果比孩子结点大则进行交换
	if mh.data[index] < smaller {
		if flag {
			// 右孩子结点小，和右孩子结点进行交换
			mh.data[index], mh.data[right] = mh.data[right], mh.data[index]
			return right, true
		}
		// 左孩子结点小，和左孩子结点进行交换
		mh.data[index], mh.data[left] = mh.data[left], mh.data[index]
		return left, true
	}

	return index, false
}

// NewMinHeap 创建一个数据为空的大顶堤
func NewMinHeap(size int) *minHeap {
	return &minHeap{
		length: 0,
		size:   size,
		data:   make([]int, size),
	}
}

// BuildMinHeap 根据给定的数组创建大顶堆
// 从数组中的第二个元素开始进行Push操作
func BuildMinHeap(data []int) *minHeap {
	length := len(data)
	mh := &minHeap{
		length: length,
		size:   length,
		data:   data,
	}

	mh.mu.Lock()
	defer mh.mu.Unlock()

	for i := 1; i < length; i++ {
		for j := i; j >= 0; j = (j - 1) / 2 {
			if !mh.shiftUp(j) {
				break
			}
		}
	}

	return mh
}
