package calc

import (
	"errors"
	"fmt"
)

// Stats 统计信息
type Stats struct {
	Operations int // 总操作次数
	Success    int // 成功次数
	Failures   int // 失败次数
}

// Calculator 计算器结构
type Calculator struct {
	stats Stats
}

// NewCalculator 创建新的计算器
func NewCalculator() *Calculator {
	return &Calculator{
		stats: Stats{},
	}
}

// Add 加法
func (c *Calculator) Add(a, b int) (int, error) {
	c.stats.Operations++
	result := a + b
	c.stats.Success++

	// 记录日志
	fmt.Printf("[DEBUG] 执行加法: %d + %d = %d\n", a, b, result)

	return result, nil
}

// Subtract 减法
func (c *Calculator) Subtract(a, b int) (int, error) {
	c.stats.Operations++
	result := a - b
	c.stats.Success++

	fmt.Printf("[DEBUG] 执行减法: %d - %d = %d\n", a, b, result)

	return result, nil
}

// Multiply 乘法
func (c *Calculator) Multiply(a, b int) (int, error) {
	c.stats.Operations++
	result := a * b
	c.stats.Success++

	fmt.Printf("[DEBUG] 执行乘法: %d × %d = %d\n", a, b, result)

	return result, nil
}

// Divide 除法
func (c *Calculator) Divide(a, b int) (int, error) {
	c.stats.Operations++

	if b == 0 {
		c.stats.Failures++
		fmt.Printf("[ERROR] 除法错误: 除数不能为0 (%d / %d)\n", a, b)
		return 0, errors.New("除数不能为0")
	}

	result := a / b
	c.stats.Success++

	fmt.Printf("[DEBUG] 执行除法: %d / %d = %d\n", a, b, result)

	return result, nil
}

// GetStats 获取统计信息
func (c *Calculator) GetStats() Stats {
	return c.stats
}

// ResetStats 重置统计
func (c *Calculator) ResetStats() {
	c.stats = Stats{}
}
