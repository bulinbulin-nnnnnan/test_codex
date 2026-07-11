// Package mathutil 提供基础数学计算工具函数
package mathutil

import "errors"

// Sum 返回任意数量整数的和
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Average 返回一组整数的平均值（浮点数）
func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := Sum(nums...)
	return float64(total) / float64(len(nums))
}

// Factorial 计算 n 的阶乘（n!），n 为负数时返回错误
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("负数没有阶乘")
	}
	if n <= 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// GCD 计算两个整数的最大公约数（欧几里得算法）
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Fibonacci 返回斐波那契数列的前 n 项
func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	seq := make([]int, n)
	seq[0], seq[1] = 0, 1
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}
