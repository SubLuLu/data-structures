package queue

// 队列接口
type Queue interface {
	// Enqueue 向队列中添加元素
	Enqueue(num int) error
	// Dequeue 从队列中弹出元素
	Dequeue() (int, error)
	// IsEmpty 判断队列是否为空
	IsEmpty() bool
	// Traverse 遍历队列，以切片形式返回
	Traverse() []int
}
