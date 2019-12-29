package queue

import (
	"errors"
)

// 使用数组的方式实现循环队列
//
// 添加一个count属性
// 可以迅速判断队列是否已满或者为空
// 不需要清除弹出的元素
// 同时避免front和rear一直自增，导致计算时取模
//
// 如果不添加count属性
// front 指向队列的第一个结点
// rear  指向队列的最后一个结点的下一个结点
// 所以当队列满时，arr中还有一个未使用
// 即arr的length应该为size + 1
// 队列为空条件 front == rear
// 队列已满条件 (rear + 1) % (size + 1) == front
type arrayQueue struct {
	size  int   // 队列的大小
	count int   // 队列中的元素个数
	arr   []int // 存储队列中元素的切片
	front int   // 头游标，0-(size-1)
	rear  int   // 尾游标，0-(size-1)
}

// Enqueue 向队列中添加元素
// num 等待添加的元素
// 返回值为添加元素过程中的错误信息
func (aq *arrayQueue) Enqueue(num int) error {
	if aq.IsFull() { // 队列已满，不能添加元素
		return errors.New("array queue is full")
	}
	aq.arr[aq.rear] = num // 在队列尾部添加元素
	aq.count++ // 队列中的元素个数增加一个
	// 避免使用取模运算
	// 当rear = size后，就会从0重新开始循环
	if t := aq.rear + 1; t == aq.size {
		aq.rear = 0
	} else {
		aq.rear = t
	}
	return nil
}

// Dequeue 从队列中弹出元素
// 返回值为队列中的头部元素或者错误信息
func (aq *arrayQueue) Dequeue() (int, error) {
	if aq.IsEmpty() { // 队列为空，不能弹出元素
		return -1, errors.New("array queue is empty")
	}
	res := aq.arr[aq.front] // 取出队列头部元素
	aq.count-- // 队列中的元素个数减少一个
	// 避免使用取模运算
	// 当front = size后，就会从0重新开始循环
	if h := aq.front + 1; h == aq.size {
		aq.front = 0
	} else {
		aq.front = h
	}
	return res, nil
}

// IsFull 判断队列是否已满
// 返回值表示队列是否已满 true表示已满
func (aq *arrayQueue) IsFull() bool {
	// 元素个数等于size表示队列已满
	return aq.count == aq.size
}

// IsEmpty 判断队列是否为空
// 返回值表示队列是否为空 true表示为空
func (aq *arrayQueue) IsEmpty() bool {
	// 元素个数为0表示队列为空
	return aq.count == 0
}

// Traverse 遍历队列，以切片形式返回
func (aq *arrayQueue) Traverse() []int {
	res := make([]int, aq.count)
	var index = aq.front
	for i:= 0; i != aq.count; i++ {
		res[i] = aq.arr[index]
		if i := index + 1; i == aq.size {
			index = 0
		} else {
			index = i
		}
	}
	return res
}

// NewArrayQueue 根据指定的size创建一个队列，并返回其指针
func NewArrayQueue(size int) *arrayQueue {
	return &arrayQueue{
		size: size,
		arr: make([]int, size),
	}
}
