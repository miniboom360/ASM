package logic

import (
	"api-gateway/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Find_sub_domain(c *gin.Context) {
	// *HostScanResult
	domain := c.Query("domain")
	if domain == "" {
		c.String(http.StatusOK, "域名不能为空")
		return
	}

	task_id, err := handler.FindSubDomain(domain, "")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	r := fmt.Sprintf(`{"task_id":%s}`, task_id)
	c.JSON(http.StatusOK, r)
	return
}
