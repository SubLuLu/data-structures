package sort

// 希尔排序 (升级版插入排序)
// 1. 按照gap步长进行分组
// 2. 在每组中按照插入排序进行排序
// 3. 改变gap值，重新分组
// 4. 继续在每组中按照插入排序进行排序
// 5. 当gap=1时，完成排序
func Shell(data []int) {
	n := len(data)
	var j int
	// 确定增量步长，按gap进行分组
	for gap := n / 2; gap > 0; gap /= 2 {
		// i从gap开始，确保从每一组的第二个元素开始遍历
		for i := gap; i < n; i++ {
			// j从i-gap开始，确保从每一组的第一个元素开始遍历
			// 找到需要插入元素的位置下标
			for j = i - gap; j >= 0 && data[j] >= data[i]; j -= gap {
			}
			// 再移动需要插入元素位置后面的数据
			// 由于前一个for循环中，
			// j -= gap多执行了一次
			// 所以moveShell函数的第二个参数为j+gap
			moveShell(data, j+gap, i, gap)
		}
	}
}

// 从start位置开始进行移动
// 以end位置为结束
// gap为步长
func moveShell(data []int, start, end, gap int) {
	val := data[end]
	for i := end; i > start; i -= gap {
		data[i] = data[i-gap]
	}
	data[start] = val
}
