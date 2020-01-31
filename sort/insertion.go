package sort

// 插入排序
// 将数组分为前后两部分
// 前面为有序数组，初始为一个元素
// 后面为无序数组，初始为n-1个元素
// 每次将无序数组中的第一个元素插入到有序数组的相应位置
func Insertion(data []int) {
	n := len(data)
	var j int
	// i表示无序数组最小下标
	for i := 1; i < n; i++ {
		// 找到插入的下标
		j = insertionIndex(data, i)
		if j != i { // 如果是最后一位，不需要移动
			// 再移动后面的数据并在j位置插入元素
			move(data, j, i)
		}
	}
}

// insertionIndex 查找要插入元素的下标
// data 待排序序列
// index 需要插入的元素下标
func insertionIndex(data []int, index int) int {
	// 取出要插入的元素
	insertion := data[index]
	// 先比较有序序列的首元素
	if insertion < data[0] {
		return 0
	}
	// 再比较有序序列的尾元素
	if insertion >= data[index-1] {
		return index
	}
	// 再用折半查找
	left, right := 0, index-1
	var (
		mid int // 中间位置索引
		mv  int // 中间索引的元素值
	)
	for left <= right {
		mid = (left + right) >> 1
		mv = data[mid]
		if insertion <= mv {
			right = mid - 1
		} else {
			if insertion < data[mid+1] {
				return mid + 1
			}
			left = mid + 1
		}
	}
	return -1
}

// 从start位置开始进行移动
// 以end位置为结束点
func move(data []int, start, end int) {
	val := data[end]
	// 后移
	for i := end; i > start; i-- {
		data[i] = data[i-1]
	}
	data[start] = val
}
