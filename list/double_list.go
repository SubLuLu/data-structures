package list

import (
	"errors"
	"sync"
)

// 双链表
type (
	// 双向链表结点
	doubleNode struct {
		data  int         // 结点元素值(数据域)
		prior *doubleNode // 前驱结点(指针域)
		next  *doubleNode // 后继结点(指针域)
	}

	// 双向链表
	doubleList struct {
		mu   sync.RWMutex // 读写锁
		size int          // 双链表中结点元素个数
		head *doubleNode  // 双链表中的头指针
		tail *doubleNode  // 双链表中的尾指针
	}
)

// Add 在双向链表头部添加结点
// 通过头指针确定首结点
// 在头结点和首结点之间添加新的结点
func (dl *doubleList) Add(val int) error {
	// 构建一个新结点
	node := &doubleNode{
		data: val,
	}

	dl.mu.Lock()
	defer dl.mu.Unlock()

	// 确定首结点
	first := dl.head.next
	// 直接插入到头结点后面
	dl.head.next = node
	// 新结点的前驱指针域指向头结点
	node.prior = dl.head
	// 如果首结点为空
	if first == nil {
		// 将尾指针指向新的结点
		dl.tail = node
	} else {
		// 新的结点的后继指针域指向之前的首结点
		node.next = first
		// 之前的首结点的前驱指针域指向新结点
		first.prior = node
	}
	// 双链表size增加
	dl.size++
	return nil
}

// Append 在双向链表末尾添加结点
// 从头结点开始遍历，找到尾结点
// 在尾结点后面添加新的结点
func (dl *doubleList) Append(val int) error {
	// 创建新结点
	node := &doubleNode{
		data: val,
	}

	dl.mu.Lock()
	defer dl.mu.Unlock()

	// 如果尾指针为空，从头开始插入
	if dl.tail == nil {
		// 头结点后继指针域指向新结点
		dl.head.next = node
		// 新结点的前驱指针域指向头结点
		node.prior = dl.head
	} else {
		// 尾结点的后继指针域指向新结点
		dl.tail.next = node
		// 新结点的前驱指针域指向尾结点
		node.prior = dl.tail
	}
	// 新结点成为新的尾结点
	dl.tail = node
	// 双链表size增加
	dl.size++
	return nil
}

// findNode 找到指定位置的结点
// 因为双链表对结点的操作可以通过自身完成
// 所以提供一个查找指定位置结点方法
func (dl *doubleList) findNode(index int) *doubleNode {
	dl.mu.RLock()
	defer dl.mu.RUnlock()

	// index无效，返回空
	if index > dl.size || index < 0 {
		return nil
	}
	t := dl.head // 临时变量指向头结点
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
func (dl *doubleList) Insert(val, index int) error {
	dl.mu.RLock()
	size := dl.size
	dl.mu.RUnlock()

	// index超出范围，返回错误
	if index > size || index < 0 {
		return errors.New("insert fail, index out of range")
	}
	// 在最后添加，直接调用Append方法
	if index == size {
		return dl.Append(val)
	}
	// 待插入位置的结点
	current := dl.findNode(index)
	// 构建新结点
	node := &doubleNode{
		data: val,
	}

	dl.mu.Lock()
	defer dl.mu.Unlock()

	// 新结点的后继指针域指向原来index位置的结点
	node.next = current
	// 新结点的前驱指针域指向原来index-1位置的结点
	node.prior = current.prior
	// 将index-1位置原来的结点的后继指针域指向新结点
	current.prior.next = node
	// 将index位置原来的结点的前驱指针域指向新结点
	current.prior = node
	// 双链表size增加
	dl.size++
	return nil
}

// Delete 删除指定位置的元素结点，并返回被删除的元素值
// index 指定的位置，从0开始
func (dl *doubleList) Delete(index int) (int, error) {
	dl.mu.RLock()
	size := dl.size
	dl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= size || index < 0 {
		return 0, errors.New("delete fail, index out of range")
	}
	// 待删除结点
	current := dl.findNode(index)
	return dl.DeleteSelf(current)
}

// DeleteSelf 删除指定的结点，并返回被删除的元素值
func (dl *doubleList) DeleteSelf(node *doubleNode) (int, error) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	// 前驱结点为空，则不在双向链表中
	if node.prior == nil {
		return 0, errors.New("delete self fail, illegal node")
	}
	if node == dl.tail { // 删除最后一个结点
		// 将要删除结点的前一个结点的next指向nil
		node.prior.next = nil
		if node.prior == dl.head { // 删除的是首结点
			// 尾指针置为nil
			dl.tail = nil
		} else {
			// 尾指针指向删除结点的前驱结点
			dl.tail = node.prior
		}
	} else {
		// 将删除结点的前一个结点的后继指针域指向删除结点的下一个结点
		node.prior.next = node.next
		// 将删除结点的下一个结点的前驱指针域指向删除结点的前一个结点
		node.next.prior = node.prior
	}
	// 清空node的前驱指针域
	node.prior = nil
	// 清空node的后继指针域
	node.next = nil
	// 双链表size减少
	dl.size--
	return node.data, nil
}

// Set 给指定位置的元素重新赋值
// index 指定的位置，从0开始
func (dl *doubleList) Set(val, index int) error {
	dl.mu.RLock()
	size := dl.size
	dl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= size || index < 0 {
		return errors.New("set fail, index out of range")
	}
	// 待修改的结点
	current := dl.findNode(index)

	dl.mu.Lock()
	current.data = val // 重新赋值
	dl.mu.Unlock()
	return nil
}

// Find 查找指定位置结点的元素值
// index 指定的位置，从0开始
func (dl *doubleList) Find(index int) (int, error) {
	dl.mu.RLock()
	defer dl.mu.RUnlock()

	// index超出范围，返回错误
	if index >= dl.size || index < 0 {
		return 0, errors.New("find fail, index out of range")
	}
	current := dl.findNode(index)
	return current.data, nil
}

// Traverse 遍历双向链表，以切片形式返回
func (dl *doubleList) Traverse() []int {
	dl.mu.RLock()
	defer dl.mu.RUnlock()

	res := make([]int, dl.size)
	t := dl.head
	for i := 0; t.next != nil; i++ {
		t = t.next
		res[i] = t.data
	}
	return res
}

// Reverse 逆向遍历双向链表，以切片形式返回
func (dl *doubleList) Reverse() []int {
	dl.mu.RLock()
	defer dl.mu.RUnlock()

	res := make([]int, dl.size)
	t := dl.tail
	for i := 0; t != dl.head; i++ {
		res[i] = t.data
		t = t.prior
	}
	return res
}

// Length 返回双向链表的长度(结点元素个数)
func (dl *doubleList) Length() int {
	dl.mu.RLock()
	defer dl.mu.RUnlock()

	// 1. 通过size返回
	return dl.size

	// 2. 通过遍历链表计算
	// var length int
	// temp := dl.head
	// for length = 0; temp.next != nil length++ {
	// 	temp = temp.next
	// }
	// return length
}

// IsEmpty 判断双向链表是否为空
func (dl *doubleList) IsEmpty() bool {
	dl.mu.RLock()
	defer dl.mu.RUnlock()

	// 1. 通过size进行判断
	// return dl.size == 0
	// 2. 通过head进行判断
	return dl.head.next == nil
	// 3. 通过判断尾指针是否为空
	// return dl.tail == nil
}

// NewDoubleList 创建一个双向链表，返回其指针
func NewDoubleList() *doubleList {
	return &doubleList{
		head: &doubleNode{},
	}
}
