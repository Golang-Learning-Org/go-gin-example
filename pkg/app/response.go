package app

import (
	"github.com/evanxzj/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, code int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

	return
}
