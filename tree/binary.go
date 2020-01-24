package tree

// 顺序存储
// 一般存储完全二叉树
// 在二叉树的性质中有关于完全二叉树结点和其孩子结点的位置关系
// 第i个结点的左孩子结点为2i，右孩子结点为2i+1
// 由此推出第i个结点的双亲结点为⌊i/2⌋(向下取整)
// 此处涉及的小标都是从1开始，而且要确保2i和2i+1小于树的结点数
// 因此为了记录结点的双亲孩子关系
// 一般二叉树顺序存储时，需要按照完全二叉树进行存储
// 即把一般二叉树转换为完全二叉树时，不存在的结点值设置为空
type arrayBinary []*int

// 前序遍历
func (ab arrayBinary) PreOrder() []int {
	n := len(ab)
	if n == 0 { // 空树，直接返回
		return []int{}
	}
	if n == 1 { // 只有一个结点
		val := ab[0]
		if val != nil { // 结点值不为空
			return []int{*val}
		}
		return []int{}
	}
	// 根据有效结点个数申请相应大小的切片
	result := make([]int, ab.count())
	var pos = index() // 索引值用闭包
	// 前序遍历从根开始
	pre(ab, result, 0, pos)
	return result
}

// 利用递归实现前序遍历
// root是结点在[]int中的索引，从0开始
// 先根
// 再左子树
// 最后右子树
func pre(data arrayBinary, result []int, root int, pos func() int) {
	n := len(data)
	// 递归退出条件为结点索引不小于data长度
	if root >= n {
		return
	}
	// 每颗子树的根结点
	if data[root] != nil {
		// 如果当前结点为有效值，则存储到result中
		result[pos()] = *data[root]
	}
	// 左子树
	// 根据完全二叉树的性质
	// i结点的左孩子是2*i(i是从1开始的)
	// 而golang的切片中索引是从0开始的
	// 所以root结点的左孩子结点是2*root+1
	pre(data, result, 2*root+1, pos)
	// 右子树
	// 根据左孩子的确定，可知
	// root结点的左孩子结点是2*root+2
	pre(data, result, 2*root+2, pos)
}

// 中序遍历
func (ab arrayBinary) InOrder() []int {
	n := len(ab)
	if n == 0 { // 空树，直接返回
		return []int{}
	}
	if n == 1 { // 只有一个结点
		val := ab[0]
		if val != nil { // 结点值不为空
			return []int{*val}
		}
		return []int{}
	}
	// 根据有效结点个数申请相应大小的切片
	result := make([]int, ab.count())
	var pos = index() // 索引值用闭包
	// 中序遍历从左子树开始
	in(ab, result, 0, pos)
	return result
}

// 利用递归实现中序遍历
// 先左子树
// 再根
// 最后右子树
func in(data arrayBinary, result []int, root int, pos func() int) {
	n := len(data)
	// 递归退出条件为结点索引不小于data长度
	if root >= n {
		return
	}
	// 左子树
	// 根据完全二叉树的性质
	// i结点的左孩子是2*i(i是从1开始的)
	// 而golang的切片中索引是从0开始的
	// 所以root结点的左孩子结点是2*root+1
	in(data, result, 2*root+1, pos)
	// 每颗子树的根结点
	if data[root] != nil {
		// 如果当前结点为有效值，则存储到result中
		result[pos()] = *data[root]
	}
	// 右子树
	// 根据左孩子的确定，可知
	// root结点的左孩子结点是2*root+2
	in(data, result, 2*root+2, pos)
}

// 后续遍历
func (ab arrayBinary) PostOrder() []int {
	n := len(ab)
	if n == 0 { // 空树，直接返回
		return []int{}
	}
	if n == 1 { // 只有一个结点
		val := ab[0]
		if val != nil { // 结点值不为空
			return []int{*val}
		}
		return []int{}
	}
	// 根据有效结点个数申请相应大小的切片
	result := make([]int, ab.count())
	var pos = index() // 索引值用闭包
	// 后续遍历从左子树开始
	post(ab, result, 0, pos)
	return result
}

// 利用递归实现后续遍历
// 先左子树
// 再右子树
// 最后根
func post(data arrayBinary, result []int, root int, pos func() int) {
	n := len(data)
	// 递归退出条件为结点索引不小于data长度
	if root >= n {
		return
	}
	// 左子树
	// 根据完全二叉树的性质
	// i结点的左孩子是2*i(i是从1开始的)
	// 而golang的切片中索引是从0开始的
	// 所以root结点的左孩子结点是2*root+1
	post(data, result, 2*root+1, pos)
	// 右子树
	// 根据左孩子的确定，可知
	// root结点的左孩子结点是2*root+2
	post(data, result, 2*root+2, pos)
	// 每颗子树的根结点
	if data[root] != nil {
		// 如果当前结点为有效值，则存储到result中
		result[pos()] = *data[root]
	}
}

// 计算二叉树中有效结点个数
func (ab arrayBinary) count() int {
	var count int
	// 遍历一遍，确定二叉树中有效结点个数
	for _, v := range ab {
		if v != nil {
			count++
		}
	}
	return count
}

// index返回一个计算索引位置的闭包函数
// 因为在递归调用中，会为每次的递归开辟一片内存空间存储当前的局部变量
// 所以当遍历完左子树后，第一次遍历右子树时，所对应的index值是第一次遍历左子树时的值
// 因此用闭包解决该问题，或者定义一个全局变量也可解决此问题
func index() func() int {
	var i = -1
	return func() int {
		i++
		return i
	}
}

// 链式存储
// 因为二叉树最多两个孩子结点
// 所以设计带有一个数据域和两个指针域的二叉链表进行存储
type listBinary struct {
	data  int         // 数据域
	left  *listBinary // 左孩子
	right *listBinary // 右孩子
}

// 前序遍历
func (lb *listBinary) PreOrder() *list {
	l := newList()
	preOrder(lb, l)
	return l
}

// 利用递归实现前序遍历
func preOrder(node *listBinary, l *list) {
	if node == nil {
		return
	}
	l.append(node.data)
	preOrder(node.left, l)
	preOrder(node.right, l)
}

// 中序遍历
func (lb *listBinary) InOrder() *list {
	l := newList()
	inOrder(lb, l)
	return l
}

// 利用递归实现中序遍历
func inOrder(node *listBinary, l *list) {
	if node == nil {
		return
	}
	inOrder(node.left, l)
	l.append(node.data)
	inOrder(node.right, l)
}

// 后序遍历
func (lb *listBinary) PostOrder() *list {
	l := newList()
	postOrder(lb, l)
	return l
}

// 利用递归实现后序遍历
func postOrder(node *listBinary, l *list) {
	if node == nil {
		return
	}
	postOrder(node.left, l)
	postOrder(node.right, l)
	l.append(node.data)
}
