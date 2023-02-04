package routers

import (
	"gin/basic/database"
	"gin/basic/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handler.Handler{
		DB:database.GetDB(),
	}
	router.GET("/books", api.GetBooks)
	router.POST("/books",api.SaveBook)
	return router
}