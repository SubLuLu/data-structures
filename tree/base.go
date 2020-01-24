package tree

import (
	"bytes"
	"errors"
	"strconv"
)

type (
	listNode struct {
		data int
		next *listNode
	}

	list struct { // 为了方便插入操作添加了tail
		head *listNode
		tail *listNode
	}
)

func (l *list) append(val int) {
	node := newListNode(val)
	if l.tail == nil {
		l.head.next = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (l *list) string() string {
	t := l.head.next
	buf := bytes.Buffer{}
	buf.WriteString("[ ")
	for t != nil {
		buf.WriteString(strconv.Itoa(t.data))
		buf.WriteString(" ")
		t = t.next
	}
	buf.WriteString("]")
	return buf.String()
}

func newListNode(val int) *listNode {
	return &listNode{
		data: val,
		next: nil,
	}
}

func newList() *list {
	node := newListNode(0)
	return &list{
		head: node,
		tail: nil,
	}
}

type (
	// 链表结点
	linkListNode struct {
		data int           // 结点元素值
		next *linkListNode // 后继结点
	}

	// 使用链表的方式实现栈
	//
	// 添加一个size属性
	// size主要用途是在遍历的时候申请长度为size的切片
	// 避免使用append函数，导致频繁分配内存
	linkListStack struct {
		size   int           // 队列的当前大小
		top    *linkListNode // 栈顶结点
		bottom *linkListNode // 栈底结点
	}
)

// Push 压栈，向栈中添加元素
func (ls *linkListStack) Push(val int) error {
	// 构建新结点
	node := &linkListNode{
		data: val,
	}
	// 将新结点的next指向原来的栈顶
	node.next = ls.top
	// 将栈顶更新为新结点
	ls.top = node
	// 栈size增加
	ls.size++
	return nil
}

// Pop 出栈，从栈顶弹出元素
func (ls *linkListStack) Pop() (int, error) {
	if ls.IsEmpty() { // 栈为空，不能弹出元素
		return 0, errors.New("link list stack is empty")
	}
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
func (ls *linkListStack) IsEmpty() bool {
	// 1. 通过size判断栈是否为空
	// return ls.size == 0
	// 2. 通过栈顶和占地判断是否为空
	return ls.bottom == ls.top
}

// NewLinkListStack 创建一个新的链表栈，返回其指针
func NewLinkListStack() *linkListStack {
	node := &linkListNode{}
	return &linkListStack{
		top:    node,
		bottom: node,
	}
}

// 定义一个存放二叉树结点指针的栈
type (
	// 链表结点
	stackNode struct {
		data *listBinary // 结点元素值
		next *stackNode  // 后继结点
	}

	// 使用链表的方式实现栈
	nodeStack struct {
		top    *stackNode // 栈顶结点
		bottom *stackNode // 栈底结点
	}
)

// Push 压栈，向栈中添加元素
func (ns *nodeStack) Push(val *listBinary) error {
	// 构建新结点
	node := &stackNode{
		data: val,
	}
	// 将新结点的next指向原来的栈顶
	node.next = ns.top
	// 将栈顶更新为新结点
	ns.top = node
	return nil
}

// Pop 出栈，从栈顶弹出元素
func (ns *nodeStack) Pop() (*listBinary, error) {
	if ns.IsEmpty() { // 栈为空，不能弹出元素
		return nil, errors.New("link list stack is empty")
	}
	// 存储栈顶数据
	val := ns.top.data
	// 更新栈顶我原来栈顶的next
	ns.top = ns.top.next
	return val, nil
}

// IsEmpty 判断栈是否为空
// 返回值表示栈是否为空 true表示为空
func (ns *nodeStack) IsEmpty() bool {
	// 通过栈顶和占地判断是否为空
	return ns.bottom == ns.top
}

// NewNodeStack 创建一个新的链表栈，返回其指针
func NewNodeStack() *nodeStack {
	node := &stackNode{}
	return &nodeStack{
		top:    node,
		bottom: node,
	}
}

// 简易链表结构队列
type (
	// 链表结点
	queueNode struct {
		data *listBinary // 结点元素值
		next *queueNode  // 后继结点
	}

	// 使用链表的方式实现队列
	queue struct {
		head *queueNode // 链表头结点(非首结点)
		tail *queueNode // 链表尾结点
	}
)

// Enqueue 向队列中添加元素
// num 等待添加的元素
// 返回值为添加元素过程中的错误信息
func (lq *queue) Enqueue(lb *listBinary) error {
	node := &queueNode{
		data: lb,
	}
	lq.tail.next = node
	lq.tail = node
	return nil
}

// Dequeue 从队列中弹出元素
// 返回值为队列中的头部元素或者错误信息
func (lq *queue) Dequeue() (*listBinary, error) {
	if lq.IsEmpty() {
		return nil, errors.New("link list queue is empty")
	}
	node := lq.head.next // 链表首结点
	if node == lq.tail { // 删除最后一个结点
		lq.tail = lq.head // 尾结点指向头结点
	}
	lq.head.next = node.next
	return node.data, nil
}

// IsEmpty 判断队列是否为空
// 返回值表示队列是否为空 true表示为空
func (lq *queue) IsEmpty() bool {
	// 通过头结点和尾结点判断队列是否为空
	return lq.head == lq.tail
}

// NewQueue 创建一个新的链表队列，返回其指针
func NewQueue() *queue {
	// 声明一个空节点
	node := &queueNode{}
	// 初始化一个链表队列
	return &queue{
		head: node,
		tail: node,
	}
}

// 后序线索化遍历使用的栈
type (
	threadNoe struct {
		data *threadedNode
		next *threadNoe
	}

	threadStack struct {
		top    *threadNoe // 栈顶结点
		bottom *threadNoe // 栈底结点
	}
)

// Push 压栈，向栈中添加元素
func (ns *threadStack) Push(val *threadedNode) error {
	// 构建新结点
	node := &threadNoe{
		data: val,
	}
	// 将新结点的next指向原来的栈顶
	node.next = ns.top
	// 将栈顶更新为新结点
	ns.top = node
	return nil
}

// Pop 出栈，从栈顶弹出元素
func (ns *threadStack) Pop() (*threadedNode, error) {
	if ns.IsEmpty() { // 栈为空，不能弹出元素
		return nil, errors.New("link list stack is empty")
	}
	// 存储栈顶数据
	val := ns.top.data
	// 更新栈顶我原来栈顶的next
	ns.top = ns.top.next
	return val, nil
}

// IsEmpty 判断栈是否为空
// 返回值表示栈是否为空 true表示为空
func (ns *threadStack) IsEmpty() bool {
	// 通过栈顶和占地判断是否为空
	return ns.bottom == ns.top
}

func (ns *threadStack) Traverse() []int {
	res := make([]int, 11)
	t := ns.top // 临时存储栈顶
	for i := 0; t != ns.bottom; i++ {
		res[i] = t.data.data
		t = t.next
	}
	return res
}

// NewThreadStack 创建一个新的链表栈，返回其指针
func NewThreadStack() *threadStack {
	node := &threadNoe{}
	return &threadStack{
		top:    node,
		bottom: node,
	}
}
