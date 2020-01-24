package search

import (
	"sort"
)

// 插值查找
// 插值查找要求序列是有序的，所以要先排序
// 对于均匀分布的有序序列，插值查找更加高效
func InterpolationSearch(data []int, val int) int {
	// 先排序
	sort.Ints(data)

	left, right := 0, len(data) - 1
	return InterpolationRecursion(data, val, left, right)
}

// 递归插值查找
func InterpolationRecursion(data []int, val, left, right int) int {
	if left > right || data[left] > val || data[right] < right {
		return -1
	}

	// 自适应查找mid位置
	// right - left 是当前需要查找的序列的长度
	// data[right] - data[left] 是当前需要查找的序列的极差
	// val - data[left] 是需要查找的值和最小元素的差
	// 当需要查找的序列均匀分布的时候
	// 把data[left], val, data[right]三个数放在数轴上看
	// 就可以确定val到data[left]的距离，并计算val的准确位置
	// 也就是
	// left + (right - left) * (val - data[left]) / (data[right] - data[left])
	// 或
	// right - (right - left) + (data[right] - val) / (data[right] - data[left])
	// 当然，这样需要确保val在data[left]和data[right]之间
	// 否则无意义，下标会越界
	mid := left + (right - left) * (val - data[left]) / (data[right] - data[left])
	mv := data[mid]

	if val < mv { // 在左边递归查找
		return InterpolationRecursion(data, val, left, mid - 1)
	} else if val > mv { // 在右边递归查找
		return InterpolationRecursion(data, val, mid + 1, right)
	} else { // 已经找到
		return mid
	}
}

// 非递归插值查找
func Interpolation9(data []int, val int) int {
	left, right := 0, len(data) - 1
	if val < data[left] || val > data[right] {
		return -1
	}
	for left <= right {
		mid := left + (right - left) * (val - data[left]) / (data[right] - data[left])
		mv := data[mid]
		if val < mv {
			right = mid - 1
		} else if val > mv {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
