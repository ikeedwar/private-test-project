#获取编译环境
FROM golang:1.18 as devel

#定义go mod环境变量
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

#执行创建文件夹操作
RUN mkdir -p /zjb/private-test-project

#设置工作目录
WORKDIR /zjb/private-test-project

#copy代码到工作目录
COPY . /zjb/private-test-project

#下载go mod的模块
RUN go mod download

#编译代码
RUN go build main.go

#暴露容器接口
EXPOSE 8080

#执行添加可执行权限
RUN chmod +x main

#执行main文件
ENTRYPOINT ["./main"]
