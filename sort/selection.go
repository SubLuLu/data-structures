package sort

// 选择排序
// 1. 默认起始元素是最小的
// 2. 每轮循环找到一个最小的
// 3. 和起始元素进行交换
func Select(data Sort) {
	n := data.Len()
	// 最小值的下标
	var index int
	for i := 0; i < n; i++ {
		index = i // 每次找寻最小值时初始化起始下标
		for j := i; j < n; j++ {
			if data.Less(j, index) {
				index = j
			}
		}
		// 第一个就是最小值，则不需要交换
		if index != i {
			data.Swap(i, index)
		}
	}
}

// IntsSelect sorts a slice of ints in increasing order.
func IntsSelect(a []int) { Select(IntSlice(a)) }

// Float64sSelect sorts a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
func Float64sSelect(a []float64) { Select(Float64Slice(a)) }

// StringsSelect sorts a slice of strings in increasing order.
func StringsSelect(a []string) { Select(StringSlice(a)) }
