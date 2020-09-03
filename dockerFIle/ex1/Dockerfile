FROM golang:1.14.4

ENV TZ Asia/Shanghai
RUN echo 'Asia/Shanghai' >/etc/timezone

ENV GO111MODULE on

ENV GOROOT /usr/local/go
ENV GOPROXY https://goproxy.cn

WORKDIR $GOPATH
ADD . $GOPATH/src/ex1

# 日志权限
#RUN chmod -R 777 $GOPATH/src/时间工具包/logs/
RUN cd $GOPATH/src/ex1/ && go mod vendor

RUN cd $GOPATH/src/ex1 && go build -o ex1 ./main.go

CMD cd $GOPATH/src/ex1 && ./ex1