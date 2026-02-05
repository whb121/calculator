package main

import (
	calc "clc/calculator"
	"fmt"
)

func main() {
	// 创建计算器实例
	calculator := calc.NewCalculator()

	fmt.Println("=== 计算器演示 ===")

	// 加法测试
	result1, err1 := calculator.Add(10, 5)
	if err1 == nil {
		fmt.Printf("10 + 5 = %d\n", result1)
	}

	// 减法测试
	result2, err2 := calculator.Subtract(10, 5)
	if err2 == nil {
		fmt.Printf("10 - 5 = %d\n", result2)
	}

	// 除法测试（正常）
	result3, err3 := calculator.Divide(10, 2)
	if err3 == nil {
		fmt.Printf("10 / 2 = %d\n", result3)
	}

	// 除法测试（错误：除零）
	result4, err4 := calculator.Divide(10, 0)
	if err4 != nil {
		fmt.Printf("10 / 0 错误: %v\n", err4)
	}
	fmt.Printf("10 / 0 = %d\n", result4)

	// 统计功能
	stats := calculator.GetStats()
	fmt.Printf("\n=== 统计信息 ===\n")
	fmt.Printf("操作次数: %d\n", stats.Operations)
	fmt.Printf("成功次数: %d\n", stats.Success)
	fmt.Printf("失败次数: %d\n", stats.Failures)
}
