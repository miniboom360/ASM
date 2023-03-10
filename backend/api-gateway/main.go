package main

import (
	"api-gateway/handler"
	"api-gateway/logic"
	"api-gateway/module"
	"github.com/gin-gonic/gin"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getsubdomain", logic.Find_sub_domain)
	r.GET("/scanByTags", logic.NucleiScanByTags)
	r.POST("/addTask", logic.AddTask)
	return r
}
func main() {
	go handler.InitScheduler()
	err := module.InitMysql()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":5002")
}
