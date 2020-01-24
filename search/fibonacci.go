package search

import (
	"sort"
)

// 斐波那契查找
// 与二分查找类似，主要区别在于计算分割点位置
// 确定分割点思路：
// 1. 确保序列长度为F(k) - 1，F(k)为第k个斐波那契数
// 2. 如果长度不是F(k) - 1，则序列末尾用最大数进行填充补齐
// 3. 因为 F(k) = F(k-1) + F(k-2)
//    序列长度为F(k) - 1，分割点占一个，剩余为F(k) - 2
//    F(k) - 2 = F(k-1) + F(k-2) - 2 = F(k-1) - 1 + F(k-2) - 1
//    所以左边剩F(k-1) - 1，右边剩F(k-2) - 1
//    此时分割点的位置为 left + F(k-1) - 1
func FibonacciSearch(data []int, val int) int {
	// 先排序
	sort.Ints(data)
	n := len(data) // 原始序列长度
	// 根据需要排序序列的长度生成斐波那契数列
	fibs := buildFibs(n)
	// 斐波那契数列的最大下标
	k := len(fibs) - 1
	src := data
	// 如果原始序列长度不是F(k)-1，进行填充
	if n != fibs[k] - 1 {
		src = fillData(data, fibs[k])
	}
	// 初始化left和right
	left, right := 0, fibs[k] - 1
	// 递归查找
	return FibonacciRecursion(src, fibs, val, left, right, k)
}

// 递归斐波那契查找
// data  是长度为F(k)-1的有序序列
// fibs  是长度为k+1的斐波那契数列
// val   是需要查找的值
// left  是序列起始下标
// right 是序列结束下标
// k     是斐波那契数列下标
func FibonacciRecursion(data, fibs []int, val, left, right, k int) int {
	if left > right { // 退出递归条件
		return -1
	}
	// 计算分割点的位置
	// 因为总长度数 F(k) - 1
	// 分割后，左边长度是F(k-1) - 1
	// 所以分割点的下标 = 起始点的下标 + 左边序列长度
	// 即 mid = left + F(k-1) - 1
	// 同理，分割后，右边长度是F(k-2) - 1
	// 所以分割点的下标 = 结束点的下标 - 右边序列长度
	// 即 mid = right - F(k-2) + 1
	mid := left + fibs[k-1] - 1
	// 取出分割点的值
	mv := data[mid]
	if val < mv { // 比分割点值小，在左边递归查找
		k-- // 左边序列长度是F(k-1) - 1，所以k--
		// 修改right值为mid-1
		return FibonacciRecursion(data, fibs, val, left, mid - 1, k)
	} else if val > mv { // 比分割点值大，在右边递归查找
		k -= 2 // 右边序列长度是F(k-2) - 1，所以k -= 2
		// 修改left值为mid+1
		return FibonacciRecursion(data, fibs, val, mid + 1, right, k)
	} else { // 已经找到
		return mid
	}
}

// 非递归斐波那契查找
// data  是长度为F(k)-1的有序序列
// fibs  是长度为k+1的斐波那契数列
// val   是需要查找的值
// left  是序列起始下标
// right 是序列结束下标
func Fibonacci(data, fibs []int, val, left, right int) int {
	k := len(fibs) - 1
	for left <= right {
		// 计算分割点的位置
		// 因为总长度数 F(k) - 1
		// 分割后，左边长度是F(k-1) - 1
		// 所以分割点的下标 = 起始点的下标 + 左边序列长度
		// 即 mid = left + F(k-1) - 1
		// 同理，分割后，右边长度是F(k-2) - 1
		// 所以分割点的下标 = 结束点的下标 - 右边序列长度
		// 即 mid = right - F(k-2) + 1
		mid := left + fibs[k-1] - 1
		// 取出分割点的值
		mv := data[mid]
		if val < mv { // 比分割点值小，在左边进行查找
			right = mid - 1 // 更新right
			k-- // 左边序列长度是F(k-1) - 1，所以k--
		} else if val > mv { // 比分割点值大，在右边进行查找
			left = mid + 1 // 更新left
			k -= 2 // 右边序列长度是F(k-2) - 1，所以k -= 2
		} else { // 已经找到
			return mid
		}
	}
	return -1
}

// 构造斐波那契数列
func buildFibs(length int) []int {
	var fibs []int // 存储斐波那契数列
	fc := fibonacci() // 生成斐波那契数列的闭包函数
	// 初始化第一个斐波那契数
	fib := fc()
	// 存储第一个斐波那契数
	fibs = append(fibs, fib)
	// 序列长度大于斐波那契数，则生成下一个
	for length > fib - 1 {
		fib = fc() // 产生下一个斐波那契数
		// 将生成的斐波那契数存储到fibs中
		fibs = append(fibs, fib)
	}
	return fibs
}

// 当要排序的序列的长度小于F(k) - 1时
// 用序列中最大的数据填充至序列长度达到F(k) - 1
func fillData(data []int, fk int) []int {
	// 申请一个长度为fk-1的切片进行存储
	src := make([]int, fk - 1)
	// 将原始序列拷贝到src中
	copy(src, data)
	n := len(data)
	// 用原始序列中的最大值进行填充
	for j := n; j < fk - 1; j++ {
		src[j] = data[n-1]
	}
	return src
}

// 产生斐波那契数列
// 利用闭包产生斐波那契数列
func fibonacci() func() int {
	i := 0 // 记录斐波那契数列中的下标
	m, n := 1, 1 // 斐波那契数列的前两项
	return func () int {
		if i > 1 {
			m, n = n, m+n
		}
		i++
		return n
	}
}
