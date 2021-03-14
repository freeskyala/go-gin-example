package controllers

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	groupRoute1 := r.Group("/api")
	groupRoute1.Use()
	{
		groupRoute1.GET("/demoTest", DemoTest)
		//获取标签列表
		groupRoute1.GET("/tagIndex", TagIndex)
	}

	return r
}
