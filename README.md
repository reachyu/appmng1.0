# 简介
* go语言开发的web应用，1.0版本前后端不分离
* 使用GIN框架实现http服务
* 部署到kubeflow，访问kubeflow数据库，实现增删改查功能
* go版本：1.13
* kubeflow版本：0.7.0
* 开发工具：IntelliJ IDEA 2019.2.3 x64
# 目录介绍
## config
配置目录
## public
静态文件目录
## src
### 本地服务
业务逻辑及api目录
## 访问示例
### 本地服务
* 前端  http://ip:9090/html
* api   http://ip:9090/app/getapps
### 部署kubeflow
* 前端  http://ip:31000/html
* api   http://ip:31000/app/getapps
# 编译
```
go build
```

# 启动
```
go run main.go
```