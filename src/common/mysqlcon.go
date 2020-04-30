package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

/*
* 数据库连接操作库
* 基于gorm封装开发
* 单例模式
 */
type MysqlConnectiPool struct {
}

var instance *MysqlConnectiPool
var once sync.Once

var db *gorm.DB
var err_db error

func GetInstance() *MysqlConnectiPool {
	once.Do(func() {
		instance = &MysqlConnectiPool{}
	})
	return instance
}

/*
* 初始化数据库连接
 */
func (m *MysqlConnectiPool) InitDataPool(dburl string) (issucc bool) {
	fmt.Println("=====================初始化mysql数据库=====================")
	db, err_db = gorm.Open("mysql", dburl)
	if err_db != nil {
		db = nil
		log.Fatal(err_db)
		return false
	}
	// 关闭数据库，db会被多个goroutine共享，可以不调用
	// defer后面的函数在defer语句所在的函数执行结束的时候会被调用
	// defer db.Close()
	return true
}

/*
* 对外获取数据库连接对象db
 */
func (m *MysqlConnectiPool) GetMysqlDB() (db_con *gorm.DB) {
	return db
}