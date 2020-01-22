package hash

type Key interface{}

type Value interface{}

type Stringer interface {
	String() string
}

// 散列表接口
type Table interface {
	// Pet 向散列表中添加元素
	Pet(key Key, value Value) error
	// Delete 从散列表中删除元素
	Delete(key Key) error
	// Get 从散列表中获取指定元素
	Get(key Key) (Value, error)
	// Len 获取散列表的表长
	Len() int
}
