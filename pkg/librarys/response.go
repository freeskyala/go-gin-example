package librarys

import (
	"github.com/EDDYCJY/go-gin-example/languages"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	DebugInfo interface{} `json:"debugInfo"`
}

//http响应
func ResponseJson(c *gin.Context,code int, data interface{},debugInfo interface{}) {
	c.JSON(200, Response{
		Code: code,
		Msg:  languages.GetErrorMsg(code),
		Data: data,
		DebugInfo: debugInfo,
	})
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  languages.GetErrorMsg(errCode),
		Data: data,
	})
	return
}
