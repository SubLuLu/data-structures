package stack

import (
	"errors"
	"sync"
)

// 使用数组存储方式实现静态栈
type arrayStack struct {
	mu   sync.RWMutex // 读写锁
	size int          // 栈的大小
	arr  []int        // 存储栈中元素的切片
	top  int          // 栈顶下标
}

// Push 压栈，向栈中添加元素
func (as *arrayStack) Push(val int) error {
	if as.IsFull() { // 栈已满，不能添加元素
		return errors.New("array stack is full")
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	as.top++             // 栈顶上移
	as.arr[as.top] = val // 元素存入栈顶
	return nil
}

// Pop 出栈，从栈顶弹出元素
func (as *arrayStack) Pop() (int, error) {
	if as.IsEmpty() { // 栈为空，不能弹出元素
		return 0, errors.New("array stack is empty")
	}

	as.mu.Lock()
	defer as.mu.Unlock()

	val := as.arr[as.top] // 存储栈顶元素
	as.top--              // 栈顶下移
	return val, nil
}

// IsEmpty 判断栈是否为空
// 返回值表示栈是否为空 true表示为空
func (as *arrayStack) IsEmpty() bool {
	as.mu.RLock()
	defer as.mu.RUnlock()

	return as.top == -1
}

// IsFull 判断栈是否已满
// 返回值表示栈是否已满 true表示已满
func (as *arrayStack) IsFull() bool {
	as.mu.RLock()
	defer as.mu.RUnlock()

	return as.top == as.size-1
}

// Traverse 遍历栈，以切片形式返回
func (as *arrayStack) Traverse() []int {
	as.mu.RLock()
	defer as.mu.RUnlock()

	res := make([]int, as.top+1)
	for i := as.top; i >= 0; i-- {
		res[as.top-i] = as.arr[i]
	}
	return res
}

// NewArrayStack 根据指定的size创建一个栈，并返回其指针
func NewArrayStack(size int) *arrayStack {
	return &arrayStack{
		size: size,
		arr:  make([]int, size),
		top:  -1,
	}
}
