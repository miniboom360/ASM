package logic

import (
	"api-gateway/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TaskReq struct {
	Plan       string   `json:"plan"`
	Org_name   string   `json:"org_name"`
	Domains    []string `json:"domains"`
	ScanPolicy string   `json:"scan_policy"`
}

func AddTask(c *gin.Context) {
	req := TaskReq{}

	c.BindJSON(&req)

	log.Printf("%v", &req)
	if len(req.Domains) == 0 {
		c.String(http.StatusOK, "获取域名列表失败，请检查输入参数内容")
		return
	}

	task_id, err := handler.AddTask(req.Org_name, req.Plan, req.ScanPolicy, req.Domains)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	r := fmt.Sprintf(`{"task_id":%s}`, task_id)
	c.JSON(http.StatusOK, r)
	return
}
