package main

import (
	"api-gateway/logic"
	"api-gateway/module"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getsubdomain", logic.Find_sub_domain)
	return r
}
func main() {
	err := module.InitMysql()
	if err != nil {
		return
	}
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":5002")
}
