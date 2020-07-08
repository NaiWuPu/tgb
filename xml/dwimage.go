package main;

import (
	"github.com/astaxie/beego/logs"
	"io"
	"net/http"
)

func main() {
	controller.Ctx.ResponseWriter.Header().Add("Content-Type", "application/octet-stream")
	controller.Ctx.ResponseWriter.Header().Add("Content-Disposition", "attachment; filename=\"4375ea1a540adc5d6525d4a0b8d7c524.jpg\"")

	res, err := http.Get("https://jjc-api-server.oss-cn-beijing.aliyuncs.com/4375ea1a540adc5d6525d4a0b8d7c524.jpg")

	if err != nil{
		logs.Error(err)
	}
	//controller.Ctx.WriteString(res.Body)

	io.Copy(controller.Ctx.ResponseWriter, res.Body)
}