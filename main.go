package main

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/controllers"
	"github.com/EDDYCJY/go-gin-example/pkg/initconfig"
	"github.com/EDDYCJY/go-gin-example/pkg/initstart"
	"github.com/EDDYCJY/go-gin-example/pkg/librarys"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	//setting.Setup()
	initconfig.InitConfig()
	initstart.InitMysqlTestDb()
	new(initstart.RedisPool).InitRedisDb()
	new(initstart.ElasticSearch).InitDefaultEs()
	librarys.InitLog()
}

// @title Golang Gin API test
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(initconfig.ServerConfig.RunMode)
	routersInit := controllers.InitRouter()
	readTimeout := initconfig.ServerConfig.ReadTimeout
	//lib.P(readTimeout,initconfig.ServerConfig.ReadTimeout)
	writeTimeout := initconfig.ServerConfig.WriteTimeout
	endPoint := fmt.Sprintf(":%d", initconfig.ServerConfig.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
