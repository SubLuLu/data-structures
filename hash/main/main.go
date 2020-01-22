package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

// 字典中单词总个数
const counter = 235886

func main() {
	// 100以内的素数
	primes := []uint64{2, 3, 5, 7,
		11, 13, 17, 19,
		23, 29,
		31, 37,
		41, 43, 47,
		53, 59,
		61, 67,
		71, 73, 79,
		83, 89, 97}
	// 字典数据
	data := readWords()

	fmt.Println("prime\t min\t max\t conflict\t rate\t")

	l := len(primes)
	for n := 0; n < l; n++ {
		conflictStatus(primes[n], data)
	}
}

// readWords 从文件中读取字典里的所有单词
func readWords() []string {
	// 2.4M words 文件并不大，直接读入内存
	fi, err := os.Open("./words")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer func() {
		// 最后关闭文件
		if err := fi.Close(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()
	// 事先已经知道单词个数
	data := make([]string, counter)
	br := bufio.NewReader(fi)
	for n := 0; ; n++ {
		// 逐行读入
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		data[n] = string(a)
	}
	return data
}

// conflictStatus 计算冲突情况
func conflictStatus(prime uint64, data []string) {
	// 用无符号的64位进行计算，防止越界变为负数
	var hc, max, min uint64
	// 最小hash值
	min = math.MaxUint64

	m := make(map[uint64]bool)
	for _, d := range data {
		hc = hashCode(prime, d)
		m[hc] = true
		if hc < min {
			min = hc
		}
		if hc > max {
			max = hc
		}
	}

	// 所有的hash值总数
	num := len(m)

	// 冲突率，冲突数/总数
	rate := float64(counter - num) / float64(counter) * 100

	fmt.Printf("%d\t %d\t %d\t %d\t %.8f%%\t\n", prime, min, max, counter - num, rate)

	m = nil
}

// HashCode 计算str的hash值
func hashCode(prime uint64, str string) uint64 {
	var hc uint64
	for _, ch := range str {
		hc = prime*hc + uint64(ch)
	}
	return hc
}
