package sort

// 基数排序
//
// 基数排序的原理是：
// 1. 找出最大数，并求它是几位数
// 2. 先按个位数进行排序
// 3. 再按十位数进行排序
// 4. 以此类推，直到达到最大数的最高位为止
//
// 传统基数排序缺点：
// 1. 不能对负数进行排序
// 2. 频繁使用除法和取模运算，效率较低
// 3. 内存开销大
//
// 改进基数排序：
// 1. 使得其能对负数进行排序
// 2. 使用位运算代替除法和取模运算
// 3. 采用一些策略动态扩容
//
// 改进思路：
// 1. 10进制不方便进行位运算，采用16进制
// 2. 负数的二进制存储是其绝对值取反 + 1
//    如果把负数转换成无符号的正数，
//    那么负数的绝对值越小，其对应的无符号的正数越大
//    按照数的排序效果看，负数的绝对值越小排在越后面
//    而这正是负数的排序效果
// 3. 计算最高位
//    因为负数转换成无符号的正数后，最高位始终为1(原来的符号位)
//    正常来说，16进制下，最高位一直是16，显然不太科学
//    对于两个负数来说，相同的高位不需要比较，只需要比较不同的低位即可
//    低位的不同，是由于负数的绝对值取反 + 1造成的
//    所以低位不同的位数，就是其绝对值得二进制有效位数
// 4. 按照以上思路进行技术排序，得到一个序列
//    在该序列中，所有非负数是有序的，所有负数也是有序的
//    当最高位大于1时，左边是非负序列，右边为负数序列
// 5. 将序列拆分为非负数序列和负数序列
// 6. 将两个序列合并到原始序列中
func Radix(data []int64) {
	n := len(data)
	if n <= 1 {
		return
	}
	// 因为16进制，余数为0-15，所以需要16个桶
	bucket := 16
	// 先计算最大值和最小值
	min, max := minAndMax(data)
	// 标识是否整个序列都为负数
	isNegative := min < 0 && max < 0
	if min < 0 { // 最小值小于0，说明有负数
		// 求最小值的绝对值
		// 负数是其绝对值取反 + 1
		minAbs := ^(min - 1)
		// 求绝对值最大的值，将值赋给max
		if minAbs > max {
			max = minAbs
		}
	}
	// 求最大有多少位(即十六进制下是几位数)
	digit := digitHex(max)
	// 二维切片作为桶
	bins := make([][]int64, bucket)
	// 记录每个桶中放了多少个元素
	// 减轻遍历负担，避免多次遍历bins
	count := make([]int, bucket)
	// 循环digit轮
	for i := 0; i < digit; i++ {
		for _, v := range data {
			// 相当于v除以16^i
			t := v >> uint(i << 2)
			// 相当于t%16
			r := t & 0xF
			// 取出第r个切片桶
			bin := bins[r]
			if len(bin) == 0 { // 当前桶为空
				// 初始化切片桶
				bin = newBin(n)
				// 赋值到二维切片中
				bins[r] = bin
			}
			// 当前桶中元素个数
			l := count[r]
			if l == len(bin) { // 当前桶已满
				// 自动扩容
				bin = autoCopy(bin)
				// 赋值到二维切片中
				bins[r] = bin
			}
			bin[l] = v // 存入桶中
			count[r]++ // 计数器自增
		}
		// 将每轮排序结果放回data中，最后一轮count不需要清零
		adjust(data, bins, count, i != digit - 1)
	}

	// 如果序列中有负数也有非负数，则最后对负数进行一次排序
	if min < 0 && !isNegative {
		// 当digit大于等于2时，负数都在右边且有序
		order(data, digit)
	}
}

// 计算序列中的最小值和最大值
func minAndMax(data []int64) (int64, int64) {
	n := len(data)
	// 默认第一个即为最大值，也为最小值
	min, max := data[0], data[0]
	for i := 1; i < n; i++  {
		if data[i] > max {
			max = data[i]
		} else if data[i] < min {
			min = data[i]
		}
	}
	return min, max
}

// 按16进制计算max有几位
func digitHex(max int64) int {
	// 因为int64是64位的，每4位二进制表示16进制的一位，所以最大值是16
	for digit := 1; digit <= 16; digit++ {
		// 需要移动的位数
		// digit << 2 相当于digit * 4
		// 因为16余数为0-15，需要4位二进制表示
		i := uint(digit << 2)
		// 相当于 max / 16^digit
		// 如果商为0，表示最高位已经找到
		if max >> i == 0 {
			return digit
		}
	}
	return 1
}

// newBin 初始化一个桶
// 如果序列长度小于10
// 那么每个桶的长度为序列长度
// 否则桶的初始长度为10
func newBin(length int) []int64 {
	if length < 10 {
		return make([]int64, length)
	} else {
		return make([]int64, 10)
	}
}

// autoCopy 复制切片并扩容
// 自动扩容简单策略
// 当长度小于1000时，翻倍扩容
// 当长度大于1000时，增加100
func autoCopy(src []int64) []int64 {
	n := len(src)
	var result []int64
	if n < 1000 {
		result = make([]int64, n << 1)
	} else {
		result = make([]int64, n + 100)
	}
	copy(result, src)
	return result
}

