package main

import (
	"com.py/router"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(gin.Logger(), gin.Recovery())

	e.POST("pyfile", router.Savepy)
	e.Run(":8080")
}
