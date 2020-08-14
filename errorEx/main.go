package main

import (
	mylogger "examples/errorEx/时间工具包"
)

var log mylogger.Logger

func main() {
	log = mylogger.NewFileLogger("info", "./", "text.log", 10*1024)
	for {
		log.Debug("DEBUG 日志")
		log.Info("INFO 日志")
		log.Warning("warning 日志")

		id := 10010
		name := "理想"

		log.Error("Error 日志%d,name%s", id, name)
		log.Fatal("fata 日志")
		//time.Sleep(1 * time.Second)
	}
}
