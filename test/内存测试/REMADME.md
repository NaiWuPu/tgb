
### 测试内存
go test -v -bench=Alloc -benchmem benchmark_test.go

```
goos: windows
goarch: amd64
Benchmark_Alloc
Benchmark_Alloc-8       11459610               102 ns/op              16 B/op          2 allocs/op
PASS
ok      command-line-arguments  1.540s

```

####第 1 行的代码中-bench后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
####第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。