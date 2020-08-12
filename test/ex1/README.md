### -v 具体信息 =.所有基准测试 -benchtime 测试5秒
go test -v -bench=. -benchmem -benchtime=5s -gcflags "-N -l"

### 测试次数    响应时间
```
BenchmarkArray
BenchmarkArray-8          445617              2545 ns/op               0 B/op          0 allocs/op
BenchmarkSlice
BenchmarkSlice-8          300799              3833 ns/op            8192 B/op          1 allocs/op
```
