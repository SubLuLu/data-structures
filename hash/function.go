package hash

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Identity 直接定址法
// H(key) = a x key + b (a != 0)
// 适用于关键字连续的场景
func Identity(key Key, a, b int) (code int, err error) {
	if key == nil {
		err = errors.New("key can't be nil")
		return
	}
	if a == 0 {
		err = errors.New("a can't be zero")
		return
	}
	var hc int
	switch key.(type) {
	case int:
		// 如果是int类型，直接类型转换
		hc, _ = key.(int)
	default:
		// 其他类型，通过计算
		hc = hashCode(key)
	}
	code = a*hc + b
	return
}

// Division 除留余数法
// H(key) = key % p (p是小于表长的最大素数)
func Division(key Key, p int) (code int, err error) {
	if key == nil {
		err = errors.New("key can't be nil")
		return
	}
	if p == 0 {
		err = errors.New("p can't be zero")
		return
	}
	var hc int
	switch key.(type) {
	case int:
		// 如果是int类型，直接类型转换
		hc, _ = key.(int)
	default:
		// 其他类型，通过计算
		hc = hashCode(key)
	}
	code = hc % p
	return
}

// DigitalAnalysis 数字分析法
// 通过对关键字找规律，取均匀分布的几位进行计算
// 该方法没有固定的公式，不同的规律给出不同的计算方法
// 本种实现以电话号码类型为例，取连续的后四位数字
// start 数字从个位数开始，个位数为1
func DigitalAnalysis(key Key, start, end int) (code int, err error) {
	if key == nil {
		err = errors.New("key can't be nil")
		return
	}
	if start > end {
		err = errors.New("end must be bigger than start")
		return
	}
	var hc int
	switch key.(type) {
	case int:
		// 如果是int类型，直接类型转换
		hc, _ = key.(int)
		// 先除以10的start-1次方
		// 将得到的数再对10的(end-start+1)次方取模
		code = (hc / int(math.Pow10(start-1))) % int(math.Pow10(end-start+1))
	default:
		// 将key格式化为字符串
		keyStr := fmt.Sprintf("%s", key)
		// 按字符分割
		keyRune := []rune(keyStr)
		if end > len(keyRune) {
			err = errors.New("end out of range of key")
			return
		}
		if start > len(keyRune) {
			err = errors.New("start out of range of key")
			return
		}
		n := len(keyRune)
		// 截取指定位置的连续字符
		// 字符串下标从左开始计算
		for i := n - end; i < n-start; i++ {
			// 将每个字符的ASCII码(或Unicode码)值加起来
			code = code + int(keyRune[i])
		}
	}
	return
}

// MidSquare 平方取中法
// 当关键字的每一位取值都不够均匀的时候，使用平方法使关键字均匀分布
// 再取新生成的关键字的中间几位
func MidSquare(key Key, start, end int) (code int, err error) {
	if key == nil {
		err = errors.New("key can't be nil")
		return
	}
	if start > end {
		err = errors.New("end must be bigger than start")
		return
	}
	var hc int
	switch key.(type) {
	case int:
		// 如果是int类型，直接类型转换
		hc, _ = key.(int)
	default:
		hc = hashCode(key)
	}
	square := hc * hc
	// 先除以10的start-1次方
	// 将得到的数再对10的(end-start+1)次方取模
	code = (square / int(math.Pow10(start-1))) % int(math.Pow10(end-start+1))
	return
}

// Folding 折叠法
// size是区域大小
// 将关键字从左到右平均分割成几部分
// 然后将这几部分进行相加得到结果
func Folding(key Key, size int) (code int, err error) {
	if key == nil {
		err = errors.New("key can't be nil")
		return
	}
	if size <= 0 {
		err = errors.New("size must be positive")
		return
	}
	var hc int
	switch key.(type) {
	case int:
		// 如果是int类型，直接类型转换
		hc, _ = key.(int)
	default:
		hc = hashCode(key)
	}
	// 将hc格式化为字符串
	codeStr := fmt.Sprintf("%d", hc)
	n := len(codeStr)
	var partVal int // 临时存储每个区域的值
	for i := 0; i < n; i += size {
		// 如果有必要，还可以对区域值partVal进行必要的转换
		if i+size < n {
			partVal, err = strconv.Atoi(codeStr[i : i+size])
		} else {
			partVal, err = strconv.Atoi(codeStr[i:n])
		}
		code += partVal
	}
	return
}

// hashCode 将其他类型的关键字转换成数值
func hashCode(key Key) int {
	var keyStr string
	if str, ok := key.(Stringer); ok {
		keyStr = str.String()
	} else {
		// 将key格式化为字符串
		keyStr = fmt.Sprintf("%s", key)
	}
	// 按字符分割
	keyRune := []rune(keyStr)
	hc := 0
	for _, ch := range keyRune {
		// 通常根据Horner法则，计算一个多项式
		// 形式为：int(keyRune[0])*k^(n-1) + int(keyRune[1])*k^(n-2) + ... + int(keyRune[n-1])
		// 其中k是一个常数，一般选择一个大小适中的素数，根据除留余数法可知，素数会减少冲突
		// 其中n是字符数组的长度
		//
		// 此处为简单实现就简单的相加
		// 将每个字符的ASCII码(或Unicode码)值加起来
		hc = hc + int(ch)
	}
	return hc
}

// maxPrime 找到不大于source的最大素数
func maxPrime(source int) (prime int) {
	if source == 2 {
		return source
	}
	// 2以外的素数肯定是奇数，所以start确定为奇数
	var start int
	// 先判断source是奇数还是偶数
	if source%2 == 0 {
		start = source - 1
	} else {
		start = source
	}

	var end int
	var flag bool
	// 从奇数中挑选素数
	for i := start; i > 2; i -= 2 {
		flag = false
		// 一个数如果能够分解成两个数相乘
		// 则一个因数肯定大于等于该数的平方根
		// 另一个因数肯定小于等于该数的平方根
		// 所以查找到该数的平方根即可
		end = int(math.Floor(math.Sqrt(float64(end))))
		for j := 3; j <= end; j += 2 {
			if i%j == 0 { // 找到因数
				flag = true
				break
			}
		}
		if !flag { // 没有找到因数
			prime = i
			return
		}
	}
	return
}
