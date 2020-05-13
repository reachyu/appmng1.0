package httpserver

import (
	"appmng/config"
	"github.com/gin-gonic/gin"
	"log"
)

/*
开启http server
封装对外提供的接口服务
*/

var router *gin.Engine

func init() {
	// Engin gin.Default() 默认是加载了一些框架内置的中间件
	router = gin.Default()
	log.Println(">>>> HTTP Server starting <<<<")
	// gin.New() 没有默认加载框架内置的中间件，根据需要自己手动加载中间件
	//router := gin.New()

	// http://127.0.0.1:9090/html/test.html
	router.Static("/html", "./public")
}

func InitHttpServer() {

	postListen()
	getListen()
	// 指定地址和端口号
	// router.Run("localhost:9090")
	router.Run(":9090")
}

// 启动post服务监听
func postListen() {
	log.Println(">>>> start post servers <<<<")
	// post接口
	for u, f := range config.PostMaps {
		router.POST("/"+u, f)
	}
}

// 启动get服务监听
func getListen() {
	log.Println(">>>> start get servers <<<<")
	// get接口
	for u, f := range config.GetMaps {
		router.GET("/"+u, f)
	}
}