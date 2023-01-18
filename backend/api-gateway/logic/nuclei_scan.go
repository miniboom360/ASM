package logic

import (
	"api-gateway/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NucleiScanByTags(c *gin.Context) {
	// *HostScanResult
	domain := c.Query("domain")
	if domain == "" {
		c.String(http.StatusOK, "域名不能为空")
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
