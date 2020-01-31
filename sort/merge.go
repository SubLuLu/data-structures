package sort

// 归并排序
// 采用分治策略，递归调用
// 1. 分解
// 递归分解序列
// 2. 合并
// 逐个合并为有序序列到临时序列中
// 3. 还原
// 将临时序列中的数据还原至母序列中
func Merge(data []int) {
	left, right := 0, len(data) - 1
	// 准备一个需要排序的子序列大小的切片进行存储排好序的数据
	result := make([]int, right - left + 1)
	// 默认对全部数据进行排序
	mergeSort(data, result, left, right)
}

// 先分解，再在result中合并为有序序列，最后还原到母序列data中
//
// data   是需要排序的母序列
// result 是临时存储需要排序序列的切片
// left   是需要排序子序列的起始下标
// right  是需要排序子序列的末尾下标
func mergeSort(data, result []int, left, right int){
	// 递归退出条件
	if left >= right {
		return
	}
	// 将需要排序的子序列分成两份，计算在data中的中间下标位置
	mid := (right + left) >> 1
	// 左边子序列递归分解
	mergeSort(data, result, left, mid)
	// 右边子序列递归分解
	mergeSort(data, result, mid + 1, right)
	// 把分解的结果合并到result中
	merge(data, result, left, right)
}

// 合并
// 分解后的序列进行合并
// 在合并的过程中进行排序
func merge(data, result []int, left, right int) {
	// 通过求中间值获得做序列和有序列的边界
	mid := (right + left) >> 1
	// l是左边序列的起点下标
	// r是右边序列的起点下标
	l, r := left, mid + 1
	var i int // result的下标
	// 把每次合并的有序序列填充到result中
	// 每次都是从第一个元素开始填充，覆盖之前的结果
	for i = 0; l <= mid && r <= right; i++ {
		// 小的放在前面，这里的=保证了排序的稳定性
		if data[l] <= data[r] {
			result[i] = data[l]
			l++
		} else {
			result[i] = data[r]
			r++
		}
	}
	if l <= mid { // 填充左序列剩余数据
		for l <= mid {
			result[i] = data[l]
			l++
			i++
		}
	}
	if r <= right { // 填充右序列剩余数据
		for r <= right {
			result[i] = data[r]
			r++
			i++
		}
	}
	// 将数据按顺序从result还原至data中
	for i, l := 0, left; l <= right; i++ {
		data[l] = result[i]
		l++
	}
}

// 归并排序go特色版
//
// 利用切片的特性，简化了代码量
// 同时提高了可阅读性
// 但由于每次合并都是重新申请一个切片
// 而且还多次调用append函数
// 导致内存分配频繁
func MergeSortSpecial(data []int) []int {
	n := len(data)
	if n <= 1 { // 递归退出条件
		return data
	}
	// 找到中间下标
	mid := n >> 1
	// 利用切片特性获取左边序列并递归分解
	left := MergeSortSpecial(data[:mid])
	// 利用切片特性获取右边序列并递归分解
	right := MergeSortSpecial(data[mid:])
	// 合并并返回新的切片
	return mergeSpecial(left, right)
}

// 有序合并
func mergeSpecial(left, right []int) (result []int) {
	l, r := 0, 0
	ln, rn := len(left), len(right)
	for l < ln && r < rn {
		// 将left和right中较小的追加到result中
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	if l < ln { // 填充左序列剩余数据
		result = append(result, left[l:]...)
	}
	if r < rn { // 填充右序列剩余数据
		result = append(result, right[r:]...)
	}
	return
}
