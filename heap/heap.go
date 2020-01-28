package heap

// 堆接口
type Heap interface {
	// Peek 返回根结点的值
	Peek() (int, error)

	// Top 返回根结点的值并删除
	Top() (int, error)

	// Push 插入元素
	Push(value int) error

	// Delete 删除元素
	Delete(index int) error

	// IsEmpty 判断是否为空
	IsEmpty() bool

	// Traverse 遍历
	Traverse() []int
}
