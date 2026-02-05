package calc

import (
	"fmt"
	"testing"
)

// TestAdd 测试加法
func TestAdd(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"正数加法", 2, 3, 5},
		{"负数加法", -2, -3, -5},
		{"正负相加", 5, -3, 2},
		{"零加法", 0, 10, 10},
		{"大数加法", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Add(tt.a, tt.b)
			if err != nil {
				t.Errorf("%s: Add(%d, %d) 返回错误: %v", tt.name, tt.a, tt.b, err)
			}
			if result != tt.expected {
				t.Errorf("%s: Add(%d, %d) = %d, 期望 %d", tt.name, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract 测试减法
func TestSubtract(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"正数减法", 5, 3, 2},
		{"负数减法", -2, -3, 1},
		{"正负相减", 5, -3, 8},
		{"零减法", 10, 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Subtract(tt.a, tt.b)
			if err != nil {
				t.Errorf("%s: 返回错误: %v", tt.name, err)
			}
			if result != tt.expected {
				t.Errorf("%s: Subtract(%d, %d) = %d, 期望 %d", tt.name, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply 测试乘法
func TestMultiply(t *testing.T) {
	calc := NewCalculator()

	result, err := calc.Multiply(4, 5)
	if err != nil {
		t.Errorf("Multiply 返回错误: %v", err)
	}
	if result != 20 {
		t.Errorf("Multiply(4, 5) = %d, 期望 20", result)
	}
}

// TestDivide 测试除法
func TestDivide(t *testing.T) {
	calc := NewCalculator()

	// 测试正常除法
	result, err := calc.Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) 返回错误: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d, 期望 5", result)
	}

	// 测试除零错误
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) 应该返回错误，但没有")
	}
	if err.Error() != "除数不能为0" {
		t.Errorf("期望错误信息 '除数不能为0'，得到 '%v'", err)
	}
}

// TestStats 测试统计功能
func TestStats(t *testing.T) {
	calc := NewCalculator()

	// 执行一些操作
	calc.Add(1, 2)
	calc.Subtract(5, 3)
	calc.Multiply(2, 3)
	calc.Divide(10, 2)
	calc.Divide(5, 0) // 这个会失败

	stats := calc.GetStats()

	if stats.Operations != 5 {
		t.Errorf("总操作次数 = %d, 期望 5", stats.Operations)
	}
	if stats.Success != 4 {
		t.Errorf("成功次数 = %d, 期望 4", stats.Success)
	}
	if stats.Failures != 1 {
		t.Errorf("失败次数 = %d, 期望 1", stats.Failures)
	}

	// 测试重置
	calc.ResetStats()
	stats = calc.GetStats()
	if stats.Operations != 0 || stats.Success != 0 || stats.Failures != 0 {
		t.Errorf("重置后统计应该全为0，得到 %+v", stats)
	}
}

// BenchmarkAdd 性能测试：加法
func BenchmarkAdd(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(i, i+1)
	}
}

// ExampleAdd 示例测试
func ExampleAdd() {
	calc := NewCalculator()
	result, _ := calc.Add(3, 4)
	fmt.Println(result)
	// Output:
	// [DEBUG] 执行加法: 3 + 4 = 7
	// 7
}
