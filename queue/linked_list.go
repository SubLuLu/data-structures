package queue

import (
	"errors"
	"sync"
)

type (
	// 链表结点
	linkedListNode struct {
		data int             // 结点元素值
		next *linkedListNode // 后继结点
	}

	// 使用链表的方式实现队列
	//
	// 添加一个size属性
	// size主要用途是在遍历的时候申请长度为size的切片
	// 避免使用append函数，导致频繁分配内存
	linkedListQueue struct {
		mu   sync.RWMutex    // 读写锁
		size int             // 队列的当前大小
		head *linkedListNode // 链表头结点(非首结点)
		tail *linkedListNode // 链表尾结点
	}
)

// Enqueue 向队列中添加元素
// num 等待添加的元素
// 返回值为添加元素过程中的错误信息
func (lq *linkedListQueue) Enqueue(num int) error {
	node := &linkedListNode{
		data: num,
	}

	lq.mu.Lock()
	defer lq.mu.Unlock()

	lq.tail.next = node
	lq.tail = node
	lq.size++
	return nil
}

// Dequeue 从队列中弹出元素
// 返回值为队列中的头部元素或者错误信息
func (lq *linkedListQueue) Dequeue() (int, error) {
	if lq.IsEmpty() {
		return 0, errors.New("link list queue is empty")
	}

	lq.mu.Lock()
	defer lq.mu.Unlock()

	node := lq.head.next // 链表首结点
	if node == lq.tail { // 删除最后一个结点
		lq.tail = lq.head // 尾结点指向头结点
	}
	lq.head.next = node.next
	lq.size--
	return node.data, nil
}

// IsEmpty 判断队列是否为空
// 返回值表示队列是否为空 true表示为空
func (lq *linkedListQueue) IsEmpty() bool {
	lq.mu.RLock()
	defer lq.mu.RUnlock()

	// 1. 通过size判断队列是否为空
	// return lq.size == 0

	// 2. 通过头结点和尾结点判断队列是否为空
	return lq.head == lq.tail
}

// Traverse 遍历队列，以切片形式返回
func (lq *linkedListQueue) Traverse() []int {
	lq.mu.RLock()
	defer lq.mu.RUnlock()

	res := make([]int, lq.size)
	t := lq.head
	for i := 0; t != lq.tail; i++ {
		t = t.next
		res[i] = t.data
	}
	return res
}

// NewLinkListQueue 创建一个新的链表队列，返回其指针
func NewLinkListQueue() *linkedListQueue {
	// 声明一个空节点
	node := &linkedListNode{
	}
	// 初始化一个链表队列
	return &linkedListQueue{
		head: node,
		tail: node,
	}
}
