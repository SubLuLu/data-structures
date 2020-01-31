package sort

// 选择排序
// 1. 默认起始元素是最小的
// 2. 每轮循环找到一个最小的
// 3. 和起始元素进行交换
func Selection(data []int) {
	n := len(data)
	// 最小值的下标
	var index int
	// 起始元素从0开始
	for i := 0; i < n; i++ {
		index = i // 每次找寻最小值时初始化为起始下标
		for j := i; j < n; j++ {
			if data[j] < data[index] {
				index = j
			}
		}
		// 第一个就是最小值，则不需要交换
		if index != i {
			data[i], data[index] = data[index], data[i]
		}
	}
}
