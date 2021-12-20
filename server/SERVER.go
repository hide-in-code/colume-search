package server

import (
	"colume-search/server/controller"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const PORT string = ":5200"

func Run(HttpServer *gin.Engine)  {
	HttpServer.POST("/:table", controller.Add)
	HttpServer.GET("/:table", controller.Search)
	HttpServer.GET("/:table/allindex", controller.AllIndex)
	HttpServer.GET("/:table/allindexcount", controller.AllIndexCount)

	err := HttpServer.Run(PORT)
	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		os.Exit(200)
	}

	return
}