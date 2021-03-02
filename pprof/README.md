命令行交互界面
我们使用go工具链里的pprof来分析一下。
```cmd
go tool pprof cpu.pprof
```
执行上面的代码会进入交互界面如下：
``` 
runtime_pprof $ go tool pprof cpu.pprof
Type: cpu
Time: Jun 28, 2019 at 11:28am (CST)
Duration: 20.13s, Total samples = 1.91mins (568.60%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)  
```

我们可以在交互界面输入top3来查看程序中占用CPU前3位的函数：
``` 
(pprof) top3
Showing nodes accounting for 100.37s, 87.68% of 114.47s total
Dropped 17 nodes (cum <= 0.57s)
Showing top 3 nodes out of 4
      flat  flat%   sum%        cum   cum%
    42.52s 37.15% 37.15%     91.73s 80.13%  runtime.selectnbrecv
    35.21s 30.76% 67.90%     39.49s 34.50%  runtime.chanrecv
    22.64s 19.78% 87.68%    114.37s 99.91%  main.logicCode
```

其中：

flat：当前函数占用CPU的耗时
flat：:当前函数占用CPU的耗时百分比
sun%：函数占用CPU的耗时累计百分比
cum：当前函数加上调用当前函数的函数占用CPU的总耗时
cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
最后一列：函数名称
在大多数的情况下，我们可以通过分析这五列得出一个应用程序的运行情况，并对程序进行优化。

我们还可以使用list 函数名命令查看具体的函数分析，例如执行list logicCode查看我们编写的函数的详细分析。
``` 
(pprof) list logicCode
Total: 1.91mins
ROUTINE ================ main.logicCode in .../runtime_pprof/main.go
    22.64s   1.91mins (flat, cum) 99.91% of Total
         .          .     12:func logicCode() {
         .          .     13:   var c chan int
         .          .     14:   for {
         .          .     15:           select {
         .          .     16:           case v := <-c:
    22.64s   1.91mins     17:                   fmt.Printf("recv from chan, value:%v\n", v)
         .          .     18:           default:
         .          .     19:
         .          .     20:           }
         .          .     21:   }
         .          .     22:}
```
通过分析发现大部分CPU资源被17行占用，我们分析出select语句中的default没有内容会导致上面的case v:=<-c:一直执行。我们在default分支添加一行time.Sleep(time.Second)即可。

图形化
或者可以直接输入web，通过svg图的方式查看程序中详细的CPU占用情况。 想要查看图形化的界面首先需要安装graphviz图形化工具。

Mac：

brew install graphviz
Windows: 下载graphviz 将graphviz安装目录下的bin文件夹添加到Path环境变量中。 在终端输入dot -version查看是否安装成功。