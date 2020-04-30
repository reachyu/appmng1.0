## 源镜像，使用了镜像大小体积只有5MB的alpine镜像
#FROM alpine:latest
## 引入最新的golan ，不设置版本即为最新版本
FROM golang
## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /www/webapp
## 设置程序在容器内的工作路径
WORKDIR /www/webapp
## 把上文编译好的main文件添加到镜像里，第一个参数是实际文件或路径，第二个参数是容器内路径
## ADD main /
COPY . /www/webapp
## 编译
#RUN go build .
## 暴露容器内部端口
EXPOSE 9090
## 入口，最后执行
#ENTRYPOINT ["/www/webapp/main"]
## 启动docker需要执行的文件
CMD go run main.go
