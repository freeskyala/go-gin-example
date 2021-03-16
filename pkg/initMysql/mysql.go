package initMysql

import (
	"fmt"
	global2 "github.com/EDDYCJY/go-gin-example/pkg/global"
	"github.com/EDDYCJY/go-gin-example/pkg/initconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)

var TEST_DB_STATUS bool = false


type MysqlPool struct {

}


//初始化mysql连接池
func InitMysqlTestDb() {
	defer func() {
		if err := recover(); err != nil {
			//TODO panic happened, need log
		}
	}()
	//加锁，防止多个对象调用初始化方法
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	if TEST_DB_STATUS {
		return
	}

	db, err := gorm.Open(initconfig.DataTestConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		initconfig.DataTestConfig.User,
		initconfig.DataTestConfig.Password,
		initconfig.DataTestConfig.Host,
		initconfig.DataTestConfig.Name))
	if err != nil {
		//TODO log error
		return
	}
	//lib.P(555,initconfig.DataTestConfig,db)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "vape_" + defaultTableName;
	//}
	db.LogMode(false)
	global2.TEST_DB = db
	TEST_DB_STATUS = true
}


