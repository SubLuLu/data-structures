package list

import (
	"errors"
	"sync"
)

// 双向循环链表
type (
	// 双向循环链表结点
	doubleCycleNode struct {
		data  int              // 结点元素值(数据域)
		prior *doubleCycleNode // 前驱结点(指针域)
		next  *doubleCycleNode // 后继结点(指针域)
	}

	// 双向循环链表
	doubleCycleList struct {
		mu   sync.RWMutex     // 读写锁
		size int              // 结点元素个数
		head *doubleCycleNode // 头指针
	}
)

// Add 在双向循环链表头部添加结点
// 通过头指针确定首结点
// 在头结点和首结点之间添加新的结点
func (dcl *doubleCycleList) Add(val int) error {
	// 构建一个新结点
	node := &doubleCycleNode{
		data: val,
	}

	dcl.mu.Lock()
	defer dcl.mu.Unlock()

	// 确定首结点
	first := dcl.head.next
	// 直接插入到头结点后面
	dcl.head.next = node
	// 新结点的前驱指针域指向头结点
	node.prior = dcl.head
	// 新的结点的后继指针域指向之前的首结点
	node.next = first
	// 之前的首结点的前驱指针域指向新结点
	first.prior = node
	// 双向循环链表size增加
	dcl.size++
	return nil
}

// Append 在双向循环链表末尾添加结点
// 利用头结点的前驱结点，找到尾结点
// 在尾结点后面添加新的结点
func (dcl *doubleCycleList) Append(val int) error {
	// 创建新结点
	node := &doubleCycleNode{
		data: val,
	}

	dcl.mu.Lock()
	defer dcl.mu.Unlock()

	tail := dcl.head.prior // 尾指针
	// 尾结点的后继指针域指向新结点
	tail.next = node
	// 新结点的前驱指针域指向尾结点
	node.prior = tail
	// 新结点的后继指针域指向头结点
	node.next = dcl.head
	// 头结点的前驱指针域指向新结点
	dcl.head.prior = node
	// 双向循环链表size增加
	dcl.size++
	return nil
}

// findNode 找到指定位置的结点
// 因为双向循环链表对结点的操作可以通过自身完成
// 所以提供一个查找指定位置结点方法
func (dcl *doubleCycleList) findNode(index int) *doubleCycleNode {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	// index无效，返回空
	if index > dcl.size || index < 0 {
		return nil
	}
	t := dcl.head // 临时变量指向头结点
	// 从头结点开始遍历找到插入位置的结点
	// 如果index为0，则t指向首结点
	// 否则t指向第index个有效结点
	for i := 0; i <= index; i++ {
		t = t.next
	}
	return t
}

// Insert 在双向链表指定位置插入新的结点
// index 指定的位置，从0开始
func (dcl *doubleCycleList) Insert(val, index int) error {
	dcl.mu.RLock()
	size := dcl.size
	dcl.mu.RUnlock()

	// index超出范围，返回错误
	if index > size || index < 0 {
		return errors.New("insert fail, index out of range")
	}
	// 在最后添加，直接调用Append方法
	if index == size {
		return dcl.Append(val)
	}
	// 找到index位置的结点
	current := dcl.findNode(index)

	dcl.mu.Lock()
	defer dcl.mu.Unlock()

	// 构建新结点
	node := &doubleCycleNode{
		data: val,
	}
	// 将新结点的后继指针域指向原来index位置的结点
	node.next = current
	// 将新结点的前驱指针域指向原来index-1位置的结点
	node.prior = current.prior
	// 将原来index-1位置的结点的后继指针域指向新结点
	current.prior.next = node
	// 将原来index位置的结点的前驱指针域指向新结点
	current.prior = node
	// 双向循环链表size增加
	dcl.size++
	return nil
}

