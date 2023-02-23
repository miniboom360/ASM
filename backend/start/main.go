package main

import (
  "backend/app/handler"
  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()
  r.POST("/scantask", handler.ScanTaskHandler)
  // r.GET("/scanByTags", logic.NucleiScanByTags)
  // r.POST("/addTask", logic.AddTask)
  return r
}

func main() {
  r := setupRouter()
  r.Run(":5002")
}
