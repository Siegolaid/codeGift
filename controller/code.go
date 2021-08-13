package controller

import (
	"codeGift/service"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func CodeController(path *gin.RouterGroup)  {
	path.GET("", createCode)
}

func createCode(c *gin.Context)  {

	name := c.Query("name")
	service.CreateCode(name)
	log.Println("入参：", name)

	err := qrcode.WriteFile("http://192.168.1.3:8080/page?name=" + name, qrcode.Medium, 256, "./resource/1.png")
	if err != nil {
		c.String(http.StatusOK, "二维码生成失败")
		return
	}

	c.File("./resource/1.png")
}