package search

// 线性查找
func Linear(data []int, val int) int {
	for i, v := range data {
		if v == val {
			return i
		}
	}
	return -1
}
