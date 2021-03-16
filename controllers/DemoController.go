package controllers

import (
	"encoding/json"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/elasticClient"
	lib "github.com/EDDYCJY/go-gin-example/pkg/librarys"
	"github.com/EDDYCJY/go-gin-example/pkg/redisClient"
	"github.com/gin-gonic/gin"
)


type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int
	PageNum  int
	PageSize int
}

func DemoTest(c *gin.Context) {
	/*
	//str := "2006-01-02 15:04:05"
		str := "Y-m-d h:i:s"
		//str := "ymdhis"
		dateTime := NowTime(str)
	 */
	//response := lib.Gin{C: c}

	scoreMap := make(map[string]interface{})
	scoreMap["user"] = 90
	scoreMap["test"] = 100
	scoreMap["fff"] = 60


	funcName := c.Query("funcName")
	if funcName == "log"{
		lib.LogInfo(11,"log 33311",scoreMap)
		c.JSON(200, gin.H{
			"message": "log",
		})
		return
	}
	if funcName == "elastic"{
		esClient := new(elasticClient.EsClientHandler)
		updateMap := make(map[string]interface{})
		updateMap["challenge_item"] = 90
		esClient.EsClientUpdateById("90856","vaffle-posts","posts",updateMap)
		result:= esClient.EsClientGetInfoById("90856","vaffle-posts","posts")
		lib.ResponseJson(c,200,result,"")
		return
	}
	if funcName == "test"{
		c.JSON(200, gin.H{
			"message": "demoTest",
		})
		return
	}

	if funcName == "tagmodel"{
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		tags, _ := models.GetTags(1, 1, maps)
		lib.ResponseJson(c,200,tags,"")
		return
	}

	if funcName == "redis"{
		redisClient := new(redisClient.RedisClientHandler)
		stringCache,_:= json.Marshal(scoreMap)
		redisClient.RedisSelect(1).RedisSet("goin3",stringCache)
		cache := redisClient.RedisSelect(1).RedisGet("goin3")
		scoreMap2 := make(map[string]interface{})
		json.Unmarshal([]byte(cache), &scoreMap2)
		lib.ResponseJson(c,200,scoreMap2,"redis")
	}



	//logging.Info(11,22,global.VAFFLE_DB)
	/*
	lib.P(global.VAFFLE_DB)
	var idsString="[1]"
	var ids = make([]int, 0)
	json.Unmarshal([]byte(idsString), &ids)
	//var data = make([]goods.VapeShops, 0)
	var data = make([]models.Tag, 0)
	global.VAFFLE_DB.Where("id in (?)", ids).Find(&data)

	 */




}
