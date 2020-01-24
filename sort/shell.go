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
		for i := gap; i < n; i++ { // i是每组中的第二个元素
			// j起始值为每组中的第一个元素
			for j = i - gap; j >= 0 && data[j] > data[i]; j -= gap {
			}
			// 再移动后面的数据
			moveShell(data, j + gap, i, gap)
		}
	}
}

// 从start位置开始进行移动
// 以end位置为结束
func moveShell(data []int, start, end, gap int) {
	val := data[end]
	for i := end; i > start; i -= gap {
		data[i] = data[i - gap]
	}
	data[start] = val
}
