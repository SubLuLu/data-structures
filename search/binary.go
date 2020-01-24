package search

import (
	"sort"
)

// 二分查找
// 二分查找要求序列是有序的，所以要先排序
func BinarySearch(data []int, val int) int {
	// 先排序
	sort.Ints(data)

	left, right := 0, len(data) - 1
	return BinaryRecursion(data, val, left, right)
}

// 递归二分查找
func BinaryRecursion(data []int, val, left, right int) int {
	if left > right || data[left] > val || data[right] < right {
		return -1
	}
	// 求中间下标
	mid := (left + right) >> 1
	// 取出中间值方便比较
	mv := data[mid]
	if val < mv {
		// 递归在左边查找
		return BinaryRecursion(data, val, left, mid - 1)
	} else if val > mv {
		// 递归在右边查找
		return BinaryRecursion(data, val, mid + 1, right)
	} else {
		// 已经找到
		return mid
	}
}

// 非递归二分查找
func Binary(data []int, val int) int {
	left, right := 0, len(data) - 1
	if val < data[left] || val > data[right] {
		return -1
	}
	for left <= right {
		mid := (left + right) >> 1
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
