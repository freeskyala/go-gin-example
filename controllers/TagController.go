package controllers

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/logics"
	lib "github.com/EDDYCJY/go-gin-example/pkg/librarys"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func TagIndex(c *gin.Context) {

	name := c.Query("name")
	fmt.Println(name)
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := TagLogic.Tag{
		Name:     name,
		State:    state,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		lib.ResponseJson(c,500,nil,"")
		return
	}

	count, err := tagService.Count()
	if err != nil {
		lib.ResponseJson(c,500,nil,"")
		return
	}
	lib.ResponseJson(c,200,map[string]interface{}{
		"lists": tags,
		"total": count,
	},"")

}

