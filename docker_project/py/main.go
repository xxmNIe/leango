package main

import (
	"com.py/router"
	"com.py/router/dockerOperator"
	"github.com/gin-gonic/gin"
)

func main() {

	e := gin.Default()
	e.Use(gin.Logger(), gin.Recovery())

	e.POST("pyfile", gin.HandlerFunc(router.Savepy))
	e.POST("downloadImage", gin.HandlerFunc(dockerOperator.ImageDownload))
	e.Run(":8080")

}
