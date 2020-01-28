package tree

// 二叉树非递归遍历
// 前序遍历：根——左——右
// 中序遍历：左——根——右
// 后序遍历：左——右——根
// 从三种遍历方式上看，主要是描述根的先后顺序
// 所以只需要解决所有根(非叶子结点)的顺序
// 由于根据根查找左右孩子结点时，需要逆序输出根的顺序
// 因此，要将根的顺序临时存储
// 而关于逆序输出最有效的数据就结构就是栈

// 顺序存储的二叉树前序遍历——非递归实现
func (ab arrayBinary) PreOrderTraverse() []int {
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
	var (
		index int
		err   error
	)
	// 存储非叶子结点
	roots := NewLinkListStack()
	for i := 0; i < n; {
		if ab[i] != nil { // 结点值是否存在
			result[index] = *ab[i]     // 将遍历过的值存入result中
			index++                    // result中的下标自增
			j := 2*i + 1               // 再找左结点
			if j < n && ab[j] != nil { // 左结点存在，入栈，作为子树的根
				err = roots.Push(i)
				i = j
				continue
			}
		}
		// 结点不存在，或者结点值为空，出栈，作为双亲结点
		// 如果栈为空，则表示遍历完成
		i, err = roots.Pop()
		if err != nil { // 栈为空出栈错误
			break
		}
		j := 2*i + 2 // 再找右结点
		// 这个循环最多执行两次
		// 因为在完全二叉树中，不可能祖孙三代都没有右孩子结点
		for j >= n {
			// 结点不存在，或者结点值为空，出栈，作为双亲结点
			// 如果栈为空，则表示遍历完成
			i, err = roots.Pop()
			if err != nil { // 栈为空出栈错误
				break
			}
			j = 2*i + 2 // 找右结点
		}
		i = j // 找到存在的右结点
	}
	return result
}

// 顺序存储二叉树的中序遍历——非递归实现
func (ab arrayBinary) InOrderTraverse() []int {
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

	var (
		index int
		err   error
	)
	// 存储非叶子结点
	roots := NewLinkListStack()
	for i := 0; i < n; {
		if ab[i] != nil { // 结点值是否存在
			err = roots.Push(i)        // 入栈，作为子树的根
			j := 2*i + 1               // 左结点是否存在
			if j < n && ab[j] != nil { // 子树存在，继续找子树
				i = j
				continue
			}
		}

		// 已经没有左子树了，出栈，作为双亲结点
		i, err = roots.Pop()
		if err != nil {
			break
		}
		// 把当前结点存入result中
		result[index] = *ab[i]
		index++ // result的下标自增

		j := 2*i + 2 // 找右结点
		// 这个循环最多执行两次
		// 因为在完全二叉树中，不可能祖孙三代都没有右孩子结点
		for j >= n {
			i, err = roots.Pop()
			if err != nil {
				break
			}
			// 把当前结点存入result中
			result[index] = *ab[i]
			index++ // result的下标自增

			j = 2*i + 2 // 找右结点
		}
		i = j // 右结点作为新的根进入下一轮循环
	}
	return result
}

// 顺序存储的二叉树后序遍历——非递归实现
func (ab arrayBinary) PostOrderTraverse() []int {
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

	var (
		index int
		err   error
		visit = -1   // 最近一次访问的右结点
		flag  = true // 表示是否第一次访问(不是通过出栈访问)
	)
	// 存储非叶子结点
	roots := NewLinkListStack()
	for i := 0; i < n; {
		// 一次性将某个结点的所有左子节点遍历完
		if ab[i] != nil && flag {
			err = roots.Push(i)        // 入栈，作为左子树的根
			j := 2*i + 1               // 左孩子结点是否存在
			if j < n && ab[j] != nil { // 左子树存在，继续找左子树的根
				i = j
				continue
			}
		}

		// 从左子树的最底层开始遍历右子树
		if flag { // 第一次遍历，此时的i是一个叶子结点
			// 先找到双亲结点，遍历右子树
			i, err = roots.Pop()
			if err != nil {
				break
			}
		} else { // 不是第一次遍历，此时的i不是叶子结点(可能子节点为空但是位置存在)
			flag = true // 重置是否第一次访问的标记
		}

		// 找到右结点
		j := 2*i + 2
		// 右结点已经访问过或者右结点不存在或者右结点为空
		if j == visit || j >= n || ab[j] == nil {
			// 当前结点存入result中
			result[index] = *ab[i]
			index++   // result下标自增
			visit = i // 标记该右结点被访问
			// 继续找双亲结点
			i, err = roots.Pop()
			if err != nil {
				break
			}
			// 此时的结点是从栈中获取的，所以已经访问过
			flag = false
		} else { // 右结点存在，没有访问过，以该节点为根进行遍历
			// 因为i是从栈中pop出来的
			// 现在把当前结点放回栈中
			err = roots.Push(i)
			i = j     // 以该节点为根，循环遍历
			visit = i // 标记当前右结点最近被访问
		}
	}

	return result
}

