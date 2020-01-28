package stack

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

	// 使用链表的方式实现栈
	//
	// 添加一个size属性
	// size主要用途是在遍历的时候申请长度为size的切片
	// 避免使用append函数，导致频繁分配内存
	linkedListStack struct {
		mu     sync.RWMutex    // 读写锁
		size   int             // 队列的当前大小
		top    *linkedListNode // 栈顶结点
		bottom *linkedListNode // 栈底结点
	}
)

// Push 压栈，向栈中添加元素
func (ls *linkedListStack) Push(val int) error {
	// 构建新结点
	node := &linkedListNode{
		data: val,
	}

	ls.mu.Lock()
	defer ls.mu.Unlock()

	// 将新结点的next指向原来的栈顶
	node.next = ls.top
	// 将栈顶更新为新结点
	ls.top = node
	// 栈size增加
	ls.size++
	return nil
}

// Pop 出栈，从栈顶弹出元素
func (ls *linkedListStack) Pop() (int, error) {
	if ls.IsEmpty() { // 栈为空，不能弹出元素
		return 0, errors.New("link list stack is empty")
	}

	ls.mu.Lock()
	defer ls.mu.Unlock()

	// 存储栈顶数据
	val := ls.top.data
	// 更新栈顶我原来栈顶的next
	ls.top = ls.top.next
	// 栈size减小
	ls.size--
	return val, nil
}

// IsEmpty 判断栈是否为空
// 返回值表示栈是否为空 true表示为空
func (ls *linkedListStack) IsEmpty() bool {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	// 1. 通过size判断栈是否为空
	// return ls.size == 0

	// 2. 通过栈顶和占地判断是否为空
	return ls.bottom == ls.top
}

// Traverse 遍历栈，以切片形式返回
func (ls *linkedListStack) Traverse() []int {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	res := make([]int, ls.size)
	t := ls.top // 临时存储栈顶
	for i := 0; t != ls.bottom; i++ {
		res[i] = t.data
		t = t.next
	}
	return res
}

// NewLinkListStack 创建一个新的链表栈，返回其指针
func NewLinkListStack() *linkedListStack {
	node := &linkedListNode{}
	return &linkedListStack{
		top:    node,
		bottom: node,
	}
}
