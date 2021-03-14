package controllers

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	lib "github.com/EDDYCJY/go-gin-example/pkg/librarys"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
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
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	logging.Info(11,22,scoreMap)
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
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0


	//tags, _ := models.DemoTest(1, 1, maps)
	tags, _ := models.GetTags(1, 1, maps)
	//lib.P(tags)
	lib.ResponseJson(c,200,tags,"")
	//response.Response(200, 200, scoreMap)
	return
	gredis.Set("test2", scoreMap, 3600)
	lib.Pmap(scoreMap)
	lib.P(222,scoreMap)
	c.JSON(200, gin.H{
		"message": "demoTest",
	})
}
