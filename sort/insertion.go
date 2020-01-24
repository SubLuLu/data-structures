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
		// j起始值为有序数组最大下标，找到data[i]需要插入的位置
		for j = i - 1; j >= 0 && data[j] > data[i]; j-- {
		}
		// 再移动后面的数据
		move(data, j + 1, i)
	}
}

// 从start位置开始进行移动
// 以end位置为结束点
func move(data []int, start, end int) {
	val := data[end]
	for i := end; i > start; i-- {
		data[i] = data[i - 1]
	}
	data[start] = val
}
