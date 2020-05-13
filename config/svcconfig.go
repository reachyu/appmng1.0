package config

import (
	hs "appmng/src/apiserver/server"

	"github.com/gin-gonic/gin"
)

/*
配置对外提供服务的访问名及其对应的实现方法
*/

var GetMaps = map[string]gin.HandlerFunc{
	"app/list": hs.AppList,
	"app/info/:id": hs.AppByid,
	"app/getusers": hs.GetUserinfo,
	"app/getapps": hs.GetAppinfo,
}

var PostMaps = map[string]gin.HandlerFunc{
	// id
	"app/del/:id": hs.DelApp,
	"app/add": hs.AddApp,
	"app/update/:id": hs.UpdateApp,
}
