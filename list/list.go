package list

// 链表方法接口
type List interface {
	// Add 从头部插入结点
	Add(val int) error
	
	// Append 从尾部插入结点
	Append(val int) error

	// 从指定位置插入结点
	Insert(val, index int) error

	// Delete 删除指定位置的结点
	Delete(index int) error

	// Set 给指定位置结点赋值
	Set(val, index int) error

	// Find 查找指定位置结点元素值
	Find(index int) (int, error)

	// Length 返回链表中有效结点个数
	Length() int

	// IsEmpty 判断链表是否为空
	IsEmpty() bool

	// Traverse 遍历链表，以切片形式返回
	Traverse() []int
}
