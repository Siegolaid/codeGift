package main

import (
	"codeGift/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")

	engine.GET("/", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "this is Code Gift App",
		})
	})

	router.InitRouters(engine)

	engine.Run()
}

