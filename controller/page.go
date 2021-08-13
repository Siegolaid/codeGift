package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PageController(path *gin.RouterGroup)  {
	path.GET("", view)
}

func view(c *gin.Context)  {

	c.HTML(http.StatusOK, "page1.tmpl", gin.H{
		"name": c.Query("name"),
	})

}
