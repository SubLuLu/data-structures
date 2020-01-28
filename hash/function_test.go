package hash

import (
	"fmt"
	"testing"
)

// 测试散列算法

type bornNum struct {
	Year int // 年份
	Num  int // 人数
}

// 测试直接定址法
// 例如：统计某个村2010年及以后出生的人数
// 年份   人数
// 2010  48
// 2011  63
// 2012  70
// 2013  51
// 2014  37
// 2015  62
// 2016  30
// 2017  49
// 2018  66
// 2019  54
// 以年份为关键字，H(key) = a x key + b
// 令a=1，b=-2010
// H(key) = key - 2010
func TestIdentity(t *testing.T) {
	bns := make([]bornNum, 10)

	var (
		index int
		err   error
	)

	bn2010 := bornNum{
		Year: 2010,
		Num:  48,
	}
	index, err = Identity(bn2010.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}

	bns[index] = bn2010

	bn2011 := bornNum{
		Year: 2011,
		Num:  63,
	}
	index, err = Identity(bn2011.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2011

	bn2012 := bornNum{
		Year: 2012,
		Num:  70,
	}
	index, err = Identity(bn2012.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2012

	bn2013 := bornNum{
		Year: 2013,
		Num:  51,
	}
	index, err = Identity(bn2013.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2013

	bn2014 := bornNum{
		Year: 2014,
		Num:  37,
	}
	index, err = Identity(bn2014.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2014

	bn2015 := bornNum{
		Year: 2015,
		Num:  62,
	}
	index, err = Identity(bn2015.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2015

	bn2016 := bornNum{
		Year: 2016,
		Num:  30,
	}
	index, err = Identity(bn2016.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2016

	bn2017 := bornNum{
		Year: 2017,
		Num:  49,
	}
	index, err = Identity(bn2017.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2017

	bn2018 := bornNum{
		Year: 2018,
		Num:  66,
	}
	index, err = Identity(bn2018.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2018

	bn2019 := bornNum{
		Year: 2019,
		Num:  54,
	}
	index, err = Identity(bn2019.Year, 1, -2010)
	if err != nil {
		t.Fatal(err)
	}
	bns[index] = bn2019

	fmt.Println("期望是： [{2010 48} {2011 63} {2012 70} {2013 51} {2014 37} {2015 62} {2016 30} {2017 49} {2018 66} {2019 54}]")

	fmt.Println("结果为：", bns)
}

// 测试除留余数法
// 例如：[12 24 36 48 60 72 84 96 108 120 132 144]
// 散列表表长为12
// 如果p取表长12，则所有的都冲突
// 如果p取不大于表长的最大素数11，则仅12和144冲突
// H(key) = key % 11
func TestDivision(t *testing.T) {
	nums := []int{12, 24, 36, 48, 60, 72, 84, 96, 108, 120, 132, 144}
	m := len(nums)   // 表长
	p := maxPrime(m) // 最大素数

	var err error
	indexes := make([]int, 12)
	for i, num := range nums {
		indexes[i], err = Division(num, p)
		if err != nil {
			t.Fatal(err)
		}
	}

	fmt.Println("期望是： [1 2 3 4 5 6 7 8 9 10 0 1]")
	fmt.Println("结果是：", indexes)
}

// 测试数字分析法
// 例如，将电话号码作为关键字
// 电话号码的前三位是表示运营商
// 中间四位是表示地区
// 也就是前七位都不是均匀分布的
// 所以用后四位进行计算
// 例如：有如下伪手机号
// 13011224856
// 13622339842
// 13833446263
// 15244557109
// 15855663981
// 17066771562
// 17377882295
// 17788997480
// 18199003861
// 18900115240
// 显然，对于手机号来说，只有最后四位数字是均匀分布的
// 所以通过对数字的分析，选择最后四位作为散列地址
func TestDigitalAnalysis(t *testing.T) {
	numbers := []int{
		13011224856,
		13622339842,
		13833446263,
		15244557109,
		15855663981,
		17066771562,
		17377882295,
		17788997480,
		18199003861,
		18900115240,
	}
	addrs := make([]int, 10)
	var err error
	for i, number := range numbers {
		addrs[i], err = DigitalAnalysis(number, 1, 4)
		if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println("期望为： [4856 9842 6263 7109 3981 1562 2295 7480 3861 5240]")
	fmt.Println("结果是：", addrs)
}

// 测试平方取中法
// 例如：有如下一组数字
// [1234 1243 1324 1342 1423 1432 2134 2143 2314 2341 2413 2431]
// 所有数字的位数不是很大，都是四位数，如果再添加其他关键字，数字的分布情况就不清楚
// 此种情况适合对关键字平法，然后取中间的几位作为散列地址
// 平方结果为
// [1522756 1545049 1752976 1800964 2024929 2050624]
// [4553956 4592449 5354596 5480281 5822569 5909761]
func TestMidSquare(t *testing.T) {
	keys := []int{1234, 1243, 1324, 1342, 1423, 1432,
		2134, 2143, 2314, 2341, 2413, 2431}

	addrs := make([]int, 12)
	var err error
	for i, key := range keys {
		addrs[i], err = MidSquare(key, 3, 5)
		if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println("期望为： [227 450 529 9 249 506 539 924 545 802 225 97]")
	fmt.Println("结果是：", addrs)
}

// 测试折叠法
// 例如：有如下一组五位数作为关键字
// [66142 70079 65250 17403 34835 94333 64882 99956 43780 76672]
// 折叠法将每个关键字都从左到右平分为3个部分
// 然后将得到的数相加作为散列地址
// 比如：66142 被折叠法分为 66 14 2 三部分
// 再相加 66 + 14 + 2 = 82
func TestFolding(t *testing.T) {
	keys := []int{66142, 70079, 65250, 17403, 34835,
		94333, 64882, 99956, 43780, 76672}

	addrs := make([]int, 10)
	var err error
	for i, key := range keys {
		addrs[i], err = Folding(key, 2)
		if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println("期望为： [82 86 90 60 122 130 154 200 121 145]")
	fmt.Println("结果是：", addrs)
}
