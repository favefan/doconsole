package app

import (
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/gin-gonic/gin"
)

// Gin includes a *gin.Context
type Gin struct {
	C *gin.Context
}

// Response includes response's content
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response return request's results to client
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
