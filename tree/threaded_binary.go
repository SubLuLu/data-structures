package tree

// 线索二叉树结点
type threadedNode struct {
	data                  int           // 结点数据
	leftTag, rightTag     bool          // false表示该位置是孩子结点
	leftChild, rightChild *threadedNode // 孩子结点
}

// 线索二叉树
type threadedBinary struct {
	head *threadedNode // 头结点
}

// 前序线索化
// 将为空的左孩子结点指向前驱结点(双亲结点)
// 将为空的右孩子结点指向后继结点
func (tb *threadedBinary) PreThreading() {
	root := tb.head
	preFunc := preNode()

	preThreading(root, preFunc)
}

// 对二叉树前序线索化的递归实现
func preThreading(node *threadedNode, preFunc func(read bool, node *threadedNode) *threadedNode) {
	if node != nil { // 当前结点作为根结点
		// 从闭包中获取当前结点的前驱结点
		// pre为当前结点(node)的前驱结点
		pre := preFunc(true, nil)
		// pre不为nil，pre的右孩子为nil
		// pre的后继结点指向当前结点
		if pre != nil && pre.rightTag {
			// pre的右孩子指向当前结点
			pre.rightChild = node
		}
		if node.leftChild == nil { // 左子树为空
			node.leftTag = true // 前驱线索标识
			node.leftChild = pre
		}
		if node.rightChild == nil { // 右子树为空
			// rightChild存储后继结点
			// 由于此轮无法确定其后继结点
			// 所以将该操作放在下一轮中
			// pre.rightChild = node
			node.rightTag = true // 后继线索标识
		}
		// 临时保存当前结点作为下一次的前驱结点
		preFunc(false, node)
		// 在递归遍历前需要确保子树存在，避免进入死循环
		// 当前序遍历到叶子结点时，其左孩子是前驱结点(即双亲结点)
		// 如果不进行判断，将继续从其前驱结点开始遍历(进入死循环)
		if !node.leftTag { // 确保当前结点的左子树存在
			preThreading(node.leftChild, preFunc) // 递归左子树
		}
		if !node.rightTag { // 确保当前结点的右子树存在
			preThreading(node.rightChild, preFunc) // 递归右子树
		}
	}
}

// 递归前序线索化时利用闭包临时保存当前结点和前驱结点
// 避免定义全局变量
func preNode() func(read bool, node *threadedNode) *threadedNode {
	var pre *threadedNode
	// read是否只读标识
	return func(read bool, node *threadedNode) *threadedNode {
		if !read { // 非只读模式
			pre = node
		}
		return pre
	}
}

// 前序线索化后的遍历
// 1. 从根结点开始访问
// 2. 依次访问左子树
//    如果左孩子结点是线索化结点(前驱结点)，则访问右子树
// 3. 每次访问时，把作为子树根的结点输出
func (tb *threadedBinary) PreThreadedTraverse() *list {
	l := newList()
	node := tb.head
	for node != nil {
		l.append(node.data)
		if node.leftChild != nil {
			if !node.leftTag {
				node = node.leftChild
			} else {
				node = node.rightChild
			}
		}
	}
	return l
}

// 中序线索化
// 将为空的左孩子结点指向前驱结点
// 将为空的右孩子结点指向后继结点
func (tb *threadedBinary) InThreading() {
	root := tb.head
	inFunc := inNode()
	inThreading(root, inFunc)
}

// 对二叉树中序线索化的递归实现
func inThreading(node *threadedNode, inFunc func(node *threadedNode) *threadedNode) {
	if node != nil {
		// 中序遍历左子树
		inThreading(node.leftChild, inFunc)
		// 获取前驱结点，并将当前结点作为下一次的前驱结点保存
		pre := inFunc(node)

		if node.leftChild == nil {
			// 左子树遍历完成了
			node.leftTag = true
			// 将左孩子指向其前驱结点
			node.leftChild = pre
		}
		// 前驱结点没有右孩子，前驱结点的后继结点指向当前结点
		if pre != nil && pre.rightChild == nil {
			pre.rightTag = true
			pre.rightChild = node
		}
		// 中序遍历右子树
		inThreading(node.rightChild, inFunc)
	}
}

// 利用闭包和defer的特性，临时存储前驱结点
// 避免定义全局变量
func inNode() func(node *threadedNode) *threadedNode {
	var pre *threadedNode
	return func(node *threadedNode) *threadedNode {
		defer func() {
			// 2.再赋值
			pre = node
		}()
		// 1.先返回
		return pre
	}
}

// 中序线索化后的遍历
// 1. 先找到最左边的结点(即中序遍历时的第一个结点)
// 2. 输出该结点
// 3. 遍历右子树
//    如果右孩子是线索化结点(即后继结点)，则输出
func (tb *threadedBinary) InThreadedTraverse() *list {
	l := newList()
	node := tb.head
	for node != nil {
		// 先遍历正常的左子树，找到最左边的结点
		for node.leftChild != nil && !node.leftTag {
			node = node.leftChild
		}

		l.append(node.data)

		// 根据线索化后的指针，查找后继结点
		for node.rightChild != nil && node.rightTag {
			node = node.rightChild
			l.append(node.data)
		}

		// 如果右孩子指针指向的不是后继结点，则遍历右子树
		node = node.rightChild
	}
	return l
}

