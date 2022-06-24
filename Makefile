#定义参数
NAME = private-test-project
VERSION = 1.0.0-SNAPSHOT

#执行终极target目标, 顺序执行image和start两个分目标
.PHONY: all
all: image start

#定义执行打镜像操作
.PHONY: image
image:
	docker build -t ${NAME}:${VERSION} .

#执行启动容器操作
.PHONY: start
start:
	docker run --name ${NAME} -p 8080:8080 -d ${NAME}:${VERSION}

