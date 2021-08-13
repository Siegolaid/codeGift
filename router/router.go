package router

import (
	"codeGift/controller"
	"github.com/gin-gonic/gin"
)

func InitRouters(engine *gin.Engine)  {

	controller.CodeController(engine.Group("code"))
	controller.PageController(engine.Group("page"))

}