package main

import (
	"colume-search/server"
	"github.com/gin-gonic/gin"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	engin := gin.Default()
	server.Run(engin)
}