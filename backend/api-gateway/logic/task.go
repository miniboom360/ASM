package logic

import (
	"api-gateway/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddTask(c *gin.Context) {
	req := handler.TaskReq{}

	c.BindJSON(&req)

	log.Printf("%v", &req)
	if len(req.Domains) == 0 {
		c.String(http.StatusOK, "获取域名列表失败，请检查输入参数内容")
		return
	}

	task_id, err := handler.AddTask(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	r := fmt.Sprintf(`{"task_id":%s}`, task_id)
	c.JSON(http.StatusOK, r)
	return
}

// 现在就写后台运行着的任务模块吧，然后将解析选项之类的，都以发消息的方式通知到任务管理器就行。
// 先写一个