// 顺序存储二叉树层级遍历
// 顺序存储本来就是按照层级存储的
// 只需要去除空结点即可
func (ab arrayBinary) LevelOrderTraverse() []int {
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
	var index int
	for _, v := range ab {
		if v != nil {
			result[index] = *v
			index++
		}
	}
	return result
}

// 链式存储的二叉树的前序遍历——非递归实现
// lb本质上是二叉树的根结点
func (lb *listBinary) PreOrderTraverse() *list {
	// 存放遍历结果的链表
	l := newList()
	// 存放非叶子结点的栈
	roots := NewNodeStack()
	var err error
	t := lb        // 申明一个临时变量进行遍历
	for t != nil { // 以t结点为根
		l.append(t.data)    // 把根结点放入list中
		err = roots.Push(t) // 把根结点放入roots中
		if t.left != nil {  // 找左孩子结点
			t = t.left // 把左孩子结点作为左子树的根
			continue
		}

		// 当前t结点的左子树已经为空，从栈中取出该结点
		t, err = roots.Pop()
		if err != nil {
			break
		}

		// 找当前结点的右孩子结点
		t = t.right

		// 因为栈中是存储的所有左孩子结点
		// 所以只需要找到一个兄弟结点不为空的结点
		for t == nil {
			// 出栈，获得双亲结点
			t, err = roots.Pop()
			if err != nil {
				break
			}
			// 继续访问右孩子
			t = t.right
		}
	}
	return l
}

// 链式存储的二叉树中序遍历——非递归实现
// lb本质上是二叉树的根结点
func (lb *listBinary) InOrderTraverse() *list {
	// 存放遍历结果的链表
	l := newList()
	// 存放非叶子结点的栈
	roots := NewNodeStack()
	var err error
	t := lb // 申明一个临时变量进行遍历
	for t != nil {
		err = roots.Push(t)
		if t.left != nil {
			t = t.left
			continue
		}

		l.append(t.data)

		// 当前t结点的左子树已经为空，从栈中取出该结点
		t, err = roots.Pop()
		if err != nil {
			break
		}

		t = t.right

		// 因为栈中是存储的所有左孩子结点
		// 所以只需要找到一个兄弟结点不为空的结点
		for t == nil {
			// 出栈，获得双亲结点
			t, err = roots.Pop()
			if err != nil {
				break
			}
			// 将当前结点放入list中
			l.append(t.data)
			// 继续访问右孩子
			t = t.right
		}
	}
	return l
}

// 链式存储的二叉树后序遍历——非递归实现
// lb本质上是二叉树的根结点
func (lb *listBinary) PostOrderTraverse() *list {
	// 存放遍历结果的链表
	l := newList()
	// 存放非叶子结点的栈
	roots := NewNodeStack()
	var (
		visit *listBinary
		err   error
		flag  = true
	)
	t := lb // 申明一个临时变量进行遍历
	for t != nil {
		if flag { // 第一次访问就直接放入roots中
			err = roots.Push(t)
			if t.left != nil {
				t = t.left
				continue
			}
		}

		if flag {
			t, err = roots.Pop()
			if err != nil {
				break
			}
		} else {
			flag = true
		}

		temp := t.right
		// 右子树已经遍历过或者右子树为空
		if temp == visit || temp == nil {
			// 当前根结点的左右子树都已经遍历完
			// 将该结点值存入list中
			l.append(t.data)
			// 标记当前结点已经被访问过
			visit = t
			// 从栈中取其双亲结点作为根，遍历其右子树
			t, err = roots.Pop()
			if err != nil {
				break
			}
			flag = false
		} else {
			// 右子树存在，把当前结点放回到roots中
			err = roots.Push(t)
			// 标识该右孩子结点已经访问过
			visit = temp
			// 将右孩子结点temp作为根结点进行遍历
			t = temp
		}
	}
	return l
}

// 链式存储的二叉树层级遍历
// 链式存储结构遍历某一层时需要知道该层所有结点的双亲结点及其顺序
// 所以需要用队列进行存储双亲结点
func (lb *listBinary) LevelOrderTraverse() *list {
	l := newList()      // 存储遍历结果链表
	roots := NewQueue() // 存储每层的双亲结点的队列
	var err error
	t := lb
	l.append(t.data)       // 把根结点放入list中
	err = roots.Enqueue(t) // 把根结点放入roots中
	for !roots.IsEmpty() { // 队列为空，退出循环
		// 从队列中取出结点作为根结点
		t, err = roots.Dequeue()
		if err != nil {
			break
		}
		// 访问该结点的左孩子
		if t.left != nil {
			l.append(t.left.data)
			err = roots.Enqueue(t.left)
		}
		// 访问该结点的右孩子
		if t.right != nil {
			l.append(t.right.data)
			err = roots.Enqueue(t.right)
		}
	}
	return l
}
