FROM golang:1.14.4

ENV TZ Asia/Shanghai
RUN echo 'Asia/Shanghai' >/etc/timezone

ENV GOROOT /usr/local/go
ENV GOPROXY https://goproxy.cn

ENV GO111MODULE on
WORKDIR $GOPATH
ADD . $GOPATH/src/ex1

# excel 表格导出权限
#RUN chmod +x $GOPATH/src/时间工具包/*
# 日志权限
#RUN chmod -R 777 $GOPATH/src/时间工具包/logs/
RUN cd $GOPATH/src/ex1/ && go mod vendor
RUN go run ./main.go