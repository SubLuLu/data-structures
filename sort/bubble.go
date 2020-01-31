package sort

// 冒泡排序
// 利用头尾指针，
// 头指针指向需要排序的元素，
// 尾指针指向已经排好序的元素
// 相邻两个元素进行比较
// 反序则交换
// 如果在某一轮中没有发生交换，
// 表示已经完成了排序
func Bubble(data []int) {
	n := len(data)
	// 如果经过某次排序后已经达到有序，
	// 则直接退出排序
	var flag bool // 是否发生过交换标识
	// 每次排序最后下标
	for i := n - 1; i > 1; i-- {
		// 每次排序起始下标
		for j := 0; j < i; j++ {
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true
			}
		}
		if !flag {
			break
		} else {
			flag = false
		}
	}
}
