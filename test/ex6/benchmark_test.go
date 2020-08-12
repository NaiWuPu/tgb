package ex6

import "testing"

func Benchmark_Add_TimerControl(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
