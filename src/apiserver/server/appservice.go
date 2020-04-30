package httpservice

import (
	"appmng/src/apiserver/resource"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppInfo struct {
	// 首字母必须大写，否则无法绑定
	AppId string `json:"AppId"`
	ApplyName string `json:"ApplyName" binding:"required"`
	ApplyType string `json:"ApplyType" binding:"required"`
	ApplyFrame string `json:"ApplyFrame" binding:"required"`
	ApplyEnvironment string `json:"ApplyEnvironment" binding:"required"`
	ApplyBrief string `json:"ApplyBrief"`
}

func AppList(c *gin.Context){
	// 中间件或处理程序中启动新的Goroutines时（简单理解，另起线程），则不应该使用其中的原始上下文，你必须使用只读副本（c.Copy()）
	// cCp := c.Copy()

	c.String(http.StatusOK, "%s", "")

}

func AppByid(c *gin.Context){
	// 中间件或处理程序中启动新的Goroutines时（简单理解，另起线程），则不应该使用其中的原始上下文，你必须使用只读副本（c.Copy()）
	// cCp := c.Copy()
	//id := c.Param("id")

	c.String(http.StatusOK, "%s", "")
}

func AddApp(c *gin.Context)  {
	var appInfo AppInfo
	err := c.BindJSON(&appInfo)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Data Error"})
		return
	} else {

		c.String(http.StatusOK, "%s", "")
	}
}

func UpdateApp(c *gin.Context)  {
	var appInfo AppInfo
	err := c.BindJSON(&appInfo)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Data Error"})
		return
	} else {

		c.String(http.StatusOK, "%s", "")
	}
}

// http://localhost:9090/iot/delApp/id
func DelApp(c *gin.Context){

	//id := c.Param("id")


	c.String(http.StatusOK, "%s", "")
	// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})

}

func GetUserinfo(c *gin.Context){
	resource.GetUserinfo()
}

func GetAppinfo(c *gin.Context){
	resource.GetAppinfo()
}