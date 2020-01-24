package sort

// 计数排序
// 只适合对整数进行排序
// 1. 计算序列中的最值，确定数据范围(min——max)
// 2. 根据数据范围，创建连续的计数桶(max - min + 1)
// 3. 将序列中相同数据的个数存储到计数桶的对应位置(value - min)
// 4. 遍历计数桶，按顺序填充原序列
func Counting(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}

	// 计算出最大值与最小值
	min, max := maxAndMin(data)
	// 确定计数桶的长度
	num := max - min + 1
	// 用切片表示计数的桶
	bins := make([]int, num)
	// 计数遍历
	for _, d := range data {
		// 计算对应桶的位置
		index := d - min
		// 计数器自增
		bins[index] += 1
	}

	// 填充时的索引
	var k int
	// 按顺序填充到原序列
	for i, v := range bins {
		for j := 0; j < v; j++ {
			data[k] = min + i
			k++
		}
	}
}

// 查找序列中的最大值和最小值
func maxAndMin(data []int) (int, int) {
	min, max := data[0], data[0]
	for _, d := range data {
		if d > max {
			max = d
		} else if d < min {
			min = d
		}
	}
	return min, max
}
