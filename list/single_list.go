package list

import (
	"errors"
	"sync"
)

// 单链表
type (
	// 单链表结点
	singleNode struct {
		data int         // 结点元素值(数据域)
		next *singleNode // 后继结点(指针域)
	}

	// 单链表
	singleList struct {
		mu   sync.RWMutex // 读写锁
		size int          // 单链表中结点元素个数
		head *singleNode  // 单链表中的头指针
		tail *singleNode  // 单链表中的尾指针
	}
)

// Add 在单链表头部添加结点
// 通过头指针确定首结点
// 在头结点和首结点之间添加新的结点
func (sl *singleList) Add(val int) error {
	// 构建一个新结点
	node := &singleNode{
		data: val,
	}

	sl.mu.Lock()
	defer sl.mu.Unlock()

	// 确定首结点
	first := sl.head.next
	// 直接插入到头结点后面
	sl.head.next = node
	// 如果首结点为空
	if first == nil {
		// 将尾指针指向新的结点
		sl.tail = node
	} else {
		// 新的结点的指针域指向之前的首结点
		node.next = first
	}
	// 单链表size增加
	sl.size++
	return nil
}

// Append 在单链表末尾添加结点
// 通过尾指针确定尾结点
// 在尾结点后面添加新的结点
func (sl *singleList) Append(val int) error {
	// 构建一个新结点
	node := &singleNode{
		data: val,
	}

	sl.mu.Lock()
	defer sl.mu.Unlock()

	// 如果尾指针为空，从头开始插入
	if sl.tail == nil {
		// 头结点指针域指向新结点
		sl.head.next = node
	} else {
		// 尾结点的指针域指向新结点
		sl.tail.next = node
	}
	// 尾指针指向新结点
	sl.tail = node
	// 单链表size增加
	sl.size++
	return nil
}

// findPrevNode 找到指定位置的前一个结点
// 因为单链表的操作，需要借助前驱结点
// 所以提供一个查找指定位置结点的前驱结点的方法
func (sl *singleList) findPrevNode(index int) *singleNode {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	// index无效，返回空
	if index > sl.size || index < 0 {
		return nil
	}
	t := sl.head // 临时变量指向头结点
	// 从头结点开始遍历找到插入位置的前一个结点
	// 如果index为0，则t指向头结点
	// 否则t指向第index-1个有效结点
	for i := 0; i < index; i++ {
		t = t.next
	}
	return t
}

// Insert 在单链表指定位置插入新的结点
// index 指定的位置，从0开始
func (sl *singleList) Insert(val, index int) error {
	sl.mu.RLock()
	size := sl.size
	sl.mu.RUnlock()

	// index超出范围，返回错误
	if index > size || index < 0 {
		return errors.New("insert fail, index out of range")
	}
	// 在最后添加，直接调用Append方法
	if index == size {
		return sl.Append(val)
	}

	// 查找前驱结点
	prev := sl.findPrevNode(index)

	// 构建新结点
	node := &singleNode{
		data: val,
	}

	sl.mu.Lock()
	defer sl.mu.Unlock()

	// 新结点的next指向index位置的结点
	node.next = prev.next
	// index位置前面结点的next指向新结点
	prev.next = node
	// 链表size增加
	sl.size++
	return nil
}

// Delete 删除指定位置的元素结点，并返回被删除的元素值
// index 指定的位置，从0开始
func (sl *singleList) Delete(index int) (int, error) {
	sl.mu.RLock()
	size := sl.size
	sl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= size || index < 0 {
		return 0, errors.New("delete fail, index out of range")
	}
	// 找到前驱结点，prev.next是index位置的结点
	prev := sl.findPrevNode(index)

	sl.mu.Lock()
	defer sl.mu.Unlock()

	val := prev.next.data     // 取出index位置结点的元素值
	if prev.next == sl.tail { // 如果删除的是最后一个结点
		// 直接将index位置的前一个结点的next赋值nil
		prev.next = nil
		// 尾指针指向最后一个结点
		sl.tail = prev
	} else { // 如果删除的不是最后一个结点
		// 需要删除的结点
		current := prev.next
		// 把index位置的前一个结点的next指向index位置的后一个结点
		prev.next = current.next
		// 断开删除结点的指针域
		current.next = nil
	}
	// 链表size减少
	sl.size--
	return val, nil
}

// Set 给指定位置的元素重新赋值
// index 指定的位置，从0开始
func (sl *singleList) Set(val, index int) error {
	sl.mu.RLock()
	size := sl.size
	sl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= size || index < 0 {
		return errors.New("set fail, index out of range")
	}
	// 找到index位置的结点
	current := sl.findPrevNode(index).next
	sl.mu.Lock()
	// 重新赋值
	current.data = val
	sl.mu.Unlock()

	return nil
}

// Find 查找指定位置结点的元素值
// index 指定的位置，从0开始
func (sl *singleList) Find(index int) (int, error) {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= sl.size || index < 0 {
		return 0, errors.New("find fail, index out of range")
	}
	// 找到index位置的结点
	current := sl.findPrevNode(index).next
	return current.data, nil
}

// Traverse 遍历单链表，以切片形式返回
func (sl *singleList) Traverse() []int {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	// 指定[]int的大小，避免频繁分配内存
	res := make([]int, sl.size)
	t := sl.head
	for i := 0; t.next != nil; i++ {
		t = t.next
		res[i] = t.data
	}
	return res
}

// Length 返回单链表的长度(结点元素个数)
func (sl *singleList) Length() int {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	// 1. 通过size返回
	return sl.size

	// 2. 通过遍历链表计算
	// var length int
	// temp := sl.head
	// for length = 0; temp.next != nil length++ {
	// 	temp = temp.next
	// }
	// return length
}

// IsEmpty 判断单链表是否为空
func (sl *singleList) IsEmpty() bool {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	// 1. 通过size进行判断
	// return sl.size == 0

	// 2. 通过head进行判断
	return sl.head.next == nil

	// 3. 通过tail进行判断
	// return sl.tail == nil
}

// NewSingleList 创建一个单链表，返回其指针
func NewSingleList() *singleList {
	return &singleList{
		head: &singleNode{},
	}
}