// 对每轮排序后的结果进行必要调整
// 1. 多定义一个count计数器，是为了避免直接遍历bins
// 2. 因为count是一维的，所以只遍历它
// 3. 对count的元素清零即可，不需要对bins进行清零
//
// data  是需要排序的序列
// bins  是进行排序的桶
// count 是每个桶中元素个数计数
// clear 是表示count是否清零
func adjust(data []int64, bins [][]int64, count []int, clear bool) {
	n := len(data)
	for i := 0; i < n; {
		// j是桶的序号，c是桶中元素个数
		for j, c := range count {
			if c != 0 { // 桶中是否有元素
				for k := 0; k < c; k++ {
					// 将桶中元素依次放入原始序列
					data[i] = bins[j][k]
					i++
				}
				// 对计数器清零，避免对桶(二维切片)进行清零
				if clear {
					count[j] = 0
				}
			}
		}
	}
}

// 如果序列中有负数，最后把负数进行排序
// 经过排序后所得序列中非负数是有序的
// 经过排序后所得序列中负数也是有序的
// 只不过非负数和负数交替存储在序列中
// 所以用两个桶分别存储有序的负数和非负数序列
// 再将两个序列合并到原始序列中
func order(data []int64, digit int) {
	n := len(data)
	if digit > 1 {
		// 找到第一个负数的下标
		index := firstNegative(data)
		// 存储非负序列切片
		positive := make([]int64, index)
		// 存储非负序列
		copy(positive, data[:index])
		// 将负数序列copy到原序列的头部
		copy(data, data[index:])
		// 计算负数序列长度
		nl := n - index
		// 依次添加非负序列到原序列
		for i := nl; i < n; i++ {
			data[i] = positive[i - nl]
		}
	} else {
		// 存储负数的桶
		negative := newBin(n)
		// 存储非负数数的桶
		positive := newBin(n)
		var i, j int // i, j分别是负数和非负数个数
		for _, d := range data {
			if d < 0 { // 负数存入negative中
				if i == len(negative) {
					// 扩容，避免下标越界
					negative = autoCopy(negative)
				}
				negative[i] = d
				i++
			} else { // 费负数存入positive中
				if j == len(positive) {
					// 扩容，避免下标越界
					positive = autoCopy(positive)
				}
				positive[j] = d
				j++
			}
		}
		// 先把负数copy到data中
		copy(data, negative)
		// 再把非负数追加到后面
		for y := i; y < n; y++ {
			data[y] = positive[y - i]
		}
	}
}

// 找到第一个负数
// 因为负数序列在后面，所以从后面开始查找
func firstNegative(data []int64) int {
	n := len(data)
	for i := n - 2; i > 0; i-- {
		if data[i] >= 0 {
			return i + 1
		}
	}
	return n - 1
}

// 基础版的基数排序
//
// 只能对非负数进行排序
// 包含负数，会将负数按其绝对值进行排序
// 采用十进制取模的方式
// 需要排序的数据量大时可能内存溢出
func RadixBase(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}
	// 因为十进制，余数为0-9，所以需要10个桶
	bucket := 10
	// 序列中的最大值
	max := maxVal(data)
	// 序列中最大值的位数
	digit := digitDec(max)
	// 二维切片作为桶
	bins := make([][]int, bucket)
	// 记录每个桶中放了多少个元素
	// 减轻遍历负担，避免多次遍历bins
	count := make([]int, bucket)
	var t = 1
	// 循环digit轮
	for i := 0; i < digit; i++ {
		for _, v := range data {
			r := (v / t) % 10
			if r < 0 { // 防止模为负数，数组越界
				r = -r
			}
			// 取出第r个切片桶
			bin := bins[r]
			if len(bin) == 0 { // 当前桶为空
				// 初始化切片桶
				bin = make([]int, n)
				// 赋值到二维切片中
				bins[r] = bin
			}
			// 当前桶中元素个数
			l := count[r]
			bin[l] = v // 存入桶中
			count[r]++ // 计数器自增
		}
		t *= 10 // 一轮循环完后，t*10
		// 将每轮排序结果放回data中，最后一轮count不需要清零
		arrange(data, bins, count, i != digit - 1)
	}
}

// 计算序列中的最大值
func maxVal(data []int) int {
	n := len(data)
	// 默认第一个即为最大值
	max := data[0]
	for i := 1; i < n; i++  {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// 按十进制计算max有几位
func digitDec(max int) int {
	i := 1
	for digit := 1; ; digit++ {
		i *= 10
		if max / i == 0 {
			return digit
		}
	}
}

// 整理桶中的数据，按顺序放入原序列中
func arrange(data []int, bins [][]int, count []int, clear bool) {
	n := len(data)
	for i := 0; i < n; {
		// j是桶的序号，c是桶中元素个数
		for j, c := range count {
			if c != 0 { // 桶中是否有元素
				for k := 0; k < c; k++ {
					// 将桶中元素依次放入原始序列
					data[i] = bins[j][k]
					i++
				}
				// 对计数器清零，避免对桶(二维切片)进行清零
				if clear {
					count[j] = 0
				}
			}
		}
	}
}
