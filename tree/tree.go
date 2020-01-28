package tree

// 树的几种基础表示法

// 1. 双亲表示法

// 每个结点用一个指针或者下标指向双亲结点
type parentTreeNode struct {
	data   int // 结点数据
	parent int // 双亲结点下标，根节点的值为-1
}

// 双亲表示法可以很好的找到结点的双亲结点
// 很容易找到根节点，但是找孩子结点需要遍历整棵树
type parentTree struct {
	tree []*parentTreeNode // 结点集合
	root int               // 根结点位置
	size int               // 结点总数
}

func NewParentTree(size int) *parentTree {
	return &parentTree{
		tree: make([]*parentTreeNode, size),
		root: 0,
		size: size,
	}
}

// 2. 孩子表示法

// 孩子结点链表
type childNode struct {
	index int        // 结点在数组中的下标
	next  *childNode // 下一个孩子结点(兄弟)
}

// 每个结点用一个指针指向孩子结点所形成的链表
type childTreeNode struct {
	data  int        // 结点数据
	child *childNode // 孩子结点位置信息
}

// 孩子表示法很容易找到结点的孩子和兄弟
// 但是要查找双亲时，需要遍历整棵树
type childTree struct {
	tree []*childTreeNode // 结点集合
	root int              // 根结点位置
	size int              // 结点总数
}

func NewChildTree(size int) *childTree {
	return &childTree{
		tree: make([]*childTreeNode, size),
		root: 0,
		size: size,
	}
}

// 3. 双亲孩子表示法
//
// 为了能同时方便查找孩子和双亲，
// 将上述两种表示方法进行整合

// 每个结点即标注双亲又记录孩子
type parentChildNode struct {
	data   int        // 结点数据
	parent int        // 双亲结点下标，根节点的值为-1
	child  *childNode // 孩子结点位置信息
}

// 这样既能像双亲表示法一样快速找到双亲
// 又能像孩子表示法一样迅速查找孩子
type parentChildTree struct {
	tree []*parentChildNode // 结点集合
	root int                // 根结点位置
	size int                // 结点总数
}

func NewParentChildTree(size int) *parentChildTree {
	return &parentChildTree{
		tree: make([]*parentChildNode, size),
		root: 0,
		size: size,
	}
}

// 孩子兄弟表示法
// 该表示法的好处在于把一颗普通树转换为二叉树

// 每个结点包含两个指针
// 一个指向第一个孩子结点
// 一个指向右边第一个兄弟结点
type childSiblingNode struct {
	data    int               // 结点数据
	child   *childSiblingNode // 第一个孩子结点
	sibling *childSiblingNode // 第一个右兄弟结点
}

func NewChildSiblingNode(data int) *childSiblingNode {
	return &childSiblingNode{
		data: data,
	}
}
