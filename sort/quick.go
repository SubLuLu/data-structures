package sort

// 快速排序
// 取任意下标作为基准数，
// 就在以下三种实现上将该下标与其中一个交换
// 最后就化解成为这三种中的一种
func Quick(data []int) {
	n := len(data)

	// 默认取中间的作为基准数
	QuickBaseMiddle(data, 0, n-1)
}

// 以中间数作为基准数的快速排序
func QuickBaseMiddle(data []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	// 确定基准数下标
	mid := (l + r) >> 1
	base := data[mid]
	for l < r {
		// 从左边找一个比基准数大的数
		// 当data[l] <= base时，继续向右找
		for l < mid && data[l] <= base {
			l++
		}

		// 从右边找一个比基准数小的数
		// 当base >= data[r]时，继续向左找
		for mid < r && base <= data[r] {
			r--
		}

		// 基准数两边都恰好满足要求
		if l == r {
			l, r = mid - 1, mid + 1
			break
		}

		if l == mid { // 左边没有比base大的
			// 交换右边找到的较小值和基准元素
			data[mid], data[r] = data[r], data[mid]
			base = data[mid] // 重新确定基准值
			l = left // 左边从头查找
		} else if r == mid { // 右边没有比base小的
			// 交换左边找到的较大值和基准元素
			data[mid], data[l] = data[l], data[mid]
			base = data[mid] // 重新确定基准值
			r = right // 右边从头查找
		} else { // 没相遇则直接交换 l < r
			data[l], data[r] = data[r], data[l]
		}
	}

	// 左边递归调用
	if left < l { // 增加判断条件，防止越界
		QuickBaseMiddle(data, left, l)
	}

	// 右边递归调用
	if r < right { // 增加判断条件，防止越界
		QuickBaseMiddle(data, r, right)
	}
}

// 以左边数作为基准数的快速排序
// 当以左边为基准数时
// 先从右边开始找比基准数小的数
// 再从左边开始找比基准数大的数
func QuickBaseLeft(data []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	base := data[left]
	for l < r {
		// 先从右往左找一个比基准数小的数
		// 当data[r] >= base时，继续向左找
		for l < r && base <= data[r] {
			r--
		}

		// 再从左往右找一个比基准数大的数
		// 当base >= data[l]时，继续向右找
		for l < r && base >= data[l] {
			l++
		}

		if l < r {
			// 两者进行交换
			data[l], data[r] = data[r], data[l]
		}
	}

	if r != left { // 基准数不是最小
		// 此时l == r
		// 交换基准数下标和相遇的位置
		data[left], data[r] = data[r], data[left]
	}

	// 左边递归调用
	if left < l { // 增加判断条件，防止越界
		QuickBaseLeft(data, left, l-1)
	}
	// 右边递归调用
	if r < right { // 增加判断条件，防止越界
		QuickBaseLeft(data, r+1, right)
	}
}

// 以右边数作为基准数的快速排序
// 当以右边为基准数时
// 先从左边开始找比基准数大的数
// 再从右边开始找比基准数小的数
func QuickBaseRight(data []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left, right
	base := data[right]
	for l < r {
		// 先从左往右找一个比基准数大的数
		// 当base >= data[l]时，继续向右找
		for l < r && base >= data[l] {
			l++
		}

		// 再从右往左找一个比基准数小的数
		// 当data[r] >= base时，继续向左找
		for l < r && data[r] >= base {
			r--
		}

		if l < r {
			// 两者进行交换
			data[l], data[r] = data[r], data[l]
		}
	}
	if l != right { // 基准数不是最大
		// 此时l == r
		// 交换基准数下标和相遇的位置
		data[right], data[l] = data[l], data[right]
	}

	// 左边递归调用
	if left < l { // 增加判断条件，防止越界
		QuickBaseRight(data, left, l-1)
	}
	// 右边递归调用
	if r < right { // 增加判断条件，防止越界
		QuickBaseRight(data, r+1, right)
	}
}
