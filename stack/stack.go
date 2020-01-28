package stack

// 栈接口
type Stack interface {
	// Push 压栈，向栈中添加元素
	Push(val int) error

	// Pop 出栈，从栈顶弹出元素
	Pop() (int, error)

	// IsEmpty 判断栈是否为空
	IsEmpty() bool

	// Traverse 遍历栈，以切片形式返回
	Traverse() []int
}
