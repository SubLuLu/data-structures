package hash

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

// 散列表
type hashTable struct {
	// 读写锁
	lock sync.RWMutex
	// 键值对，key为hash值
	items map[int]Value

}

// HashCode 计算key的hash值
// 选取素数31计算多项式的和
func HashCode(k Key) int {
	var key string
	if str, ok := k.(Stringer); ok {
		key = str.String()
	} else {
		// 将k格式化为字符串
		key = fmt.Sprintf("%s", k)
	}
	// 按字符分割
	keyRune := []rune(key)
	hc := 0
	for _, ch := range keyRune {
		// 通常根据Horner法则，计算一个多项式
		// 形式为：int(keyRune[0])*k^(n-1) + int(keyRune[1])*k^(n-2) + ... + int(keyRune[n-1])
		// 其中k是一个常数，一般选择一个大小适中的素数，根据除留余数法可知，素数会减少冲突
		// 其中n是字符数组的长度
		hc = 31*hc + int(ch)
	}
	return hc
}

// Put 向散列表中添加键值对
func (ht *hashTable) Put(k Key, v Value) error {
	if k == nil { // key不能为nil
		return errors.New("key can not be nil")
	}
	if v == nil { // value不能为nil
		return errors.New("value can not be nil")
	}

	ht.lock.Lock()
	defer ht.lock.Unlock()
	code := HashCode(k)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[code] = v
	return nil
}

// Get 根据关键字从散列表获取值
func (ht *hashTable) Get(k Key) (Value, error) {
	if k == nil { // key不能为nil
		return nil, errors.New("key can not be nil")
	}

	ht.lock.RLock()
	defer ht.lock.RUnlock()
	if len(ht.items) == 0 {
		return nil, errors.New("table is empty")
	}
	code := HashCode(k)
	return ht.items[code], nil
}

// Delete 根据关键字从散列表中删除元素
func (ht *hashTable) Delete(k Key) error {
	if k == nil { // key不能为nil
		return errors.New("key can not be nil")
	}

	ht.lock.Lock()
	defer ht.lock.Unlock()
	if len(ht.items) == 0 {
		return errors.New("table is empty")
	}
	code := HashCode(k)
	delete(ht.items, code)
	return nil
}

// Len 返回散列表中的键值对总数
func (ht *hashTable) Len() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)
}

func NewHashTable() *hashTable {
	return &hashTable{
		items: make(map[int]Value),
	}
}

// 高级版散列表
type seniorTable struct {
	lock   sync.RWMutex // 读写锁
	values []*pair      // 每个链表
	total  uint64       // 记录元素个数
	prime  int          // 除留余数法的素数p
}

// 键值对链表结构
type pair struct {
	key     Key            // 关键字
	element unsafe.Pointer // Value值的内存地址
	next    unsafe.Pointer // 下一个pair的内存地址
}

// 获取key
func (p *pair) Key() Key {
	return p.key
}

// 获取value
func (p *pair) Element() Value {
	pointer := atomic.LoadPointer(&p.element)
	if pointer == nil {
		return nil
	}
	return *(*Value)(pointer)
}

// 设置value值，为覆盖提供服务
func (p *pair) SetElement(v Value) {
	// 确保v != nil
	atomic.StorePointer(&p.element, unsafe.Pointer(&v))
}

// 获取链表的下一个键值对
func (p *pair) Next() *pair {
	pointer := atomic.LoadPointer(&p.next)
	if pointer == nil {
		return nil
	}
	return (*pair)(pointer)
}

// 设置链表的下一个键值对
func (p *pair) SetNext(next *pair) {
	if next == nil {
		atomic.StorePointer(&p.next, nil)
		return
	}

	atomic.StorePointer(&p.next, unsafe.Pointer(next))
}

// newPair 创建新的键值对
func newPair(key Key, value Value) *pair {
	p := &pair{
		key: key,
	}
	p.element = unsafe.Pointer(&value)
	return p
}

// Put 向散列表中添加键值对
func (st *seniorTable) Put(k Key, v Value) error {
	if k == nil { // key不能为nil
		return errors.New("key can not be nil")
	}
	if v == nil { // value不能为nil
		return errors.New("value can not be nil")
	}
	// 根据除留余数法计算散列表的位置
	hc, err := Division(k, st.prime)
	if err != nil {
		return err
	}

	st.lock.Lock()
	defer st.lock.Unlock()

	// 取出已经存在的键值对
	pe := st.values[hc]
	if pe == nil { // 不存在，则直接添加
		st.values[hc] = newPair(k, v)
		// 使用原子操作，确保并发安全，散列表中的大小加1
		atomic.AddUint64(&st.total, 1)
		return nil
	}
	var target *pair
	// 链表不为空，查找key是否已经存在
	for p := pe; p != nil; p = p.Next() {
		if k == p.Key() {
			target = p
			break
		}
	}
	if target != nil { // 已经存在，直接覆盖value值
		target.SetElement(v)
		return nil
	}

	p := newPair(k, v) // 创建新的键值对
	p.SetNext(pe)      // 将当前键值对设置为表头
	st.values[hc] = p  // 将表头添加到散列表中
	atomic.AddUint64(&st.total, 1)
	return nil
}

// Get 根据关键字从散列表获取值
func (st *seniorTable) Get(k Key) (Value, error) {
	if k == nil { // key不能为nil
		return nil, errors.New("key can not be nil")
	}
	// 根据除留余数法计算散列表的位置
	hc, err := Division(k, st.prime)
	if err != nil {
		return nil, err
	}

	st.lock.RLock()
	defer st.lock.RUnlock()

	// 取出已经存在的键值对
	pe := st.values[hc]
	// 遍历链表找到value值
	for p := pe; p != nil; p = p.Next() {
		if k == p.Key() {
			return p.Element(), nil
		}
	}
	return nil, nil
}

// Delete 根据关键字从散列表中删除元素
func (st *seniorTable) Delete(k Key) error {
	if k == nil { // key不能为nil
		return errors.New("key can not be nil")
	}
	// 根据除留余数法计算散列表的位置
	hc, err := Division(k, st.prime)
	if err != nil {
		return err
	}

	st.lock.Lock()
	defer st.lock.Unlock()

	// 取出已经存在的键值对
	pe := st.values[hc]
	var prev, target *pair // 查找需要删除的键值对
	for p := pe; p != nil; p = p.Next() {
		if k == p.Key() {
			target = p
			break
		}
		prev = p
	}
	if target == nil {
		return errors.New("key does not exist")
	}
	if prev == nil { // 要删除的是表头
		pe.SetNext(nil)
		st.values[hc] = target
	} else {
		// target是最后一个
		if target.Next() == nil {
			prev.SetNext(nil)
		} else {
			prev.SetNext(target.Next())
			target.SetNext(nil)
		}
	}
	// 使用原子操作，确保并发安全，散列表中的大小减1
	atomic.AddUint64(&st.total, ^uint64(0))
	return nil
}

// Len 返回散列表中的键值对总数
func (st *seniorTable) Len() int {
	return int(atomic.LoadUint64(&st.total))
}

// NewSeniorTable 创建一个指定表长的散列表
func NewSeniorTable(size int) *seniorTable {
	return &seniorTable{
		values: make([]*pair, size),
		total:  0,
		prime:  maxPrime(size),
	}
}
