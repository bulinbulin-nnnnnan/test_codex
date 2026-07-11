package main

import (
	"fmt"
	"test_codex/mathutil"
)

func main() {
	// 求和
	fmt.Println("=== 数学工具包示例 ===")
	fmt.Println("Sum(1, 2, 3, 4, 5)        =", mathutil.Sum(1, 2, 3, 4, 5))

	// 平均值
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Average([10,20,30,40,50]) =", mathutil.Average(nums))

	// 阶乘
	if f, err := mathutil.Factorial(6); err == nil {
		fmt.Println("Factorial(6)               =", f)
	}
	if _, err := mathutil.Factorial(-3); err != nil {
		fmt.Println("Factorial(-3)              = 错误:", err)
	}

	// 最大公约数
	fmt.Println("GCD(48, 18)                =", mathutil.GCD(48, 18))

	// 斐波那契数列
	fmt.Println("Fibonacci(10)              =", mathutil.Fibonacci(10))
}
