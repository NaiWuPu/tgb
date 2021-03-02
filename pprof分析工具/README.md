###go-torch和火焰图
#####火焰图（Flame Graph）是 Bredan Gregg 创建的一种性能分析图表，因为它的样子近似 🔥而得名。上面的 profiling 结果也转换成火焰图，如果对火焰图比较了解可以手动来操作，不过这里我们要介绍一个工具：go-torch。这是 uber 开源的一个工具，可以直接读取 golang profiling 数据，并生成一个火焰图的 svg 文件。

##安装go-torch
```cmd
   go get -v github.com/uber/go-torch
```
####火焰图 svg 文件可以通过浏览器打开，它对于调用图的最优点是它是动态的：可以通过点击每个方块来 zoom in 分析它上面的内容。

####火焰图的调用顺序从下到上，每个方块代表一个函数，它上面一层表示这个函数会调用哪些函数，方块的大小代表了占用 CPU 使用的长短。火焰图的配色并没有特殊的意义，默认的红、黄配色是为了更像火焰而已。

####go-torch 工具的使用非常简单，没有任何参数的话，它会尝试从
http://localhost:8080/debug/pprof/profile
####获取 profiling 数据。它有三个常用的参数可以调整：
```cmd
-u –url：要访问的 URL，这里只是主机和端口部分
-s –suffix：pprof profile 的路径，默认为 /debug/pprof/profile
–seconds：要执行 profiling 的时间长度，默认为 30s
```
###安装 FlameGraph
####要生成火焰图，需要事先安装 FlameGraph工具，这个工具的安装很简单（需要perl环境支持），只要把对应的可执行文件加入到环境变量中即可。
####下载安装perl：https://www.perl.org/get.html
####下载FlameGraph：git clone https://github.com/brendangregg/FlameGraph.git
####将FlameGraph目录加入到操作系统的环境变量中。
#####Windows平台的同学，需要把go-torch/render/flamegraph.go文件中的GenerateFlameGraph按如下方式修改，然后在go-torch目录下执行go install即可。
```cmd
// GenerateFlameGraph runs the flamegraph script to generate a flame graph SVG. func GenerateFlameGraph(graphInput []byte, args ...string) ([]byte, error) {
flameGraph := findInPath(flameGraphScripts)
if flameGraph == "" {
	return nil, errNoPerlScript
}
if runtime.GOOS == "windows" {
	return runScript("perl", append([]string{flameGraph}, args...), graphInput)
}
  return runScript(flameGraph, args, graphInput)
}
```

####压测工具wrk
#####推荐使用https://github.com/wg/wrk 或 https://github.com/adjust/go-wrk

####使用go-torch
####使用wrk进行压测:
```cmd
go-wrk -n 50000 http://127.0.0.1:8080/book/list
```
####在上面压测进行的同时，打开另一个终端执行:
```cmd
go-torch -u http://127.0.0.1:8080 -t 30
```
####30秒之后终端会初夏如下提示：Writing svg to torch.svg

####然后我们使用浏览器打开torch.svg就能看到如下火焰图了。 火焰图 火焰图的y轴表示cpu调用方法的先后，x轴表示在每个采样调用时间内，方法所占的时间百分比，越宽代表占据cpu时间越多。通过火焰图我们就可以更清楚的找出耗时长的函数调用，然后不断的修正代码，重新采样，不断优化。

####此外还可以借助火焰图分析内存性能数据：
```cmd
go-torch -inuse_space http://127.0.0.1:8080/debug/pprof/heap
go-torch -inuse_objects http://127.0.0.1:8080/debug/pprof/heap
go-torch -alloc_space http://127.0.0.1:8080/debug/pprof/heap
go-torch -alloc_objects http://127.0.0.1:8080/debug/pprof/heap
pprof与性能测试结合
```
####go test命令有两个参数和 pprof 相关，它们分别指定生成的 CPU 和 Memory profiling 保存的文件：

####-cpuprofile：cpu profiling 数据要保存的文件地址
####-memprofile：memory profiling 数据要报文的文件地址
####我们还可以选择将pprof与性能测试相结合，比如：

####比如下面执行测试的同时，也会执行 CPU profiling，并把结果保存在 cpu.prof 文件中：

go test -bench . -cpuprofile=cpu.prof
####比如下面执行测试的同时，也会执行 Mem profiling，并把结果保存在 cpu.prof 文件中：
```cmd
go test -bench . -memprofile=./mem.prof
```
####需要注意的是，Profiling 一般和性能测试一起使用，这个原因在前文也提到过，只有应用在负载高的情况下 Profiling 才有意义。