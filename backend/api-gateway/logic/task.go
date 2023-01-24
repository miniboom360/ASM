package logic

import (
	"api-gateway/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTask(c *gin.Context) {

	org_name := c.Query("org_name")
	scan_policy := c.Query("scan_policy")
	domains, ok := c.GetQueryArray("domains")
	if !ok || len(domains) == 0 {
		c.String(http.StatusOK, "获取域名列表失败，请检查输入参数内容")
		return
	}

	tags := c.Query("tags")
	if tags == "" {
		c.String(http.StatusOK, "tags不能为空")
		return
	}

	domains := make([]string, 0)
	domains = append(domains, domain)

	task_id, err := handler.NucleiScanByTags(tags, "", domains)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	r := fmt.Sprintf(`{"task_id":%s}`, task_id)
	c.JSON(http.StatusOK, r)
	return
}