// Delete 删除指定位置的元素结点，并返回被删除的元素值
// index 指定的位置，从0开始
func (dcl *doubleCycleList) Delete(index int) (int, error) {
	dcl.mu.RLock()
	size := dcl.size
	dcl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= size || index < 0 {
		return 0, errors.New("delete fail, index out of range")
	}
	// 待删除结点
	current := dcl.findNode(index)
	return dcl.DeleteSelf(current)
}

// DeleteSelf 删除指定的结点，并返回被删除的元素值
func (dcl *doubleCycleList) DeleteSelf(node *doubleCycleNode) (int, error) {
	dcl.mu.Lock()
	defer dcl.mu.Unlock()

	// 前驱结点或后继结点为空，则不在双向循环链表中
	if node.prior == nil || node.next == nil {
		return 0, errors.New("delete self fail, illegal node")
	}
	// 将删除结点的前一个结点的next指向删除结点的下一个结点
	node.prior.next = node.next
	// 将删除结点的下一个结点的prior指向删除结点的前一个结点
	node.next.prior = node.prior
	// 清空node的前驱指针域
	node.prior = nil
	// 清空node的后继指针域
	node.next = nil
	// 双向循环链表size减少
	dcl.size--
	return node.data, nil
}

// Set 给指定位置的元素重新赋值
// index 指定的位置，从0开始
func (dcl *doubleCycleList) Set(val, index int) error {
	dcl.mu.RLock()
	size := dcl.size
	dcl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= size || index < 0 {
		return errors.New("set fail, index out of range")
	}
	// 待修改的结点
	current := dcl.findNode(index)
	dcl.mu.Lock()
	defer dcl.mu.Unlock()
	current.data = val // 重新赋值
	return nil
}

// Find 查找指定位置结点的元素值
// index 指定的位置，从0开始
func (dcl *doubleCycleList) Find(index int) (int, error) {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= dcl.size || index < 0 {
		return 0, errors.New("find fail, index out of range")
	}
	current := dcl.findNode(index)
	return current.data, nil
}

// Traverse 遍历双向链表，以切片形式返回
func (dcl *doubleCycleList) Traverse() []int {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	res := make([]int, dcl.size)
	t := dcl.head
	for i := 0; t.next != dcl.head; i++ {
		t = t.next
		res[i] = t.data
	}
	return res
}

// Reverse 逆向遍历双向循环链表，以切片形式返回
func (dcl *doubleCycleList) Reverse() []int {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	res := make([]int, dcl.size)
	t := dcl.head.prior // 尾指针
	for i := 0; t != dcl.head; i++ {
		res[i] = t.data
		t = t.prior
	}
	return res
}

// Cycle 循环遍历双向循环链表，以切片形式返回
// num 为循环遍历圈数
func (dcl *doubleCycleList) Cycle(num int) []int {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	total := dcl.size * num
	res := make([]int, total)
	t := dcl.head
	for i := 0; i < total; i++ {
		if t == dcl.head {
			t = t.next
		}
		res[i] = t.data
		t = t.next
	}
	return res
}

// Length 返回双向链表的长度(结点元素个数)
func (dcl *doubleCycleList) Length() int {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	// 1. 通过size返回
	return dcl.size

	// 2. 通过遍历链表计算
	// var length int
	// temp := dcl.head
	// for length = 0; temp.next != nil length++ {
	// 	temp = temp.next
	// }
	// return length
}

// IsEmpty 判断双向链表是否为空
func (dcl *doubleCycleList) IsEmpty() bool {
	dcl.mu.RLock()
	defer dcl.mu.RUnlock()

	// 1. 通过size进行判断
	// return dcl.size == 0
	// 2. 通过head进行判断
	return dcl.head.next == dcl.head
}

// NewDoubleCycleList 创建一个双向循环链表，返回其指针
func NewDoubleCycleList() *doubleCycleList {
	node := &doubleCycleNode{}
	node.next = node  // 后继指针域指向自身
	node.prior = node // 前驱指针域指向自身
	return &doubleCycleList{
		head: node,
	}
}