// 后序线索化
// 将为空的左孩子结点指向前驱结点
// 将为空的右孩子结点指向后继结点
func (tb *threadedBinary) PostThreading() {
	root := tb.head
	postFunc := postNode()

	postThreading(root, postFunc)
}

// 对二叉树后序线索化的递归实现
func postThreading(node *threadedNode, postFunc func(node *threadedNode) *threadedNode) {
	if node != nil {
		postThreading(node.leftChild, postFunc)
		postThreading(node.rightChild, postFunc)
		// 获取前驱结点，并将当前结点作为下一次的前驱结点保存
		pre := postFunc(node)

		// 左结点为nil，存储前驱结点
		if node.leftChild == nil {
			node.leftTag = true
			node.leftChild = pre
		}

		// 将当前结点作为前驱结点的后继结点
		if pre != nil && pre.rightChild == nil {
			pre.rightTag = true
			pre.rightChild = node
		}
	}
}

// 利用闭包和defer的特性，临时存储前驱结点
// 避免定义全局变量
func postNode() func(node *threadedNode) *threadedNode {
	var pre *threadedNode
	return func(node *threadedNode) *threadedNode {
		defer func() {
			// 2.再赋值
			pre = node
		}()
		// 1.先返回
		return pre
	}
}

// 后序线索化后的遍历
// 1. 先找到最左边的结点
// 2. 将当前结点的双亲结点入栈
// 3. 如果右子树已经遍历完成，从栈中找双亲结点
//    可以将双亲结点在线索化的售后直接保存在结点中
// 4. 如果遍历到根，而且右子树已经访问过，直接将根输出
func (tb *threadedBinary) PostThreadedTraverse() *list {
	l := newList()

	var (
		err       error         // 接受栈操作的error
		pre       *threadedNode // 前驱结点
		traversed bool          // 根结点是否被访问标识
	)

	roots := NewThreadStack()
	root := tb.head // 存储根结点，整棵树的根和右子树的根
	node := tb.head // 遍历时的中间变量，从根结点开始

	// 由于threadedNode中没有parent结点的记录
	// 所以借助栈进行查找双亲结点
	for node != nil {
		// 先遍历正常的左子树，找到最左边的结点
		for node.leftChild != nil && !node.leftTag && node != pre {
			// 压栈，供后序操作查找双亲结点
			err = roots.Push(node)
			// 继续向左查找
			node = node.leftChild
		}

		// 根据线索化后的指针，查找后继结点
		// 此时该结点的左孩子结点已经在之前访问完了
		// 右孩子结点是后继结点，所以输出当前结点
		for node.rightChild != nil && node.rightTag {
			// 输出当前结点
			l.append(node.data)
			// 把当前结点作为下一次的pre结点
			pre = node
			// 继续向右查找
			node = node.rightChild
			// 因为threadedNode中没有存储双亲结点
			// 为了能够通过栈找到双亲结点，此处需要处理
			// 1. 一个结点的后继结点可能是其兄弟结点，也可能是双亲结点
			// 2. 如果是双亲结点，则需要从栈中弹出，以免查找双亲结点出错
			if !roots.IsEmpty() {
				// 先从栈中弹出
				t, e := roots.Pop()
				if e != nil {
					break
				}
				// 判断后继结点是否与栈中的结点相同(即是否是双亲结点)
				if t != node {
					// 不是双亲结点，把弹出的结点再次入栈
					err = roots.Push(t)
				}
			}
		}

		// 判断当前结点的右子树是否访问过
		// 如果右孩子不为nil，而且右孩子等于上一次访问的结点，则说明右子树被访问过
		if node.rightChild != nil && node.rightChild != pre {
			// 右子树没被访问过，遍历右子树
			node = node.rightChild
		} else {
			// 右子树已经被访问，将当前结点输出
			l.append(node.data)
			// 从栈中弹出元素找双亲结点
			pre, err = roots.Pop()
			if err != nil {
				break
			}
			// 将前驱结点和当前结点都指向从栈中弹出的结点
			// 此处的目的是标记该结点的左子树已经遍历过
			// 对应遍历左子树时的node != pre的条件
			node = pre
		}

		// node == pre 说明左子树已经被访问
		// node == root 说明已经访问到了根
		if node == pre && node == root {
			// 右子树存在(根的右子树肯定不是线索化的)
			// !traversed == true表示根第一被访问
			if node.rightChild != nil && !traversed {
				// 修改标识
				traversed = true
				// 开始访问根右子树
				node = node.rightChild
				// 将右子树作为一棵独立的二叉树
				root = node
			} else {
				// 右子树不存在，已经到根了，直接输出
				l.append(node.data)
				break
			}
		}
	}

	// 因为右子树的根结点访问完后就退出了循环
	// 所以整棵树的根结点还没有输出
	// 最后输出根结点
	l.append(tb.head.data)

	return l
}
