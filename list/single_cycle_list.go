package list

import (
	"errors"
)

// 单向循环链表
type (
	// 单向循环链表结点
	singleCycleNode struct {
		data int              // 结点元素值
		next *singleCycleNode // 后继结点
	}

	// 单向循环链表
	singleCycleList struct {
		size int              // 结点元素个数
		head *singleCycleNode // 头指针
		tail *singleCycleNode // 尾指针
	}
)

// Add 在单向循环链表头部添加结点
// 通过头指针确定首结点
// 在头结点和首结点之间添加新的结点
func (scl *singleCycleList) Add(val int) error {
	// 构建一个新结点
	node := &singleCycleNode{
		data: val,
	}
	// 确定首结点
	first := scl.head.next
	// 直接插入到头结点后面
	scl.head.next = node
	// 如果首结点为空
	if first == nil {
		// 将尾指针指向新的结点
		scl.tail = node
		// 新结点的指针域指向头结点形成环
		node.next = scl.head
	} else {
		// 新的结点的指针域指向之前的首结点
		node.next = first
	}
	// 单向循环链表size增加
	scl.size++
	return nil
}

// Append 在单向循环链表末尾添加结点
// 从头结点开始遍历，找到尾结点
// 在尾结点后面添加新的结点
func (scl *singleCycleList) Append(val int) error {
	// 构建一个新结点
	node := &singleCycleNode{
		data: val,
	}
	// 如果尾指针为空，从头开始插入
	if scl.tail == nil {
		// 头结点指针域指向新结点
		scl.head.next = node
	} else {
		// 尾结点的指针域指向新结点
		scl.tail.next = node
	}
	// 将尾结点的next指向新结点
	scl.tail = node
	// 将新结点的next指向新结点
	node.next = scl.head
	// 单向循环链表size增加
	scl.size++
	return nil
}

// findPrevNode 找到指定位置的前一个结点
// 因为单向循环链表的操作，需要借助前驱结点
// 所以提供一个查找指定位置结点的前驱结点的方法
func (scl *singleCycleList) findPrevNode(index int) *singleCycleNode {
	// index无效，返回空
	if index > scl.size || index < 0 {
		return nil
	}
	t := scl.head // 临时变量指向头结点
	// 从头结点开始遍历找到插入位置的前一个结点
	// 如果index为0，则t指向头结点
	// 否则t指向第index-1个有效结点
	for i := 0; i < index; i++ {
		t = t.next
	}
	return t
}

// Insert 在单向循环链表指定位置插入新的结点
// index 指定的位置，从0开始
func (scl *singleCycleList) Insert(val, index int) error {
	// index超出范围，返回错误
	if index > scl.size || index < 0 {
		return errors.New("insert fail, index out of range")
	}
	// 在最后添加，直接调用Append方法
	if index == scl.size {
		return scl.Append(val)
	}
	// 查找前驱结点
	prev := scl.findPrevNode(index)
	// 构建新结点
	node := &singleCycleNode{
		data: val,
	}
	// 新结点的next指向index位置的结点
	node.next = prev.next
	// index位置前面结点的next指向新结点
	prev.next = node
	// 链表size增加
	scl.size++
	return nil
}

// Delete 删除指定位置的元素结点，并返回被删除的元素值
// index 指定的位置，从0开始
func (scl *singleCycleList) Delete(index int) (int, error) {
	// index超出范围，返回错误
	if index >= scl.size || index < 0 {
		return 0, errors.New("delete fail, index out of range")
	}
	// 存储需要返回的值和计数器
	var val int
	// 查找前驱结点
	prev := scl.findPrevNode(index)
	current := prev.next // 需要删除的结点
	val = current.data // 取出index位置结点的元素值
	// 把index位置的前一个结点的next指向index位置的后一个结点
	prev.next = current.next
	// 断开删除结点的指针域
	current.next = nil
	// 链表size减少
	scl.size--
	return val, nil
}

// Set 给指定位置的元素重新赋值
// index 指定的位置，从0开始
func (scl *singleCycleList) Set(val, index int) error {
	// index超出范围，返回错误
	if index >= scl.size || index < 0 {
		return errors.New("set fail, index out of range")
	}
	// 找到index位置的结点
	current := scl.findPrevNode(index).next
	// 重新赋值
	current.data = val
	return nil
}

// Find 查找指定位置结点的元素值
// index 指定的位置，从0开始
func (scl *singleCycleList) Find(index int) (int, error) {
	// index超出范围，返回错误
	if index >= scl.size || index < 0 {
		return 0, errors.New("find fail, index out of range")
	}
	// 找到index位置的结点
	current := scl.findPrevNode(index).next
	return current.data, nil
}

// Traverse 遍历单向循环链表，以切片形式返回
func (scl *singleCycleList) Traverse() []int {
	res := make([]int, scl.size)
	t := scl.head
	for i := 0; t.next != scl.head; i++ {
		t = t.next
		res[i] = t.data
	}
	return res
}

// Cycle 循环遍历单向循环链表，以切片形式返回
// num 为循环遍历圈数
func (scl *singleCycleList) Cycle(num int) []int {
	total := scl.size * num
	res := make([]int, total)
	t := scl.head
	for i := 0; i < total; i++ {
		if t == scl.head {
			t = t.next
		}
		res[i] = t.data
		t = t.next
	}
	return res
}

// Length 返回单向循环链表的长度(结点元素个数)
func (scl *singleCycleList) Length() int {
	// 1. 通过size返回
	return scl.size
	// 2. 通过遍历链表计算
	// var length int
	// temp := scl.head
	// for length = 0; temp.next != nil length++ {
	// 	temp = temp.next
	// }
	// return length
}

// IsEmpty 判断单向循环链表是否为空
func (scl *singleCycleList) IsEmpty() bool {
	// 1. 通过size进行判断
	// return scl.size == 0
	// 2. 通过head进行判断
	return scl.head.next == scl.head
	// 3. 通过tail进行判断
	// return sl.tail == nil
}

// NewSingleCycleList 创建一个单向循环链表，返回其指针
func NewSingleCycleList() *singleCycleList {
	// 头指针
	node := &singleCycleNode{}
	// 头结点的指针域指向自己
	node.next = node
	return &singleCycleList{
		head: node,
	}
}
