##linux 下
```
$ go build -o cpu cpu.go
$ ./cpu
$ go tool pprof --pdf cpu cpu.pprof > cpu.pdf
```
###代码说明如下：
#####第 1 行将 cpu.go 编译为可执行文件 cpu。
#####第 2 行运行可执行文件，在当前目录输出 cpu.pprof 文件。
#####第 3 行，使用 go tool 工具链输入 cpu.pprof 和 cpu 可执行文件，生成 PDF 格式的输出文件，将输出文件重定向为 cpu.pdf 文件。这个过程中会调用 Graphviz 工具，Windows 下需将 Graphviz 的可执行目录添加到环境变量 PATH 中。