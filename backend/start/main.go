package main

import (
	"backend/app/handler"
	"backend/app/module"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/scantask", handler.ScanTaskHandler)
	r.POST("login", handler.LoginHandler)
	// r.GET("/scanByTags", logic.NucleiScanByTags)
	// r.POST("/addTask", logic.AddTask)
	return r
}

func main() {
	err := module.InitMysql()
	if err != nil {
		panic(err)
	}
	r := setupRouter()
	r.Run(":5002")
}
