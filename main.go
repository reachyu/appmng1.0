package main

import (
	"appmng/src/common"
	hs "appmng/src/httpserver"
	"log"
	"os"
)

func main() {

	// 初始化数据库
	dburl := "root@tcp(mysql:3306)/mlpipeline?maxAllowedPacket=0&charset=utf8&loc=Local&parseTime=True"
	// dburl := "root:20191014@tcp(172.16.4.175:3306)/yhb?charset=utf8"
	issucc := common.GetInstance().InitDataPool(dburl)
	if !issucc {
		log.Println("init database pool failure...")
		os.Exit(1)
	}

	// 开启http监听服务
	hs.InitHttpServer()
}